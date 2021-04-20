package link_shorter

import (
	"github.com/sirupsen/logrus"
)

type LinkShorterService struct {
	config *Config
	logger *logrus.Logger
}

func New(config *Config) *LinkShorterService {
	return &LinkShorterService{config: config, logger: logrus.New()}

}

func (a *LinkShorterService) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	a.logger.Info("Start Link Shorter Service")
	return nil
}
func (a *LinkShorterService) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.Log.Level)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)
	return nil
}
