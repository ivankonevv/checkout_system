package offers

import (
	"checkout_system/cart"

	"github.com/sirupsen/logrus"
)

func ApplyIpadPrice(cart *cart.ProductCart) *cart.ProductCart {
	if cart.Contains("ipd") && cart.Get("ipd").Len("ipd") > 4 {
		cart.UpdateItemsPrice("ipd", 499.99)
	}
	logrus.Debug("Applying Ipad special...")

	return cart
}

func ApplyAppleTVSpecialPrice(cart *cart.ProductCart) *cart.ProductCart {
	if cart.Contains("atv") && cart.Get("atv").Len("atv") == 3 {

		cart.UpdateOneItemPrice("atv", 0.0)
	}
	logrus.Debug("Applying Apple TV special...")

	return cart
}

func ApplyVGASuperDeal(cart *cart.ProductCart) *cart.ProductCart {
	if cart.Contains("mbp") && cart.Contains("vga") {
		for range cart.Get("mbp").Products["mbp"] {
			cart.UpdateOneItemPrice("vga", 0.0)
		}
	}
	logrus.Debug("Applying VGA special...")

	return cart
}
