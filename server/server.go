package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
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
	options *ServerOptions
	git     git.GitProvider
	model   ServerModel
}

func NewServer(options *ServerOptions) (*Server, error) {
	git, err := git.NewProvider(options.BaseURL, options.Token, options.Provider)
	if err != nil {
		return nil, err
	}

	model := &MemoryModel{
		pattern: regexp.MustCompile("^([^/]+)/.*$"),
	}

	return &Server{options: options, git: git, model: model}, nil
}

func (s *Server) ListenAndServe() error {
	router := s.createRouter()

	addr := s.options.Bind + ":" + strconv.Itoa(s.options.Port)

	log.Printf("Server running on http://%s", addr)

	return http.ListenAndServe(addr, router)
}

func (s *Server) createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/reload", s.reloadHandler())
	router.HandleFunc("/list", s.listHandler())

	return router
}

func (s *Server) reloadHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Reload()

		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "reload")
	})
}

func (s *Server) listHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data, err := json.Marshal(s.model.List())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
