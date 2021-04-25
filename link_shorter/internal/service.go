package link_shorter

import (
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/http_server"
	"github.com/sirupsen/logrus"
)

type LinkShorterService struct {
	config     *common.Config
	logger     *logrus.Logger
	httpServer *http_server.HttpServer
}

func Start(config *common.Config) error {
	a := LinkShorterService{config: config, logger: logrus.New()}
	err := a.configureLogger()
	if err != nil {
		return err
	}
	err = http_server.InitHttpServer(config)

	a.logger.Info("Start Link Shorter Service")

	return err
}

func (a *LinkShorterService) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.Log.Level)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)
	return nil
}
