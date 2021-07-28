package app

import (
	"client/domain"
	"log"
)

type Service struct {
	userInterface domain.UI
	commandParser domain.Parser
	client domain.Client
}

func NewService(ui domain.UI, p domain.Parser, cli domain.Client) *Service {
	return &Service{
		userInterface: ui,
		commandParser: p,
		client: cli,
	}
}

func (s *Service) Run() {
Loop:
	for{
		input := s.userInterface.GetCommand()

		command, url, err := s.commandParser.Parse(input)
		if err != nil {
			log.Println(err)
			continue Loop
		}

		if _, ok := domain.Commands[command]; !ok {
			log.Println(domain.ErrorInvalidCommand)
			continue Loop
		}

		switch command {
		case "reduce":
			reducedLink := s.client.ReduceURL(url)
			s.userInterface.ShowMessage(reducedLink)
		case "open":
			s.client.OpenBrowser(url)
		default:
			log.Println(domain.ErrorInvalidCommand)
			continue Loop
		}
	}
}