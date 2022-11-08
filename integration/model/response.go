package model

type Response struct {
	Data []Entry `json:"data"`
}

type Entry struct {
	Timestamp string  `json:"timestamp"`
	Value     float32 `json:"value"`
	Topic     string  `json:"metric.topic"`
}
