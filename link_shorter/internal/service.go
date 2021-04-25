package link_shorter

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/db_facade"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type LinkShorterService struct {
	config   *common.Config
	logger   *logrus.Logger
	router   *mux.Router
	dbFacade *db_facade.DbFacade
}

func linkShorterService(config *common.Config) *LinkShorterService {
	service := &LinkShorterService{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
	return service
}

func Start(config *common.Config) error {
	a := linkShorterService(config)

	err := a.configureLogger()
	if err != nil {
		return err
	}

	a.dbFacade, err = db_facade.InitDbFacade(config.DataBase)
	if err != nil {
		return err
	}

	defer a.dbFacade.Db_.Close()
	a.configureRouter()
	a.logger.Info("Start Link Shorter Service")

	return http.ListenAndServe(a.config.Server.Port, a.router)
}

func (a *LinkShorterService) configureLogger() error {
	level, err := logrus.ParseLevel(a.config.Log.Level)
	if err != nil {
		return err
	}
	a.logger.SetLevel(level)
	return nil
}

func (a *LinkShorterService) configureRouter() {
	a.router.HandleFunc("/HELLO", a.handleHello())
}

func (a *LinkShorterService) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello")
	}
}
