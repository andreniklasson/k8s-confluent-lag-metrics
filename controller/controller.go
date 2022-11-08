package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alexcesaro/log"
	"github.com/andreniklasson/k8s-confluent-lag-metrics/controller/model"
	"github.com/andreniklasson/k8s-confluent-lag-metrics/integration"
	"github.com/gorilla/mux"
)

type Controller struct {
	confluentIntegration *integration.ConfluentIntegration
	logger               log.Logger
}

func NewController(ci *integration.ConfluentIntegration, logger log.Logger) *Controller {
	return &Controller{
		confluentIntegration: ci,
		logger:               logger,
	}
}

func (c *Controller) HealthCheck(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (c *Controller) Consumerlag(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		vars := mux.Vars(req)
		consumerGroup := vars["consumerGroup"]

		consumerLag, err := c.confluentIntegration.QueryConsumerLag(consumerGroup)
		if err != nil {
			c.logger.Error(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response := model.NewResponse(consumerGroup, consumerLag)
			w.Header().Set("Content-Type", "application/json")

			c.logger.Info(consumerGroup + " is currenty " + fmt.Sprintf("%f", consumerLag) + " messages behind.")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
