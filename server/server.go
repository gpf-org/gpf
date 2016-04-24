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

func (self *Server) Start(options *ServerOptions) error {
	self.reload()

	router := self.createRouter()
	return listenAndServe(options.Bind, options.Port, router)
}

func (self *Server) reload() {
	self.status = StatusMaintenance
	// reload project information
	// ensure each project has webhook set
	self.status = StatusOK
}

func (self *Server) createRouter() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/reload", reloadHandler(self))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Status: %s", ServerStatusText(self.status))
		switch self.status {
		case StatusOK:
			router.ServeHTTP(w, r)
		case StatusMaintenance:
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})
}

func listenAndServe(bind string, port int, handler http.Handler) error {
	addr := bind + ":" + strconv.Itoa(port)
	return http.ListenAndServe(addr, handler)
}

func reloadHandler(server *Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		server.reload()

		w.WriteHeader(http.StatusOK)

		io.WriteString(w, "reload")
	})
}
