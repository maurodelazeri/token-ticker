# Token Ticker

[![Build Status](https://travis-ci.org/maurodelazeri/token-ticker.svg?branch=master)](https://travis-ci.org/maurodelazeri/token-ticker)
[![codecov](https://codecov.io/gh/maurodelazeri/token-ticker/branch/master/graph/badge.svg)](https://codecov.io/gh/maurodelazeri/token-ticker)
[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/maurodelazeri/token-ticker/pulls)
[![Go Report Card](https://goreportcard.com/badge/github.com/maurodelazeri/token-ticker)](https://goreportcard.com/report/github.com/maurodelazeri/token-ticker)

> NEVER LEAVE YOUR TERMINAL

![token-ticker](https://user-images.githubusercontent.com/2657334/40175207-ff9e6504-5a09-11e8-9a3d-a887ebc4895a.png)

Token-ticker (or `tt` for short) is a CLI tool for those who are both **Crypto investors** and **Engineers**, allowing you to track token prices and changes in your favorite exchanges on the terminal.

### Features

 * Auto refresh on a specified interval, watch prices in live update mode
 * Proxy aware HTTP request, for easy access to blocked exchanges
 * Real-time prices from 12+ exchanges

### Supported Exchanges

 * [Binance](https://www.binance.com/)
 * [CoinMarketCap](https://coinmarketcap.com/)
 * [Bitfinex](https://www.bitfinex.com/)
 * [Huobi.pro](https://www.huobi.pro/)
 * [ZB](https://www.zb.com/)
 * [OKEx](https://www.okex.com/)
 * [Gate.io](https://gate.io/)
 * [Bittrex](https://bittrex.com/)
 * [HitBTC](https://hitbtc.com/)
 * [BigONE](https://big.one/)
 * [Poloniex](https://poloniex.com/)
 * [Kraken](https://www.kraken.com/)
 * _still adding..._
 
### Installation

#### Homebrew

```bash
# work in process 
```

#### `curl | bash` style downloads to `/usr/local/bin`
```bash
$ curl -sfL https://raw.githubusercontent.com/maurodelazeri/token-ticker/master/install.sh | bash -s -- -d -b /usr/local/bin
```

#### Using [Go](https://golang.org/) (1.9+)
```bash
$ go get -u github.com/maurodelazeri/token-ticker
```

#### Manually
Download from [release page](https://github.com/maurodelazeri/token-ticker/releases/latest) and extact the tarbal into /usr/bin or your `PATH` directory.

### Usage

```
$ tt --help

Usage: tt [Options] [Exchange1.Token1 Exchange2.Token2 ...]

Track token prices of your favorite exchanges in the terminal

Options:
  -v, --version              Show version number
  -d, --debug                Enable debug mode
  -l, --list-exchanges       List supported exchanges
  -r, --refresh int          Auto refresh on every specified seconds, note every exchange has a rate limit,
                             too frequent refresh may cause your IP banned by their servers
  -c, --config-file string   Config file path, refer to "token_ticker.example.yaml" for the format,
                             by default token-ticker uses "token_ticker.yml" in current directory or $HOME as config file
  -p, --proxy string         Proxy used when sending HTTP request
                             (eg. "http://localhost:7777", "https://localhost:7777", "socks5://localhost:1080")
  -t, --timeout int          HTTP request timeout in seconds (default 20)

Exchange.Token Pairs:
  Specify which exchange and token pair to query, different exchanges use different forms to express tokens/trading pairs, refer to their URLs to find the format, eg. to get BitCoin price from Bitfinex and CoinMarketCap you should use query string "Bitfinex.BTCUSDT CoinMarketCap.Bitcoin"

Find help/updates from here - https://github.com/maurodelazeri/token-ticker
```

* #### Display latest market prices for for `BNBUSDT`, `BTCUSDT` from `Binance` and `HTUSDT` from `Huobi`

```bash
$ tt binance.BNBUSDT binance.BTCUSDT Huobi.HTUSDT
```

Here `Binance` and `Huobi` can be replaced by any supported exchanges, and different exchanges use different forms to express tokens/symbols/markets, refer to their URLs to find the format.

* #### Auto-refresh on every 10 seconds

```bash
$ tt -r 10 binance.BNBUSDT binance.BTCUSDT Huobi.HTUSDT
```

NOTE: some exchanges has a strict rate limit, too frequent refresh may cause your IP banned by their servers.

* #### Run with options from a configuration file

```bash
$ tt -c token_ticker.example.yaml
```

Token-ticker can also read options from configuration file, see the attached [token_ticker.example.yaml](token_ticker.example.yaml) for its format. By default token-ticker searches configuration file `token_ticker.yml` in current directory and `$HOME`, so you can compose a `token_ticker.yml`, place it in your `$HOME` and just type `tt` to get all pre-defined prices. 

```bash
$ # Create your configuration file by copying attached `token_ticker.example.yaml` to `token_ticker.yaml`
$ cp token_ticker.example.yaml token_ticker.yaml
$ # Or copy to your $HOME directory
$ cp token_ticker.example.yaml ~/token_ticker.yaml
$
$
$ # Token-ticker will search for configuration file "token_ticker.yml" in current directory and "$HOME" by default
$ tt       # <--- This is also the way I used most freqently 
```

### Thanks

 * Inspired by [coinmon](https://github.com/bichenkk/coinmon)

### License

The MIT License (MIT) - see [LICENSE.md](https://github.com/maurodelazeri/token-ticker/blob/master/LICENSE) for more details
