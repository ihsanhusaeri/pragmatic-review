package service

import (
	"errors"
	"math/rand"

	"github.com/ihsan/pragmatic-review/entity"
	"github.com/ihsan/pragmatic-review/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type postService struct {
	repo repository.PostRepository
}

func NewPostService(r repository.PostRepository) PostService {
	return &postService{repo: r}
}
func (ps *postService) Validate(post *entity.Post) error {
	if post == nil {
		return errors.New("The post is empty")
	}
	if post.Title == "" {
		return errors.New("The title is empty")
	}
	return nil
}

func (ps *postService) Save(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return ps.repo.Save(post)
}

func (ps *postService) FindAll() ([]entity.Post, error) {
	return ps.repo.FindAll()
}
