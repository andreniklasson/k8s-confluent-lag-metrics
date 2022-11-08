package model

import (
	"fmt"
	"time"
)

type Response struct {
	Kind       string   `json:"kind"`
	ApiVersion string   `json:"apiVersion"`
	MetaData   SelfLink `json:"metadata"`
	Items      []Item   `json:"items"`
}

type SelfLink struct {
	Link string `json:"selfLink"`
}

type Item struct {
	MetricName      string          `json:"metricName"`
	Timestamp       string          `json:"timestamp"`
	Value           string          `json:"value"`
	DescribedObject DescribedObject `json:"describedObject"`
}

type DescribedObject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

func NewResponse(name string, consumerLag float32) Response {
	describedObject := DescribedObject{
		Kind: "Pod",
		Name: name,
	}
	t := time.Now()
	ts := t.Format("2006-01-02T15:04:05Z07:00")
	item := Item{
		MetricName:      "consumer-lag",
		Timestamp:       ts,
		Value:           fmt.Sprintf("%f", consumerLag),
		DescribedObject: describedObject,
	}
	selfLink := SelfLink{
		Link: "/apis/custom.metrics.k8s.io/v1beta1/",
	}
	response := Response{
		Kind:       "MetricValueList",
		ApiVersion: "custom.metrics.k8s.io/v1beta1",
		MetaData:   selfLink,
		Items: []Item{
			item,
		},
	}
	return response
}
