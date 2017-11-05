package models

import (
	"encoding/json"
	"github.com/go-sql-driver/mysql"
	"time"
)

type Category struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt mysql.NullTime `json:"updatedAt"`
}

func (c *Category) MarshalJSON() ([]byte, error) {
	// TODO: Find a better way to set updatedAt to nil
	if !c.UpdatedAt.Valid {
		return json.Marshal(struct {
			ID        int             `json:"id"`
			Name      string          `json:"name"`
			CreatedAt time.Time       `json:"createdAt"`
			UpdatedAt *mysql.NullTime `json:"updatedAt"`
		}{c.ID, c.Name, c.CreatedAt, nil})
	}

	return json.Marshal(struct {
		ID        int       `json:"id"`
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}{c.ID, c.Name, c.CreatedAt, c.UpdatedAt.Time})
}
