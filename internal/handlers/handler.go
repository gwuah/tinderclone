package handlers

import (
	"github.com/gwuah/tinderclone/internal/lib"
	"github.com/gwuah/tinderclone/internal/queue"
	"github.com/gwuah/tinderclone/internal/repository"
)

type Handler struct {
	repo *repository.Repository
	sms  *lib.Termii
	q    *queue.Que
}

func New(repo *repository.Repository, sms *lib.Termii, q *queue.Que) *Handler {
	return &Handler{
		repo: repo,
		sms:  sms,
		q:    q,
	}
}
