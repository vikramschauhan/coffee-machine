package svc

func (s *Server) routes() {

	route := s.router.PathPrefix("/employee").Subrouter()
	route.HandleFunc("", s.GetAllEmployees).Methods("GET")

}
