package models

type Genre struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Slug     string `db:"slug"`
	Image    string `db:"image"`
	Position int    `db:"position"`
	Status   int    `db:"status"`
	Total    int    `db:"total"`
}
type GenreWithMovie struct {
	Genre      Genre
	Movie      []Movie
	Page       int
	PageSize   int
	TotalPages int
}