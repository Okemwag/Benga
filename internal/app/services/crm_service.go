package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"encoding/json"
	"github.com/gorilla/mux"
)

// CrmService is a service that provides CRM operations

type CrmService struct {
	logger *log.Logger
	
}

// NewCrmService creates a new crm service

func NewCrmService(logger *log.Logger) *CrmService {
	return &CrmService{logger: logger}
}

// Start starts the crm service

func (s *CrmService) Start() error {
	s.logger.Println("Starting crm service...")
	router := http.NewServeMux()
	router.HandleFunc("/crm", s.handleCrm)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			s.logger.Fatal(err)
		}
	}()
	s.logger.Println("Crm service started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	s.logger.Println("Shutting down crm service...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return server.Shutdown(ctx)
}