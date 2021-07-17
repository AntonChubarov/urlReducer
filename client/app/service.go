package app

import (
	"client/domain"
	"log"
)

type Service struct {
	userInterface domain.UI
	commandParser domain.Parser
	linkValidator domain.Validator
	client domain.Client
}

func NewService(ui domain.UI, p domain.Parser, v domain.Validator, cli domain.Client) *Service {
	return &Service{
		ui,
		p,
		v,
		cli,
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

		//if !s.linkValidator.Validate(url) {
		//	log.Println(domain.ErrorInvalidURL)
		//	continue Loop
		//}

		switch command {
		case "red":
			reducedLink := s.client.ReduceURL(url)
			s.userInterface.ShowMessage(reducedLink)
		case "get":
			initialLink := s.client.RestoreURL(url)
			s.userInterface.ShowMessage(initialLink)
		case "open":
			initialLink := s.client.RestoreURL(url)
			s.client.OpenBrowser(initialLink)
		default:
			log.Println(domain.ErrorInvalidCommand)
			continue Loop
		}
	}
}