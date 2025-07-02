package models

type Genre struct {
	Id       int    `db:"id"`
	Name     string `db:"name"`
	Slug     string `db:"slug"`
	Image    string `db:"image"`
	Position int    `db:"position"`
	Status   int    `db:"status"`
}