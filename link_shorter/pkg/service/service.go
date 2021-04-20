package service

import "github.com/kardianos/service"

var logger service.Logger

type program struct {
}

func (p *program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *program) run() {
}

func (p *program) Stop(s service.Service) error {
	return nil
}
