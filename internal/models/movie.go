package models

type Movie struct {
	Id           int     `db:"id"`
	Name         string  `db:"name"`
	Origin_name  string  `db:"origin_name"`
	Slug         string  `db:"slug"`
	Type         string  `db:"type"`
	Release_date int     `db:"release_date"`
	Rating       float32 `db:"rating,omitempty"`
	Content      string  `db:"content,omitempty"`
	Runtime      string  `db:"runtime"`
	Age          string  `db:"age,omitempty"`
	Trailer      string  `db:"trailer,omitempty"`
	Image        Image   `db:"image"`
	Genres       []Genre 
}

type Image struct {
	Poster string `db:"poster,omitempty"`
	Thumb  string `db:"thumb,omitempty"`
}