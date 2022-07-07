package svc

import (
	"go-rest-api/datastore"
	"net/http"
)

func (s *Server) GetAllEmployees(w http.ResponseWriter, req *http.Request) {
	employees, err := s.svc.GetAllEmployees()
	if err != nil {
		s.respond(w, nil, nil, http.StatusInternalServerError, err)
		return
	}
	type Response struct {
		Results []datastore.Employee `json:"results"`
	}
	s.respond(w, req, Response{Results: employees}, http.StatusOK, nil)
}
