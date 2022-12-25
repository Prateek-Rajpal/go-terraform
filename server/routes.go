package server

import "golang-app/server/handlers"

func (s *server) routes() {
	s.router.Methods("POST").Path("/upload").Handler(handlers.AddEmployeeHandler(s.postgres))
	s.router.Methods("GET").Path("/employees/all").Handler(handlers.GetEmployeesHandler(s.postgres))

}
