package books

import (
	"errors"
	"regexp"
)

var (
	ErrNotValidate = errors.New("not validate fields")
	re             = regexp.MustCompile(`^[А-ЯЁ][а-яё]+(?:-[А-ЯЁ][а-яё]+)?\s+[А-ЯЁ]\.\s?[А-ЯЁ]\.$`)
)

type BookDTO struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Author string `json:"author"`
	NumberPages int `json:"number_pages"`
}

type BookResponse struct {
	ID int `json:"id"`
	BookDTO
}

func (b *BookDTO) Validate() error {
	if b.Name == "" || b.Description == "" || b.Author == "" || b.NumberPages <= 0 {
		return ErrNotValidate
	}

	if !re.MatchString(b.Author) {
		return ErrNotValidate
	}

	return nil
}