package server

import (
	"io"
	"log"
	"net/http"
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

type GitData struct {
	projects      []*git.Project
	branches      []*git.Branch
	mergeRequests []*git.MergeRequest
}

type Server struct {
	options *ServerOptions
	git     git.GitProvider
	data    GitData
}

func NewServer(options *ServerOptions) (*Server, error) {
	git, err := git.NewProvider(options.BaseURL, options.Token, options.Provider)
	if err != nil {
		return nil, err
	}

	return &Server{options: options, git: git}, nil
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

	return router
}

func (s *Server) reloadHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		s.Reload()

		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "reload")
	})
}
