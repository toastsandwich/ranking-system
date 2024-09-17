package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/toastsandwich/realtime-ranking-system/utils"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func ConvertToStandarHandler(handler Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			url := r.URL.Path
			log.Printf("error : PATH = %s ERROR = %s", url, err.Error())
			errorResp := map[string]interface{}{
				"errorCode":    http.StatusInternalServerError,
				"errorMessage": err.Error(),
			}
			respJSON, err := json.Marshal(errorResp)
			if err != nil {
				log.Printf("error: %s", err.Error())
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(respJSON)
		}
	}
}

func (s *ApiServer) Submit(w http.ResponseWriter, r *http.Request) error {
	username := r.URL.Query().Get("username")
	country := r.URL.Query().Get("country")
	state := r.URL.Query().Get("state")
	score_str := r.URL.Query().Get("score")

	score, err := strconv.ParseFloat(score_str, 32)
	if err != nil {
		return err
	}
	scr := utils.ScoreData{
		Username: username,
		Country:  country,
		State:    state,
		Score:    score,
	}
	if err := s.SubmitScoreService(scr); err != nil {
		return err
	}
	return nil
}

func (s *ApiServer) GetRank(w http.ResponseWriter, r *http.Request) error {
	username := r.URL.Query().Get("username")

	score, err := s.GetRankService(username)
	if err != nil {
		return err
	}

	scoreJSON, err := json.Marshal(score)
	if err != nil {
		return err
	}
	log.Println(string(scoreJSON))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(scoreJSON)
	return nil
}

func (s *ApiServer) ListTopN(w http.ResponseWriter, r *http.Request) error {
	n_str := r.URL.Query().Get("n")
	location := r.URL.Query().Get("location")
	rankBy := r.URL.Query().Get("rank_by")
	n, err := strconv.Atoi(n_str)
	if err != nil {
		return err
	}
	scores, err := s.ListTopNService(n, location, rankBy)
	if err != nil {
		return err
	}
	scoresJSON, err := json.Marshal(scores)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(scoresJSON)
	return nil
}
