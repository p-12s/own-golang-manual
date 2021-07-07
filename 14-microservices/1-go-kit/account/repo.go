package account

import (
	"database/sql"
	"errors"
	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("unable to handle Repo Request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) *repo {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}
