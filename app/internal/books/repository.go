package books

import (
	"github.com/PoliakovIvan2606/Fiber/pkg/db/postgres"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type BookRepository interface {
    GetByID(ctx context.Context, id int) (BookResponse, error)
    Create(ctx context.Context, book BookDTO) (int, error)
    Update(ctx context.Context, book BookResponse) (int, error)
    Delete(ctx context.Context, id int) (int, error)
}


type RepositoryBook struct {
	postgres.Client
}

func NewBookRepository(pstg postgres.Client) *RepositoryBook {
    return &RepositoryBook{pstg}
}

func (r *RepositoryBook) GetByID(ctx context.Context, id int) (BookResponse, error) {
	const op = "books.repository.GetByID"
	query := `SELECT ID, name, description, author, number_pages FROM book WHERE id = $1`

	var book BookResponse

	err := r.QueryRow(ctx, query, id).Scan(
		&book.ID, &book.Name, &book.Description, 
		&book.Author, &book.NumberPages,
	)
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return BookResponse{}, fmt.Errorf("%s: book with id %d not found", op, id)
        }
        return BookResponse{}, fmt.Errorf("%s: failed to get book by id: %w", op, err)
    }
	return book, nil
}

func (r *RepositoryBook) Create(ctx context.Context, book BookDTO) (int, error) {
	const op = "books.repository.Create"
	query := `INSERT INTO book (name, description, author, number_pages) VALUES ($1, $2, $3, $4) RETURNING id`

	var id int
	err := r.QueryRow(ctx, query, 
		book.Name, book.Description, book.Author, book.NumberPages,
		).Scan(&id)
    if err != nil {
        // Проверяем на нарушение уникальности (email уже существует)
        if pgxErr, ok := err.(*pgconn.PgError); ok && pgxErr.Code == "23505" {
            return 0, fmt.Errorf("%s: book with id %s already exists", op, book.Name)
        }
        return 0, fmt.Errorf("%s: failed to add book: %w", op, err)
    }
	return id, nil
}

func (r *RepositoryBook) Update(ctx context.Context, book BookResponse) (int, error) {
	const op = "books.repository.Update"
	query := `UPDATE book SET name = $1, description = $2, author = $3, number_pages = $4 WHERE id = $5 RETURNING id`
	var updatedID int
	err := r.QueryRow(ctx, query, book.Name, book.Description, book.Author, book.NumberPages, book.ID).Scan(&updatedID)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to update book: %w", op, err)
	}
	return updatedID, nil
}

func (r *RepositoryBook) Delete(ctx context.Context, id int) (int, error) {
	const op = "books.repository.Delete"
	query := `DELETE FROM book WHERE id = $1 RETURNING id`

	var DeletedID int
	err := r.QueryRow(ctx, query, id).Scan(&DeletedID)
	if err != nil {
		return 0, fmt.Errorf("%s: failed to delete book: %w", op, err)
	}
	return DeletedID, nil
}