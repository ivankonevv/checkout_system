package main

import (
	"checkout_system/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Products struct {
	Products []Product `json:"products"`
}

type Product struct {
	SKU   string  `json:"SKU"`
	Name  string  `json:"Name"`
	Price float32 `json:"Price"`
}

func (p *Products) ReadProducts() {
	log := logrus.New().WithField("function", "ReadProducts()")

	inputFile, err := os.Open(os.Getenv("INPUT_FILENAME"))
	if err != nil {
		log.Error("invalid path: ", err)
		return
	}

	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Error("cannot read input file: ", err)
		return
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Errorf("failed to close `%s`", inputFile.Name())
		}
	}(inputFile)

	err = json.Unmarshal(data, &p)
	if err != nil {
		log.Error("unmarshalling error: ", err)

	}
	log.Debug("successfully read %d products", len(p.Products))
}

func (p Product) WithDropSpecial(cnt int) {
	if cnt >= 4 {
		pricePerOne := p.Price - 50.0
		finalPrice := pricePerOne * float32(cnt)
		utils.LogInfo(cnt, p.Name, finalPrice)

		return
	}
	finalPrice := p.Price * float32(cnt)
	utils.LogInfo(cnt, p.Name, finalPrice)

	return
}

func (p Product) ThreeForTwoSpecial(cnt int) {
	if cnt == 3 {
		finalPrice := p.Price*float32(cnt) - p.Price
		utils.LogInfo(cnt, p.Name, finalPrice)

		return
	}
	finalPrice := p.Price * float32(cnt)
	utils.LogInfo(cnt, p.Name, finalPrice)

	return
}

func (p Product) DefaultPrice(cnt int) {
	finalPrice := p.Price * float32(cnt)
	utils.LogInfo(cnt, p.Name, finalPrice)

	return

}

func (p Product) FreeVGASpecial(cnt int, isMac bool, macCnt int) {
	if isMac {
		if cnt == macCnt {
			finalPrice := float32(0.0)
			utils.LogInfo(cnt, p.Name, finalPrice)
			return
		}
		if cnt > macCnt {
			finalPrice := p.Price * (float32(cnt) - float32(macCnt))
			utils.LogInfo(cnt, p.Name, finalPrice)
			return
		}
		if cnt < macCnt && macCnt > 1 {
			finalPrice := p.Price * (float32(macCnt) - float32(cnt))
			utils.LogInfo(cnt, p.Name, finalPrice)
			return
		}
		return
	}
	finalPrice := p.Price * float32(cnt)
	utils.LogInfo(cnt, p.Name, finalPrice)

	return
}

func (p Products) GetFinalPrice(skus map[string]int) {
	var isMac bool
	var macCnt int
	for _, product := range p.Products {
		for i, cnt := range skus {
			if product.SKU == i {
				switch i {
				case "ipd":
					product.WithDropSpecial(cnt)
				case "atv":
					product.ThreeForTwoSpecial(cnt)
				case "mbp":
					isMac = true
					macCnt = cnt
					product.DefaultPrice(cnt)
				case "vga":
					if isMac {
						product.FreeVGASpecial(cnt, true, macCnt)
					} else {
						product.FreeVGASpecial(cnt, false, 0)
					}
				}
			}
		}
	}
	logrus.Infof("finished")

}

func init() {
	log := logrus.New().WithField("function", "main()")
	if err := godotenv.Load(".env"); err != nil {
		log.Error("an error occurred while load env variables:", err)
		return
	}
}

func main() {
	var p Products
	var skuSlice = []string{"vga", "ipd", "ipd", "ipd", "ipd"}

	p.ReadProducts()

	skuMap := utils.DupCounter(skuSlice)

	p.GetFinalPrice(skuMap)
}
