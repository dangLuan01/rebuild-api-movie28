package models

type Movie struct {
	Id           int     `db:"id"`
	Name         string  `db:"name"`
	Origin_name  string  `db:"origin_name"`
	Slug         string  `db:"slug"`
	Type         string  `db:"type"`
	Release_date int     `db:"release_date"`
	Rating       float32 `db:"rating"`
	Content      string  `db:"content,omitempty"`
	Runtime      string  `db:"runtime,omitempty"`
	Age          string  `db:"age,omitempty"`
	Trailer      string  `db:"trailer,omitempty"`
	Image        Image  
	Genres       []Genre
	Servers      []Server 
}

type Image struct {
	Poster string `db:"poster,omitempty"`
	Thumb  string `db:"thumb,omitempty"`
}

type Episode struct {
	Server_id int    `db:"server_id"`
	Episode   string `db:"episode"`
	Hls       string `db:"hls"`
}

type Server struct {
	Id       int       `db:"id"`
	Name     string    `db:"name"`
	Episodes []Episode `db:"episodes"`
}
