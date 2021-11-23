package upload

import (
	"checkout_system/models"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

type Products struct {
	Products []models.Product `json:"products"`
}

func (p *Products) LoadProductPrices() error {
	log := logrus.New().WithField("function", "ReadProducts()")

	inputFile, err := os.Open(os.Getenv("INPUT_FILENAME"))
	if err != nil {
		log.Error("invalid path: ", err)
		return errors.Wrap(err, "failed to open input file.")
	}

	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Error("cannot read input file: ", err)
		return errors.Wrap(err, "failed to read input file.")
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
		return errors.Wrap(err, "failed to unmarshall input file.")
	}

	log.Infof("successfully read %d products", len(p.Products))

	return nil
}
