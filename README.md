#Checkout system

##Installation:
```
$ git clone git@github.com:ivankonevv/checkout_system.git
$ cd checkout_system
```

You can simply run or build script using Makefile:
`$ make run`, `$ make build`

The script accepts JSON catalog as input in the following format:
```
[
  {
    "sku": string,
    "name": string,
    "price": float
  },
  {
    "sku": string,
    "name": string,
    "price": float
  }
]
```

The script takes data from the catalog and puts the selected ones in the shopping cart. If one of the three special offers is valid for the selected products, the cost of the products calculating taking the special offer.

###Available specials:
-*3 for 2 deal on Apple TVs.*

-*Price reduction to $499.99 for iPad Super if there are more than 4 in the cart.*

-*Free VGA for every MacBook Pro **(if VGA already is in the cart)***

You can add items to the cart using the `Scan()` method:
`myShoppingCart.Scan("ipd")`


###Example:
```go
package main

import (
	"checkout_system/cart",
	"checkout_system/offers"
)


func main() {
	cart.LoadCatalog()
	myShoppingCart := cart.New()

	// Add 4 iPads to the cart
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")
	myShoppingCart.Scan("ipd")

	// Available special offers.
	specialOffers := []func(productCart *cart.ProductCart) *cart.ProductCart{
		offers.ApplyIpadPrice,
	}

	// Apply special offers.
	for _, o := range specialOffers {
		myShoppingCart = o(myShoppingCart)
	}

	// Get final amount of your cart.
	myShoppingCart.Total()
}
```

*Result:*
`
INFO[0000] Super iPad                                    SKU=ipd count=4 total=1999.96
`