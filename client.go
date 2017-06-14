// Package paypal provides a PayPal SDK.
package paypal

// Client keeps the data to make a correct request to PayPal.
type Client struct {
	Token      string
	Production bool
}

// URL returns the URL used for a client request (whether client is set for production or not)
func (c *Client) URL() string {
	if c.Production {
		return "https://www.paypal.com/cgi-bin/webscr"
	}
	return "https://www.sandbox.paypal.com/cgi-bin/webscr"
}
