package repository

import (
	"github.com/ihsan/pragmatic-review/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
