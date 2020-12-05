package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

// Run starts up the server
func Run() (err error) {
	s := newServer()

	// Start that server!
	port := ":3030"
	fmt.Printf("Now listening on %s...\n", port)
	http.ListenAndServe(port, s.router)

	return err
}

type server struct {
	logger *log.Logger
	router *mux.Router
}

// newServer returns a server with logging, database, and routing
func newServer() *server {
	s := server{
		router: mux.NewRouter(),
	}

	logger := log.New()
	logger.Out = os.Stdout
	logger.SetFormatter(&log.TextFormatter{PadLevelText: true})
	s.logger = logger

	s.routes()

	return &s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) logMsg(msg string) {
	s.logger.Info(msg)
}

func (s *server) logErr(msg string, err error) {
	s.logger.Error(fmt.Sprintf("%s\n\t%s\n", msg, err))
}

func (s *server) panic(msg string, err error) {
	s.logger.Panic(fmt.Sprintf("%s\n\t%s\n", msg, err))
}
