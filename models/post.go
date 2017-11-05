package models

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Post struct {
	ID        int            `json:"id"`
	Title     string         `json:"title"`
	Slug      string         `json:"slug"`
	Body      string         `json:"body"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt mysql.NullTime `json:"updatedAt"`
	Author    *User          `json:"author"`
	Category  *Category      `json:"category"`
}

func (p *Post) MarshalJSON() ([]byte, error) {
	// TODO: Find a better way to set updatedAt to nil
	if !p.UpdatedAt.Valid {
		return json.Marshal(struct {
			ID        int             `json:"id"`
			Title     string          `json:"title"`
			Slug      string          `json:"slug"`
			Body      string          `json:"body"`
			CreatedAt time.Time       `json:"createdAt"`
			UpdatedAt *mysql.NullTime `json:"updatedAt"`
			Author    *User           `json:"author"`
			Category  *Category       `json:"category"`
		}{p.ID,
			p.Title,
			p.Slug,
			p.Body,
			p.CreatedAt,
			nil,
			p.Author,
			p.Category,
		})
	}

	return json.Marshal(struct {
		ID        int       `json:"id"`
		Title     string    `json:"title"`
		Slug      string    `json:"slug"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		Author    *User     `json:"author"`
		Category  *Category `json:"category"`
	}{p.ID,
		p.Title,
		p.Slug,
		p.Body,
		p.CreatedAt,
		p.UpdatedAt.Time,
		p.Author,
		p.Category,
	})
}
