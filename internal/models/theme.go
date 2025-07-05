package models

type Theme struct {
	Id         int     `db:"id"`
	Name       string  `db:"name"`
	Genre_id   *int    `db:"genre_id,omitempty"`
	Country_id *int    `db:"country_id,omitempty"`
	Type       *string `db:"type,omitempty"`
	Year       *int    `db:"year,omitempty"`
	Limit      int     `db:"limit"`
	Layout     int     `db:"layout"`
}
