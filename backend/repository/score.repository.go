package repository

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/toastsandwich/realtime-ranking-system/model"
)

type ScoreRepository struct {
	DB *sql.DB
}

func CreateScoreRepository(db *sql.DB) *ScoreRepository {
	return &ScoreRepository{
		DB: db,
	}
}

func (s *ScoreRepository) SubmitUserScore(score model.Score) error {
	query := `insert into scores(username, country, state, score) values(?, ?, ?, ?)`
	_, err := s.DB.Exec(query, score.Username, score.Country, score.State, score.Score)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			if mysqlErr.Number == 1062 {
				return fmt.Errorf("duplicate entry found")
			} else {
				return fmt.Errorf("error : %s", mysqlErr.Message)
			}
		} else {
			return err
		}
	}
	return nil
}

func (s *ScoreRepository) GetRank(username string) (model.Score, error) {
	query := `select * from scores where username = ?`
	row := s.DB.QueryRow(query, username)
	score := model.Score{}
	err := row.Scan(&score.Id, &score.Username, &score.Country, &score.State, &score.Score)
	if err == sql.ErrNoRows {
		return model.Score{}, fmt.Errorf("Empty rows")
	}
	return score, nil
}

func (s *ScoreRepository) ListTopN(n int, location string, grpBy string) ([]model.Score, error) {
	var query string
	var params []any
	
	switch grpBy {
		case "country":
		query = `SELECT * FROM scores WHERE country = ? ORDER BY score DESC LIMIT ?`
		params = append(params, location, n)
		case "state":
		query = `SELECT * FROM scores WHERE state = ? ORDER BY score DESC LIMIT ?`
		params = append(params, location, n)
		default:
		query = `SELECT * FROM scores ORDER BY score DESC LIMIT ?`
		params = append(params, n)
	}
	
	rows, err := s.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	scores := []model.Score{}
	for rows.Next() {
		score := model.Score{}
		rows.Scan(&score.Id, &score.Username, &score.Country, &score.State, &score.Score)
		scores = append(scores, score)
	}
	return scores, nil
}
