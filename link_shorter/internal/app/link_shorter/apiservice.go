package link_shorter

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type LinkShorterService struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	db     *sql.DB
}

func New(config *Config) *LinkShorterService {
	return &LinkShorterService{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
	}
}

func (a *LinkShorterService) Start() error {
	if err := a.configureLogger(); err != nil {
		return err
	}
	a.configureRouter()

	if err := a.ConnectDB(); err != nil {
		return err
	}

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
func (s *LinkShorterService) ConnectDB() error {
	full_url :=
		s.config.MySqlDB.User + ":" +
			s.config.MySqlDB.Password + "@/" +
			s.config.MySqlDB.DbName
	db, err := sql.Open("mysql", full_url)
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	s.db = db
	return nil
}
