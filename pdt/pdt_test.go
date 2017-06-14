package pdt

import (
	"fmt"
	"strings"
	"testing"
)

const rawTransaction = `SUCCESS
transaction_subject=
payment_date=01%3A01%3A01+May+01%2C+2017+PDT
txn_type=web_accept
last_name=buyer
residence_country=BE
item_name=My+Product
payment_gross=
mc_currency=EUR
business=mybusiness%40example.com
payment_type=instant
protection_eligibility=Ineligible
payer_status=verified
tax=0.00
payer_email=buyer%40example.com
txn_id=EPC66XON1D4EE27M9
quantity=1
receiver_email=mybusiness%40example.com
first_name=test
payer_id=KNTZ5PLNVAQLR
receiver_id=ZAB4D5P0VGMAI
item_number=myproduct
handling_amount=0.00
payment_status=Completed
payment_fee=
mc_fee=1.34
shipping=0.00
mc_gross=29.00
custom=
charset=windows-125`

func TestParseTransaction(t *testing.T) {
	tx, err := parseTransaction(strings.NewReader(rawTransaction))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(tx)
}

func BenchmarkParseTransaction(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseTransaction(strings.NewReader(rawTransaction))
	}
}
