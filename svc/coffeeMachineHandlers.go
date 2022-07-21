package svc

import (
	"go-rest-api/datastore"
	"net/http"
)

func (s *Server) GetRawMaterials(w http.ResponseWriter, req *http.Request) {
	rawMaterials, err := s.svc.GetRawMaterials()
	if err != nil {
		s.respond(w, nil, nil, http.StatusInternalServerError, err)
		return
	}
	type Response struct {
		Results []datastore.RawMaterial `json:"results"`
	}
	s.respond(w, req, Response{Results: rawMaterials}, http.StatusOK, nil)
}

func (s *Server) GetCoffeeTypes(w http.ResponseWriter, req *http.Request) {
	coffeeTypes, err := s.svc.GetCoffeeTypes()
	if err != nil {
		s.respond(w, nil, nil, http.StatusInternalServerError, err)
		return
	}
	type Response struct {
		Results []string `json:"results"`
	}
	s.respond(w, req, Response{Results: coffeeTypes}, http.StatusOK, nil)
}
