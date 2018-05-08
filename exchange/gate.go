package exchange

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// https://www.zb.com/i/developer
const gateBaseApi = "http://data.gateio.io/api2/1/"

// ZB api is very similar to OKEx, who copied whom?

type gateClient struct {
	exchangeBaseClient
	AccessKey string
	SecretKey string
}

type gateCommonResponse struct {
	Result  string
	Message *string
}

type gateTickerResponse struct {
	gateCommonResponse
	Last float64 `json:",string"`
}

type gateKlineResponse struct {
	gateCommonResponse
	Data [][]float64
}

func (resp *gateTickerResponse) getCommonResponse() gateCommonResponse {
	return resp.gateCommonResponse
}

func (resp *gateKlineResponse) getCommonResponse() gateCommonResponse {
	return resp.gateCommonResponse
}

// Any way to hold the common response, instead of adding an interface here?
type gateCommonResponseProvider interface {
	getCommonResponse() gateCommonResponse
}

func NewGateClient(httpClient *http.Client) *gateClient {
	return &gateClient{exchangeBaseClient: *newExchangeBase(gateBaseApi, httpClient)}
}

func (client *gateClient) GetName() string {
	return "Gate"
}

func (client *gateClient) decodeResponse(body io.ReadCloser, respJSON gateCommonResponseProvider) error {
	defer body.Close()

	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&respJSON); err != nil {
		return err
	}

	// All I need is to get the common part, I don't like this
	commonResponse := respJSON.getCommonResponse()
	if commonResponse.Message != nil {
		return errors.New(*commonResponse.Message)
	}
	return nil
}

func (client *gateClient) GetKlinePrice(symbol string, groupedSeconds int, size int) (float64, error) {
	symbol = strings.ToLower(symbol)
	rawUrl := client.buildUrl("candlestick2/"+symbol, map[string]string{
		"group_sec":  strconv.Itoa(groupedSeconds),
		"range_hour": strconv.Itoa(size),
	})
	resp, err := client.HTTPClient.Get(rawUrl)
	if err != nil {
		return 0, err
	}

	var respJSON gateKlineResponse
	err = client.decodeResponse(resp.Body, &respJSON)
	if err != nil {
		return 0, err
	}
	if len(respJSON.Data) == 0 {
		return 0, fmt.Errorf("%s - get a zero size kline response", client.GetName())
	}
	logrus.Debugf("%s - Kline for %v hour(s) uses price at %s", client.GetName(), size,
		time.Unix(int64(respJSON.Data[0][0])/1000, 0))
	return respJSON.Data[0][5], nil
}

func (client *gateClient) GetSymbolPrice(symbol string) (*SymbolPrice, error) {
	rawUrl := client.buildUrl("ticker/"+symbol, map[string]string{})
	resp, err := client.HTTPClient.Get(rawUrl)
	if err != nil {
		return nil, err
	}

	var respJSON gateTickerResponse
	err = client.decodeResponse(resp.Body, &respJSON)
	if err != nil {
		return nil, err
	}

	var percentChange1h, percentChange24h = math.MaxFloat64, math.MaxFloat64
	price1hAgo, err := client.GetKlinePrice(symbol, 60, 1)
	if err != nil {
		logrus.Warnf("%s - Failed to get price 1 hour ago, error: %v\n", client.GetName(), err)
	} else if price1hAgo != 0 {
		percentChange1h = (respJSON.Last - price1hAgo) / price1hAgo * 100
	}

	price24hAgo, err := client.GetKlinePrice(symbol, 300, 24) // Seems gate.io only supports 60, 300, 600 etc. seconds
	if err != nil {
		logrus.Warnf("%s - Failed to get price 24 hours ago, error: %v\n", client.GetName(), err)
	} else if price24hAgo != 0 {
		percentChange24h = (respJSON.Last - price24hAgo) / price24hAgo * 100
	}

	return &SymbolPrice{
		Symbol:           symbol,
		Price:            strconv.FormatFloat(respJSON.Last, 'f', -1, 64),
		UpdateAt:         time.Now(),
		Source:           client.GetName(),
		PercentChange1h:  percentChange1h,
		PercentChange24h: percentChange24h,
	}, nil
}

func init() {
	register((&gateClient{}).GetName(), func(client *http.Client) ExchangeClient {
		// Limited by type system in Go, I hate wrapper/adapter
		return NewGateClient(client)
	})
}
