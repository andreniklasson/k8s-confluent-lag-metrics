package main

import (
	"net/http"

	"github.com/alexcesaro/log/stdlog"
	"github.com/andreniklasson/k8s-confluent-lag-metrics/config"
	"github.com/andreniklasson/k8s-confluent-lag-metrics/controller"
	"github.com/andreniklasson/k8s-confluent-lag-metrics/integration"
	"github.com/gorilla/mux"
)

func main() {
	logger := stdlog.GetFromFlags()
	secret, err := config.ReadSecret("/mnt/secrets/confluent-lag-metrics")
	if err != nil {
		panic(err)
	}

	confluentIntegration := integration.NewConfluentIntegration(
		secret.ConfluentUri,
		secret.ConfluentApiKey,
		secret.ConfluentApiSecret,
		secret.ClusterId,
		secret.Timeout,
	)

	controller := controller.NewController(confluentIntegration, logger)

	r := mux.NewRouter()
	r.HandleFunc("/apis/custom.metrics.k8s.io/v1beta1", controller.HealthCheck)
	r.HandleFunc("/apis/custom.metrics.k8s.io/v1beta1/namespaces/{namespace}/pods/{consumerGroup}/consumer-lag", controller.Consumerlag)

	http.ListenAndServeTLS(":8090", "/tmp/serving-certs/tls.crt", "/tmp/serving-certs/tls.key", r)
	http.Handle("/", r)
}
