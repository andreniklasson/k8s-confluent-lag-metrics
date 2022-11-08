package integration

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"sort"
	"time"

	"github.com/andreniklasson/k8s-confluent-lag-metrics/integration/model"
)

type ConfluentIntegration struct {
	authorizationHeader string
	httpClient          http.Client
	baseUrl             string
	clusterId           string
}

func NewConfluentIntegration(url string, apiKey string, apiSecret string, clusterId string, timeOut int) *ConfluentIntegration {
	authorizationHeader := b64.StdEncoding.EncodeToString([]byte(apiKey + ":" + apiSecret))
	httpClient := http.Client{
		Timeout: (time.Second * time.Duration(timeOut)),
	}
	return &ConfluentIntegration{
		authorizationHeader: "Basic " + authorizationHeader,
		httpClient:          httpClient,
		baseUrl:             url,
		clusterId:           clusterId,
	}
}

func (ci *ConfluentIntegration) QueryConsumerLag(consumerGroup string) (float32, error) {
	request := model.CreateRequest(consumerGroup, ci.clusterId)
	httpRequest, err := httpRequest(request, ci.baseUrl, ci.authorizationHeader)

	res, err := ci.httpClient.Do(httpRequest)
	if err != nil {
		return 0.0, err
	}

	response, err := handleHttpResponse(res)
	if err != nil {
		return 0.0, err
	}
	return getHighestConsumerLagValue(response), nil
}

func handleHttpResponse(httpResponse *http.Response) (model.Response, error) {
	if httpResponse.Body != nil {
		defer httpResponse.Body.Close()
	}

	if httpResponse.StatusCode != 200 {
		return model.Response{}, errors.New("Received non 200 code from contentful")
	}

	body, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return model.Response{}, err
	}

	response := model.Response{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return model.Response{}, err
	}
	return response, nil
}

func httpRequest(request model.Request, baseUrl string, authorizationHeader string) (*http.Request, error) {
	jsonRequest, err := json.Marshal(request)
	if err != nil {
		return &http.Request{}, err
	}

	httpRequest, err := http.NewRequest("POST", baseUrl, bytes.NewBuffer(jsonRequest))
	if err != nil {
		return &http.Request{}, err
	}
	httpRequest.Header.Add("Authorization", authorizationHeader)
	httpRequest.Header.Add("Content-Type", "application/json")
	return httpRequest, nil
}

func getHighestConsumerLagValue(response model.Response) float32 {
	if len(response.Data) == 0 {
		return 0.0
	}

	sort.Slice(response.Data, func(i, j int) bool {
		return response.Data[i].Value > response.Data[j].Value
	})
	return response.Data[0].Value
}
