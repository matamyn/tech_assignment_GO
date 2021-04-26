package http_server

import (
	"github.com/gorilla/mux"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/common"
	"github.com/matamyn/tech_assignment_GO/link_shorter/internal/db_facade"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type HttpServer struct {
	router   *mux.Router
	logger   *logrus.Logger
	dbFacade *db_facade.DbFacade
}

func newHttpServer(facade *db_facade.DbFacade) *HttpServer {
	s := &HttpServer{router: mux.NewRouter(), logger: logrus.New(), dbFacade: facade}
	s.configureRouter()
	return s
}

func InitHttpServer(config *common.Config) error {
	facade, err := db_facade.InitDbFacade(&config.DataBase)
	if err != nil {
		return err
	}
	defer facade.Db_.Close()
	s := newHttpServer(facade)
	s.configureRouter()
	return http.ListenAndServe(config.Server.Port, s.router)
}

func (s *HttpServer) ServeHttp(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

//Handler Factory
func (a *HttpServer) configureRouter() {
	a.router.Use(a.setRequestID)
	//a.router.Use(a.logRequest)
	a.router.HandleFunc("/HELLO", a.handleHello())
	a.router.HandleFunc("/GetLink", a.handlerGetLink()).Methods("POST")
	a.router.HandleFunc("/SetLink", a.handlerCreateShortLink()).Methods("POST")
}

func (a *HttpServer) handleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "hello")
	}
}
