package domain

type Request struct {
	InitialURL string `json:"initial_url"`
}

type Response struct {
	URL string `json:"url"`
}

type UrlDTO struct {
	Id string `db:"hash"`
	InitialURL string `db:"initial_url"`
}