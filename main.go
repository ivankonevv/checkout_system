package main

import (
	"checkout_system/cart"
	"checkout_system/offers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	log := logrus.New().WithField("function", "main()")
	if err := godotenv.Load(".env"); err != nil {
		log.Error("an error occurred while load env variables:", err)
		return
	}
}

func main() {
	myShoppingCart := cart.New()
	myShoppingCart.Add("ipd")
	myShoppingCart.Add("ipd")
	myShoppingCart.Add("ipd")
	myShoppingCart.Add("ipd")
	myShoppingCart.Add("ipd")

	myShoppingCart.Add("atv")
	myShoppingCart.Add("atv")
	myShoppingCart.Add("atv")

	myShoppingCart.Add("mbp")
	myShoppingCart.Add("mbp")
	myShoppingCart.Add("vga")
	myShoppingCart.Add("vga")

	specialOffers := []func(productCart *cart.ProductCart) *cart.ProductCart{
		offers.ApplyIpadPrice,
		offers.ApplyAppleTVSpecialPrice,
		offers.ApplyVGASuperDeal,
	}

	for _, o := range specialOffers {
		myShoppingCart = o(myShoppingCart)
	}
	fmt.Println(myShoppingCart.Total())
	fmt.Println(myShoppingCart.Products)
}
