package domain

type Parser interface {
	Parse(string) (string, string, error)
}

type UI interface {
	GetCommand() string
	ShowMessage(string)
}

type Client interface {
	ReduceURL(string) string
	RestoreURL(string) string
	OpenBrowser(string)
}