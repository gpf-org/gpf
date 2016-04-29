package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gpf-org/gpf/git"
)

type ServerOptions struct {
	Provider  string
	Token     string
	BaseURL   string
	PublicURL string
	Bind      string
	Port      int
}

type Server struct {
	status ServerStatus
	git    git.GitProvider
}

func (s *Server) Start(options *ServerOptions) error {
	gp, err := git.NewProvider(options.BaseURL, options.Token, options.Provider)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	s.git = gp

	s.reload()

	router := s.createRouter()
	return listenAndServe(options.Bind, options.Port, router)
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
	log.Printf("Server running on http://%s", addr)
	return http.ListenAndServe(addr, handler)
}
