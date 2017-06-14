package pdt

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/arthurwhite/paypal"
	"io"
	"net/http"
	"net/mail"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

const (
	timeLayout = "15:04:05 Jan 02, 2006 MST"
)

// Errors
var (
	ErrTransactionNotFound = errors.New("pdt: transaction not found")
)

// Transaction is a PDT transaction.
// See https://developer.paypal.com/docs/classic/ipn/integration-guide/IPNandPDTVariables.
//
// BUG(arthurwhite): Can't handle multiple items from a shopping cart transaction.
type Transaction struct {
	ID                    string        `paypal:"txn_id"`
	Type                  string        `paypal:"txn_type"`
	Subject               string        `paypal:"transaction_subject"`
	Business              *mail.Address `paypal:"business"`
	Custom                string        `paypal:"custom"`
	FirstName             string        `paypal:"first_name"`
	HandlingAmount        float64       `paypal:"handling_amount"`
	ItemID                string        `paypal:"item_number"`
	ItemName              string        `paypal:"item_name"`
	LastName              string        `paypal:"last_name"`
	MerchantCurrency      string        `paypal:"mc_currency"`
	MerchantFee           float64       `paypal:"mc_fee"`
	MerchantGross         float64       `paypal:"mc_gross"`
	PayerEmail            *mail.Address `paypal:"payer_email"`
	PayerID               string        `paypal:"payer_id"`
	PayerStatus           string        `paypal:"payer_status"`
	PaymentDate           time.Time     `paypal:"payment_date"`
	PaymentFee            string        `paypal:"payment_fee"`
	PaymentGross          string        `paypal:"payment_gross"`
	PaymentStatus         string        `paypal:"payment_status"`
	PaymentType           string        `paypal:"payment_type"`
	ProtectionEligibility string        `paypal:"protection_eligibility"`
	Quantity              string        `paypal:"quantity"`
	ReceiverID            string        `paypal:"receiver_id"`
	ReceiverEmail         *mail.Address `paypal:"receiver_email"`
	ResidenceCountry      string        `paypal:"residence_country"`
	Shipping              float64       `paypal:"shipping"`
	Tax                   float64       `paypal:"tax"`
}

func (tx *Transaction) String() (s string) {
	txValue := reflect.ValueOf(tx).Elem()
	txType := reflect.TypeOf(*tx)
	s += fmt.Sprintf("%21v  %22v  %v\n", "FIELD", "PAYPAL VARIABLE", "VALUE")
	for i := 0; i < txValue.NumField(); i++ {
		field := txType.Field(i)
		if tagVal, ok := field.Tag.Lookup("paypal"); ok {
			s += fmt.Sprintf("%21v  %22v  %v\n", field.Name, tagVal, txValue.Field(i))
		}
	}
	return
}

// GetTransaction retreives a transaction from id, with client.
// If transaction cannot be found, error is ErrTransactionNotFound.
func GetTransaction(c *paypal.Client, id string) (*Transaction, error) {
	res, err := http.PostForm(c.URL(), url.Values{
		"cmd": {"_notify-synch"},
		"at":  {c.Token},
		"tx":  {id},
	})
	if err != nil {
		return nil, err
	}
	return parseTransaction(res.Body)
}

func parseTransaction(r io.Reader) (*Transaction, error) {
	bs := bufio.NewScanner(r)
	bs.Scan()
	if bs.Text() != "SUCCESS" {
		return nil, ErrTransactionNotFound
	}
	tx := new(Transaction)
	txValue := reflect.ValueOf(tx).Elem()
	txType := reflect.TypeOf(*tx)
	for bs.Scan() {
		t := bs.Text()
		i := strings.IndexByte(t, '=')
		if i == -1 {
			continue
		}
		key := t[:i]
		val, err := url.QueryUnescape(t[i+1:])
		if err != nil {
			return nil, err
		}
		for i := 0; i < txType.NumField(); i++ {
			field := txType.Field(i)
			if tagVal, ok := field.Tag.Lookup("paypal"); !ok || tagVal != key {
				continue
			}
			switch field.Type.String() {
			case "string":
				txValue.Field(i).SetString(val)
			case "float64":
				f, err := strconv.ParseFloat(val, 64)
				if err != nil {
					continue
				}
				txValue.Field(i).SetFloat(f)
			case "*mail.Address":
				m, err := mail.ParseAddress(val)
				if err != nil {
					continue
				}
				txValue.Field(i).Set(reflect.ValueOf(m))
			case "time.Time":
				t, err := time.Parse(timeLayout, val)
				if err != nil {
					continue
				}
				txValue.Field(i).Set(reflect.ValueOf(t))
			}
		}
	}
	return tx, nil
}
