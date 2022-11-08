package model

type Request struct {
	Aggregations []Metric `json:"aggregations"`
	Filter       Filter   `json:"filter"`
	Granularity  string   `json:"granularity"`
	GroupBy      []string `json:"group_by"`
	Intervals    []string `json:"intervals"`
	Limit        int      `json:"limit"`
}

type Filter struct {
	Operator string          `json:"op"`
	Filters  []FilterElement `json:"filters"`
}

type FilterElement struct {
	Field    string `json:"field"`
	Operator string `json:"op"`
	Value    string `json:"value"`
}

type Metric struct {
	Metric string `json:"metric"`
}

func CreateRequest(consumerGroup string, clusterId string) Request {
	metric := Metric{
		Metric: "io.confluent.kafka.server/consumer_lag_offsets",
	}
	aggregations := []Metric{metric}

	clusterFilter := FilterElement{
		Field:    "resource.kafka.id",
		Operator: "EQ",
		Value:    clusterId,
	}
	consumerGroupFilter := FilterElement{
		Field:    "metric.consumer_group_id",
		Operator: "EQ",
		Value:    consumerGroup,
	}
	filters := []FilterElement{
		clusterFilter,
		consumerGroupFilter,
	}
	filter := Filter{
		Operator: "AND",
		Filters:  filters,
	}
	request := Request{
		Aggregations: aggregations,
		Filter:       filter,
		Granularity:  "PT1M",
		GroupBy: []string{
			"metric.topic",
		},
		Intervals: []string{
			"PT3M/now",
		},
		Limit: 10,
	}
	return request
}
