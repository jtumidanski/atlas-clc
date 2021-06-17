package rest

import (
	"atlas-clc/session"
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

	sRouter := router.PathPrefix("/sessions").Subrouter()
	sRouter.HandleFunc("", session.HandleGetSessions(l)).Methods(http.MethodGet)
	sRouter.HandleFunc("/{sessionId}/errors/{errorId}", session.HandleLoginError(l)).Methods(http.MethodGet)

	return router
}
