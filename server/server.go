package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/gpf-org/gpf/core"
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
	options  *ServerOptions
	provider git.GitProvider
	store    *core.Store
}

func NewServer(options *ServerOptions) (*Server, error) {
	provider, err := git.NewProvider(options.BaseURL, options.Token, options.Provider)
	if err != nil {
		return nil, err
	}

	store := core.NewStore(core.NewIssueNameRegexp("^([0-9]+)/.*$"))

	return &Server{options: options, provider: provider, store: store}, nil
}

func (s *Server) ListenAndServe() error {
	router := s.createRouter()

	addr := s.options.Bind + ":" + strconv.Itoa(s.options.Port)

	log.Printf("Server running on http://%s", addr)

	return http.ListenAndServe(addr, router)
}

func (s *Server) Reload() error {
	log.Printf("Reloading the server. It may take awhile.")

	s.store.Reset()

	projs, err := s.provider.ListAllProjects()
	if err != nil {
		return err
	}

	log.Printf("Projects available: %d", len(projs))

	for _, proj := range projs {
		s.store.AddProject(proj)

		// stop creating hooks till we can handle them
		// log.Printf("Project %s: reloading webhook", proj.Name)
		// s.provider.CreateOrUpdateProjectHook(proj.ID, s.options.PublicURL)

		log.Printf("Project %s: reloading branches", proj.Name)
		branches, err := s.provider.ListAllBranches(proj.ID)
		if err != nil {
			return err
		}

		for _, branch := range branches {
			s.store.AddBranch(branch)
		}

		log.Printf("Project %s: reloading merge requests", proj.Name)
		mrs, err := s.provider.ListMergeRequests(proj.ID)
		if err != nil {
			return err
		}

		for _, mergeRequest := range mrs {
			s.store.AddMergeRequest(mergeRequest)
		}
	}

	return nil
}

func (s *Server) createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/reload", s.reloadHandler()).
		Methods("GET")
	router.HandleFunc("/list", s.listHandler()).
		Methods("GET")

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
		data, err := json.Marshal(core.ListIssues(s.store))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}
