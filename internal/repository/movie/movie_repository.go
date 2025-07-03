package movierepository

import (
	"fmt"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/doug-martin/goqu/v9"
)

type SqlMovieRepository struct {
	db *goqu.Database
}

func NewSqlMovieRepository(DB *goqu.Database) MovieRepository {
	return &SqlMovieRepository {
		db: DB,
	}
}

func (mr *SqlMovieRepository) FindByHot(limit int) ([]v1dto.MovieRawDTO, error) {
	var movies []v1dto.MovieRawDTO
	
	thumbSubquery := mr.db.From(goqu.T("movie_images").As("mi")).
		Where(
			goqu.I("mi.movie_id").Eq(goqu.I("m.id")),
			goqu.I("mi.is_thumbnail").Eq(1),
		).
		Select(
			goqu.Func("CONCAT", goqu.I("mi.path"), goqu.I("mi.image")),
		).
		Limit(1)

	genreSubquery := mr.db.From(goqu.T("genres").As("g")).
		Join(
			goqu.T("movie_genres").As("mg"), goqu.On(goqu.I("g.id").Eq(goqu.I("mg.genre_id"))),
		).
		Where(
			goqu.I("mg.movie_id").Eq(goqu.I("m.id")),
		).
		Select(goqu.I("g.name")).
		Limit(1)

	ds := mr.db.From(goqu.T("movies").As("m")).
	Where(goqu.Ex{
		"m.status": 1,
		"m.hot": 1,
	}).
	Select(
		goqu.I("m.name"),
		goqu.I("m.origin_name"),
		goqu.I("m.slug"),
		goqu.I("m.type"),
		goqu.I("m.release_date"),
		goqu.I("m.rating"),
		thumbSubquery.As("thumb"),
		genreSubquery.As("genre"),
	).
	Order(goqu.I("m.updated_at").Desc()).Limit(uint(limit))
	
	if err := ds.ScanStructs(&movies); err != nil {
		return nil, fmt.Errorf("Faile scantrucs movies:%v", err)
	}
	//m := v1dto.MapMovieRawToMovieDTO(movies)

	return movies, nil
}

func (mr *SqlMovieRepository) FindAll(page, pageSize int) ([]v1dto.MovieRawDTO, v1dto.Paginate, error) {

	var movies []v1dto.MovieRawDTO

	posterSubquery := mr.db.From(goqu.T("movie_images").As("mi")).
		Where(
			goqu.I("mi.movie_id").Eq(goqu.I("m.id")),
			goqu.I("mi.is_thumbnail").Eq(0),
		).
		Select(
			goqu.Func("CONCAT", goqu.I("mi.path"), goqu.I("mi.image")),
		).Limit(1)

	genreSubquery := mr.db.From(goqu.T("genres").As("g")).
		Join(
			goqu.T("movie_genres").As("mg"), goqu.On(goqu.I("g.id").Eq(goqu.I("mg.genre_id"))),
		).
		Where(
			goqu.I("mg.movie_id").Eq(goqu.I("m.id")),
		).
		Select(goqu.I("g.name")).Limit(1)

	ds := mr.db.From(goqu.T("movies").As("m")).
	Where(
		goqu.I("m.status").Eq(1),
	).
	Select(
		goqu.I("m.name"),
		goqu.I("m.origin_name"),
		goqu.I("m.slug"),
		goqu.I("m.type"),
		goqu.I("m.release_date"),
		goqu.I("m.rating"),
		posterSubquery.As("poster"),
		genreSubquery.As("genre"),
	).
	Order(goqu.I("m.updated_at").Desc())
	count, err := ds.Count()
	if err != nil {
		return nil, v1dto.Paginate{} ,fmt.Errorf("Faile count total movies:%v", err)
	}
	totalPages := count/int64(pageSize)
	if totalPages == 0 {
		totalPages = 1
	}
	if err := ds.Limit(uint(pageSize)).Offset(uint((page - 1) * pageSize)).ScanStructs(&movies); err != nil {
		return nil, v1dto.Paginate{} ,fmt.Errorf("Faile scantructs movies:%v", err)
	}
	
	return movies, v1dto.Paginate{
		Page: page,
		PageSize: pageSize,
		TotalPages: totalPages,
	}, nil
}