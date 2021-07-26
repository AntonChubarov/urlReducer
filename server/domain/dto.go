package domain

type Request struct {
	InitialURL string `json:"initial_url"`
}

type Response struct {
	URL string `json:"url"`
}

type URLDTO struct {
	ID         string `db:"hash"`
	InitialURL string `db:"initial_url"`
}