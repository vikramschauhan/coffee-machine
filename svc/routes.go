package svc

func (s *Server) routes() {

	route := s.router.PathPrefix("/coffee-machine/startMachine").Subrouter()
	route.HandleFunc("", s.GetRawMaterials).Methods("GET")

	route = s.router.PathPrefix("/coffee-machine/rawMaterialsAvailable").Subrouter()
	route.HandleFunc("", s.GetRawMaterials).Methods("GET")

	route = s.router.PathPrefix("/coffee-machine/coffeeTypes").Subrouter()
	route.HandleFunc("", s.GetCoffeeTypes).Methods("GET")

}
