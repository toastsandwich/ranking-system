package server

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/toastsandwich/realtime-ranking-system/config"
	"github.com/toastsandwich/realtime-ranking-system/repository"
)

type ApiServer struct {
	Addr            string
	Router          *mux.Router
	ScoreRepository *repository.ScoreRepository
}

func NewApiServer(addr string) *ApiServer {
	return &ApiServer{
		Addr:            addr,
		Router:          mux.NewRouter(),
		ScoreRepository: repository.CreateScoreRepository(config.DB),
	}
}

func (s *ApiServer) Start() error {
	s.routes()
	server := http.Server{
		Addr:    s.Addr,
		Handler: s.Router,
	}
	log.Println("server started ", s.Addr)
	return server.ListenAndServe()
}

func (s *ApiServer) routes() {
	corsMiddleware := mux.CORSMethodMiddleware(s.Router)
	s.Router.Use(corsMiddleware, s.LoggerMiddleware)

	submitHandler := ConvertToStandarHandler(s.Submit)
	getRank := ConvertToStandarHandler(s.GetRank)
	listTopN := ConvertToStandarHandler(s.ListTopN)

	s.Router.HandleFunc("/submit", submitHandler).Methods("POST")
	s.Router.HandleFunc("/get-rank", getRank).Methods("GET")
	s.Router.HandleFunc("/list-top-n", listTopN).Methods("GET")
}
