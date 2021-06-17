package rest

import (
	"atlas-clc/rest/resources"
	"context"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"sync"
)

func CreateRestService(l *logrus.Logger, ctx context.Context, wg *sync.WaitGroup) {
	go NewServer(l, ctx, wg, ProduceRoutes)
}

func ProduceRoutes(l logrus.FieldLogger) http.Handler {
	router := mux.NewRouter().PathPrefix("/ms/clc").Subrouter().StrictSlash(true)
	router.Use(CommonHeader)
	s := resources.NewSessionResource(l)
	sRouter := router.PathPrefix("/sessions").Subrouter()
	sRouter.HandleFunc("", s.GetSessions)
	sRouter.HandleFunc("/{sessionId}/errors/{errorId}", s.LoginError)

	return router
}
