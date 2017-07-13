# paypal [![GoDoc](https://godoc.org/github.com/arthurwhite/paypal?status.svg)](https://godoc.org/github.com/arthurwhite/paypal) [![Build](https://travis-ci.org/arthurwhite/paypal.svg?branch=master)](https://travis-ci.org/arthurwhite/paypal) [![Coverage](https://coveralls.io/repos/github/arthurwhite/paypal/badge.svg?branch=master)](https://coveralls.io/github/arthurwhite/paypal?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/arthurwhite/paypal)](https://goreportcard.com/report/github.com/arthurwhite/paypal) ![Status Testing](https://img.shields.io/badge/status-testing-orange.svg)

Package [paypal](https://godoc.org/github.com/arthurwhite/paypal) provides a PayPal SDK.

## Installing

1. Get package:

   ```Shell
   go get -u github.com/arthurwhite/paypal/...
   ````

2. Import it in your code:

   ```Go
   import "github.com/arthurwhite/paypal"
   ```

## Usage

Make a new client:

```Go
pp := &paypal.Client{
	Token:      "G-ddvHQfRB2wqzrHCgdkbx0uXEcgKTcWbG2GjlI581zbPbGxKekGXgyVwU0",
	Production: false,
}
```

### Payment Data Transfer

Import package [pdt](https://godoc.org/github.com/arthurwhite/paypal/pdt) in your code:

```Go
import "github.com/arthurwhite/paypal/pdt"
```

#### Get transaction

Use [GetTransaction](https://godoc.org/github.com/arthurwhite/paypal#paypal.GetTransaction) to retreive a transaction by its ID:

```Go
tx, _ := pdt.GetTransaction(pp, "EPC66XON1D4EE27M9")
fmt.Println(tx)
```
