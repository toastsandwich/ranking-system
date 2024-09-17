package server

import (
	"github.com/toastsandwich/realtime-ranking-system/model"
	"github.com/toastsandwich/realtime-ranking-system/utils"
)

func (s *ApiServer) SubmitScoreService(score utils.ScoreData) error {
	scr := model.Score{
		Username: score.Username,
		Country:  score.Country,
		State:    score.State,
		Score:    score.Score,
	}
	return s.ScoreRepository.SubmitUserScore(scr)
}

func (s *ApiServer) GetRankService(username string) (scr utils.ScoreData, err error) {
	score, err := s.ScoreRepository.GetRank(username)
	if err != nil {
		return
	}
	scr = utils.ScoreData{
		Username: score.Username,
		Country:  score.Country,
		State:    score.State,
		Score:    score.Score,
	}
	return
}

func (s *ApiServer) ListTopNService(n int, location string, grpBy string) (scrs []utils.ScoreData, err error) {
	scores, err := s.ScoreRepository.ListTopN(n, location, grpBy)
	if err != nil {
		return
	}
	for _, scr := range scores {
		score := utils.ScoreData{
			Username: scr.Username,
			Country:  scr.Country,
			State:    scr.State,
			Score:    scr.Score,
		}
		scrs = append(scrs, score)
	}
	return
}
