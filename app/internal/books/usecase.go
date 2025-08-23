package books

import (
	"context"
	"errors"
	"fmt"
)

var ErrIncorrectParameter = errors.New("incorrect parameter")

type UseCaseInt interface {
	GetBook(ctx context.Context, id int) (BookResponse, error)
	CreateBook(ctx context.Context, book BookDTO) (int, error)
	UpdateBook(ctx context.Context, book BookResponse) (int, error)
	RemoveBook(ctx context.Context, id int) (int, error)
}

type UseCase struct {
    repo BookRepository
}

func NewUseCase(repo BookRepository) UseCaseInt {
    return &UseCase{repo: repo}
}


func (u *UseCase) GetBook(ctx context.Context, id int) (BookResponse, error) {
	const op = "books.usecase.GetBook"
	if id == 0 {
		return BookResponse{}, fmt.Errorf("%s: %w", op, ErrIncorrectParameter)
	}
	book, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return BookResponse{}, fmt.Errorf("%s: %w", op, err)
	}
    return book, nil
}

func (u *UseCase) CreateBook(ctx context.Context, book BookDTO) (int, error) {
	const op = "books.usecase.CreateBook"
	if err := book.Validate(); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := u.repo.Create(ctx, book)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
    return id, nil
}

func (u *UseCase) UpdateBook(ctx context.Context, book BookResponse) (int, error) {
	const op = "books.usecase.UpdateBook"
	if err := book.Validate(); err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
	id, err := u.repo.Update(ctx, book)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
    return id, nil
}

func (u *UseCase) RemoveBook(ctx context.Context, id int) (int, error) {
	const op = "books.usecase.RemoveBook"
	if id == 0 {
		return 0, fmt.Errorf("%s: %w", op, ErrIncorrectParameter)
	}
	id, err := u.repo.Delete(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
    return id, nil
}