package server

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ServerOptions struct {
	Token     string
	BaseURL   string
	PublicURL string
	Bind      string
	Port      int
}

type Server struct {
	status ServerStatus
}

func (s *Server) Start(options *ServerOptions) error {
	s.reload()

	router := s.createRouter()
	return listenAndServe(options.Bind, options.Port, router)
}

func (s *Server) reload() {
	s.status = StatusMaintenance
	// reload project information
	// ensure each project has webhook set
	s.status = StatusOK
}

func (s *Server) createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/reload", s.reloadHandler())

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Status: %s", ServerStatusText(s.status))
		switch s.status {
		case StatusOK:
			router.ServeHTTP(w, r)
		case StatusMaintenance:
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
}

func (s *Server) reloadHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.reload()

		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "reload")
	})
}

func listenAndServe(bind string, port int, handler http.Handler) error {
	addr := bind + ":" + strconv.Itoa(port)
	return http.ListenAndServe(addr, handler)
}
