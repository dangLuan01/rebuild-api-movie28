package v1dto

type MovieDTO struct {
	Name         string   `json:"name"`
	Origin_name  string   `json:"origin_name"`
	Slug         string   `json:"slug"`
	Type         string   `json:"type"`
	Release_date int      `json:"release_date"`
	Rating       float32  `json:"rating"`
	Content      string   `json:"content,omitempty"`
	Runtime      string   `json:"runtime,omitempty"`
	Age          string   `json:"age,omitempty"`
	Trailer      string   `json:"trailer,omitempty"`
	Image        ImageDTO `json:"image"`
	Genres       []GenreDTO `json:"genres"`
}

type ImageDTO struct {
	Thumb  string `json:"thumb,omitempty"`
	Poster string `json:"poster,omitempty"`
}