package models

type Category struct {
	Id     int    `db:"id"`
	Name   string `db:"name"`
	Slug   string `db:"slug"`
	Status int    `db:"status"`
}