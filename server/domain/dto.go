package domain

type Request struct {
	InitialURL string `json:"initial_url"`
}

type Response struct {
	URL string `json:"url"`
}