package cart

import (
	"os/user"
	"testing"
	"github.com/stretchr/testify/assert"
	"time"

	"example.com/modules/product"
	"github.com/Rhymond/go-money"
)

func TestTotalPrice(t *testing.T) {
    items := []Item{
        {
            Product: product.Product{
                ID:    "p-1254",
                Name:  "Product test",
                Price: money.New(1000, "EUR"),
            },
            Quantity: 2,
        },
        {
            Product: product.Product{
                ID:    "p-1255",
                Name:  "Product test 2",
                Price: money.New(2000, "EUR"),
            },
            Quantity: 1,
        },
    }
    c := Cart{
        ID:           "1254",
        CreatedAt:    time.Now(),
        UpdatedAt:    time.Now(),
        User:         user.User{},
        Items:        items,
        CurrencyCode: "EUR",
    }
    actual, err := c.TotalPrice()
    assert.NoError(t, err)
    assert.Equal(t, money.New(4000, "EUR"), actual)
}