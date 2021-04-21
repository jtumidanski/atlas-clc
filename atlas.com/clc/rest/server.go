package rest

import (
	"atlas-clc/rest/resources"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

type Server struct {
	l  logrus.FieldLogger
	hs *http.Server
}

func NewServer(l *logrus.Logger) *Server {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/ms/clc").Subrouter()
	router.Use(commonHeader)

	s := resources.NewSessionResource(l)
	sRouter := router.PathPrefix("/sessions").Subrouter()
	sRouter.HandleFunc("", s.GetSessions)
	sRouter.HandleFunc("/{sessionId}/errors/{errorId}", s.LoginError)

	w := l.Writer()
	defer w.Close()

	hs := http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     log.New(w, "", 0), // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
	return &Server{l, &hs}
}

func (s *Server) Run() {
	s.l.Infof("Starting server on port 8080.")
	err := s.hs.ListenAndServe()
	if err != nil {
		s.l.WithError(err).Fatalf("Error starting server.")
	}
}

func commonHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
