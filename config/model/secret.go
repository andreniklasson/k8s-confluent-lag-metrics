package model

type Secret struct {
	ConfluentUri       string `json:"confluentUri"`
	ConfluentApiKey    string `json:"confluentApiKey"`
	ConfluentApiSecret string `json:"confluentApiSecret"`
	ClusterId          string `json:"clusterId"`
	Timeout            int    `json:"timeout"`
}
