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

// SetCountry sets the country of the invoice
func (i *Invoice) SetCountry(country string) { i.country = country }

// Country() returns the country of the invoice
func (i *Invoice) Country() string { return i.country }

// SetCity sets the city of the invoice
func (i *Invoice) SetCity(city string) { i.city = city }

// City() returns the city of the invoice
func (i *Invoice) City() string { return i.city }

// Total() returns the total of the invoice
func (i *Invoice) Total() float64 { return i.total }

// SetClient sets the client of the invoice
func (i *Invoice) SetClient(client customer.Customer) { i.client = client }

// Client() returns the client of the invoice
func (i *Invoice) Client() customer.Customer { return i.client }

// Items() returns the items of the invoice
func (i *Invoice) Items() []invoiceitem.Item { return i.items }

// AddItem() adds an item to the invoice
func (i *Invoice) AddItem(items ...invoiceitem.Item) {
	for _, v := range items {
		i.items = append(i.items, v)
		i.total += v.Value()
	}
}

// TODO RemoveItem() removes an item from the invoice
/*
func (i *Invoice) RemoveItem(items ...invoiceitem.Item) {
	for _, v := range items {
		for _, i := range i.items {
			if v.ID() == i.ID() {
				i.items = (LÓGICA PARA REMOVER)
				i.total -= v.Value()
			}
		}
	}
}
*/
