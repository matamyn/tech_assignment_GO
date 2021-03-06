package http_server

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	pb "github.com/matamyn/tech_assignment_GO/link_shorter/internal/model"
	"net/http"
)

const (
	sessionName = "link_shorter"
	ctxKeyRequestID
)

func (s *HttpServer) setRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

//func (s *HttpServer) logRequest(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		logger := s.logger.WithFields(logrus.Fields{
//			"remote_addr": r.RemoteAddr,
//			"request_id":  r.Context().Value(ctxKeyRequestID),
//		})
//		logger.Infof("started %s %s", r.Method, r.RequestURI)
//
//		start := time.Now()
//		rw := &responseWriter{w, http.StatusOK}
//		next.ServeHTTP(rw, r)
//
//		var level logrus.Level
//		switch {
//		case rw.code >= 500:
//			level = logrus.ErrorLevel
//		case rw.code >= 400:
//			level = logrus.WarnLevel
//		default:
//			level = logrus.InfoLevel
//		}
//		logger.Logf(
//			level,
//			"completed with %d %s in %v",
//			rw.code,
//			http.StatusText(rw.code),
//			time.Now().Sub(start),
//		)
//	})
//}

func (s *HttpServer) handlerGetLink() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &pb.GetLinkRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		linkDbRow, err := s.dbFacade.GetLink(req.ShortUrl)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		l := pb.GetLinkResponse{Url: linkDbRow.Link_}
		s.respond(w, r, http.StatusOK, l)
	}
}
func (s *HttpServer) handlerCreateShortLink() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		req := &pb.AddShortLinkRequest{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		linkDbRow, err := s.dbFacade.CreateShortLink(req.Url)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		l := pb.AddShortLinkResponse{ShortUrl: linkDbRow.ShortLinkKey_}
		s.respond(w, r, http.StatusCreated, l)
	}
}

func (s *HttpServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}
func (s *HttpServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}
