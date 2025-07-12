package genrerepository

import (
	"fmt"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/models"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type SqlGenreRepository struct {
	//genres []models.Genre
	db    *goqu.Database
}

func NewSqlGenreRepository(DB *goqu.Database) GenreRepository {
	return &SqlGenreRepository{
		//genres: make([]models.Genre, 0),
		db:    DB,
	}
}

func (g *SqlGenreRepository) FindAll() ([]models.Genre, error) {
	
	ds := g.db.From(goqu.T("genres").As("g")).
	InnerJoin(
		goqu.T("movie_genres").As("mg"), goqu.On(goqu.I("g.id").Eq(goqu.I("mg.genre_id"))),
	).
	Where(
		goqu.I("g.status").Eq(1),
	).
	GroupBy(goqu.I("g.id")).
    Order(goqu.I("g.position").Asc()).
	Select(
		goqu.I("g.id"),
		goqu.I("g.name"),
		goqu.I("g.slug"),
		goqu.I("g.image"),
		goqu.I("g.position"),
		goqu.I("g.status"),
		goqu.COUNT(goqu.I("mg.movie_id")).As("total"),
	)
	var genres []models.Genre
	if err := ds.ScanStructs(&genres); err != nil {
		return nil, fmt.Errorf("Faile get all genres:%v", err)
	}

	return genres, nil
}

func (g *SqlGenreRepository)FindBySlug(slug string, page, pageSize int64) (models.GenreWithMovie, error)  {
	var (
		listMovie []models.Movie
		genreInfo models.Genre
		movie []v1dto.MovieRawDTO
	)

	queryGenre := g.db.From("genres").Where(
		goqu.C("slug").Eq(slug),
	).Select(goqu.I("id"), goqu.I("name"))
	
	found, err := queryGenre.ScanStruct(&genreInfo)

	if err != nil || !found {
		return models.GenreWithMovie{}, fmt.Errorf("Faile get genre:%v", err)
	}
	
	posterSubquery := g.db.From(goqu.T("movie_images").As("mi")).
		Where(
			goqu.I("mi.movie_id").Eq(goqu.I("m.id")),
			goqu.I("mi.is_thumbnail").Eq(0),
		).
		Select(goqu.Func("CONCAT", goqu.I("mi.path"), goqu.I("mi.image"))).
		Limit(1)
	queryMovie := g.db.From(goqu.T("movies").As("m")).
	LeftJoin(
		goqu.T("movie_genres").As("mg"),
		goqu.On(goqu.I("m.id").Eq(goqu.I("mg.movie_id"))),
	).
	Where(
		goqu.Ex{"mg.genre_id":genreInfo.Id},
	).
	Select(
		goqu.I("m.name"),
		goqu.I("m.origin_name"),
		goqu.I("m.slug"),
		goqu.I("m.type"),
		goqu.I("m.release_date"),
		goqu.I("m.rating"),
		posterSubquery.As("poster"),
	).
	Order(goqu.I("m.updated_at").Desc())
	
	totalSize, err := queryMovie.Count()

	if err != nil {
		return models.GenreWithMovie{}, fmt.Errorf("Faile count total movies:%v", err)
	}
	
	if err := queryMovie.Limit(uint(pageSize)).Offset(uint((page - 1) * pageSize)).ScanStructs(&movie); err != nil {
		return models.GenreWithMovie{}, fmt.Errorf("Faile scantrucs movies:%v", err)
	}
	for _, item := range movie {
		listMovie = append(listMovie, models.Movie{
			Name: item.Name,
			Origin_name: item.Origin_name,
			Slug: item.Slug,
			Type: item.Type,
			Release_date: item.Release_date,
			Rating: utils.ConvertRating(float32(item.Rating)),
			Image: models.Image{
				Poster: item.Poster,
			},
			Genres: []models.Genre{
				{
					Name: genreInfo.Name,
				},
			},
		})
	}
	return models.GenreWithMovie{
		Movie: listMovie,
		Genre: genreInfo,
		Page: page,
		PageSize: pageSize,
		TotalPages: utils.TotalPages(totalSize, pageSize),
	}, nil
}