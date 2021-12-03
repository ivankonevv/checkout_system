package upload

import (
	"checkout_system/models"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func ReadJSONCatalog() ([]models.Product, error) {

	log := logrus.New().WithField("function", "ReadProducts()")

	inputFile, err := os.Open(os.Getenv("INPUT_FILENAME"))
	if err != nil {
		log.Error("invalid path: ", err)
		return nil, errors.Wrap(err, "failed to open input file.")
	}

	data, err := ioutil.ReadAll(inputFile)
	if err != nil {
		log.Error("cannot read input file: ", err)
		return nil, errors.Wrap(err, "failed to read input file.")
	}

	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
			log.Errorf("failed to close `%s`", inputFile.Name())
		}
	}(inputFile)

	var productList []models.Product
	err = json.Unmarshal(data, &productList)

	if err != nil {
		log.Error("unmarshalling error: ", err)
		return nil, errors.Wrap(err, "failed to unmarshall input file.")
	}

	return productList, nil
}
