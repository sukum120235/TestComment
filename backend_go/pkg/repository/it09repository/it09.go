package it09repository

import (
	"sync"
	"test_project/pkg/model/it09model"
	"time"

	"gorm.io/gorm"
)

type ICommentRepository interface {
	GetComments() (*[]it09model.Comment, error)
	InsertComment(comment it09model.Comment) error
}

type CommentRepository struct{}

var (
	commentStore   = make([]it09model.Comment, 0)
	commentStoreMu sync.RWMutex
)

func NewCommentRepository(_ *gorm.DB) ICommentRepository {
	return &CommentRepository{}
}

func (r *CommentRepository) GetComments() (*[]it09model.Comment, error) {
	commentStoreMu.RLock()
	defer commentStoreMu.RUnlock()

	comments := append([]it09model.Comment(nil), commentStore...)
	return &comments, nil
}

func (r *CommentRepository) InsertComment(comment it09model.Comment) error {
	commentStoreMu.Lock()
	defer commentStoreMu.Unlock()

	if comment.CreatedAt.IsZero() {
		comment.CreatedAt = time.Now()
	}

	commentStore = append(commentStore, comment)
	return nil
}
