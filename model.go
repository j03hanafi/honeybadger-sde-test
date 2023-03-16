package main

type Config struct {
	Addr    string
	Key     string
	Channel string
}

type Payload struct {
	RecordType    string `json:"RecordType"`
	Type          string `json:"Type"`
	TypeCode      int    `json:"TypeCode"`
	Name          string `json:"Name"`
	Tag           string `json:"Tag"`
	MessageStream string `json:"MessageStream"`
	Description   string `json:"Description"`
	Email         string `json:"Email"`
	From          string `json:"From"`
	BouncedAt     string `json:"BouncedAt"`
}
