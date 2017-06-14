package pdt_test

import (
	"fmt"
	"github.com/arthurwhite/paypal"
	"github.com/arthurwhite/paypal/pdt"
)

func Example() {
	pp := &paypal.Client{
		Token:      "G-ddvHQfRB2wqzrHCgdkbx0uXEcgKTcWbG2GjlI581zbPbGxKekGXgyVwU0",
		Production: false,
	}

	tx, err := pdt.GetTransaction(pp, "EPC66XON1D4EE27M9")
	if err != nil {
		panic(err)
	}
	fmt.Println(tx)
}
