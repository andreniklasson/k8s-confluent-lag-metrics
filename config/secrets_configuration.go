package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/andreniklasson/k8s-confluent-lag-metrics/config/model"
)

func ReadSecret(location string) (model.Secret, error) {
	content, err := ioutil.ReadFile(location)
	if err != nil {
		return model.Secret{}, err
	}

	var secret model.Secret
	err = json.Unmarshal(content, &secret)
	if err != nil {
		return model.Secret{}, err
	}
	return secret, nil
}
