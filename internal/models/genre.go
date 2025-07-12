package models

type Genre struct {
	Id       int    `db:"id"`
	Name     string `db:"name,omitempty"`
	Slug     string `db:"slug,omitempty"`
	Image    string `db:"image,omitempty"`
	Position int    `db:"position,omitempty"`
	Status   int    `db:"status,omitempty"`
	Total    int    `db:"total,omitempty"`
}
type GenreWithMovie struct {
	Genre      Genre
	Movie      []Movie
	Page       int
	PageSize   int
	TotalPages int
}