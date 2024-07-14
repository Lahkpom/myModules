package invoice

import (
	"github.com/Lahkpom/myModules/ex.3.2/customer"
	"github.com/Lahkpom/myModules/ex.3.2/invoiceitem"
)

// Invoice contains the information of an invoice
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	items   []invoiceitem.Item
}

// New() returns a new invoice
func New(country, city string) *Invoice {
	return &Invoice{country, city, 0, customer.Customer{}, []invoiceitem.Item{}}
}
