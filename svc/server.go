package svc

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"

	"go-rest-api/config"
	"go-rest-api/utils"

	"github.com/pkg/errors"
)

type Server struct {
	router *mux.Router
	config *config.HTTPServer
	httpS  http.Server
	logger log.Logger
	svc    *Services
}

func NewServer(conf *config.HTTPServer, logger log.Logger, router *mux.Router, services *Services) *Server {
	s := Server{
		router: router,
		config: conf,
		logger: logger,
		svc:    services,
	}

	s.httpS = http.Server{
		Addr:         *conf.ListenAddr,
		WriteTimeout: *conf.WriteTimeout,
		ReadTimeout:  *conf.ReadTimeout,
		IdleTimeout:  *conf.IdleTimeout,
		Handler:      s.CorsMW(),
	}
	s.routes()
	return &s
}

func (s *Server) Start() error {
	return s.httpS.ListenAndServe()
}

// Shutdown gracefully terminates the server
func (s *Server) Shutdown() error {
	return s.httpS.Shutdown(context.Background())
}

func (s *Server) respond(w http.ResponseWriter, req *http.Request, data interface{}, status int, err error) {
	defer s.logger.Log("status", status, "err", err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err != nil {
		errJ := json.NewEncoder(w).Encode(struct{ Err string }{Err: err.Error()})
		if errJ != nil {
			s.logger.Log("err", errors.Errorf("Failed to Encode result to JSON:%v", err))
		}
	}
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.logger.Log("err", errors.Errorf("Failed to Encode result to JSON:%v", err))
		}
	}
}

func (s *Server) respondWithErrors(w http.ResponseWriter, statusCode int, err ...*utils.Error) {
	defer s.logger.Log("status", statusCode, "err", err)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	if err != nil {
		err := json.NewEncoder(w).Encode(err)
		if err != nil {
			s.logger.Log("err", errors.Errorf("Failed to Encode result to the JSON:%v", err))
		}
	}
}
