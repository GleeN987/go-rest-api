package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch the comment by id")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Interface storing methods which repository implements
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Struct that will handle business logic
type Service struct {
	Store Store
}

type Comment struct {
	ID     string
	Slug   string //PATH to where comment in placed ex. post/1
	Body   string
	Author string
}

// Composite literal because Go doesn't have constructors
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Getting comment")
	comment, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return comment, nil
}

func (s *Service) UpdateComment(ctx context.Context, comment Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, comment Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
