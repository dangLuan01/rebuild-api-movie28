package movierepository

import (
	"encoding/json"
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

func (mr *SqlMovieRepository) FindBySlug(slug string) (*v1dto.MovieDetailDTO, error) {
	var (
		movie_raw v1dto.MovieRawDTO
		genres []v1dto.GenreDTO
	)

	thumbSubquery := mr.db.From(goqu.T("movie_images").As("mi")).
		Where(
			goqu.I("mi.movie_id").Eq(goqu.I("m.id")),
			goqu.I("mi.is_thumbnail").Eq(1),
		).
		Select(goqu.Func("CONCAT", goqu.I("mi.path"), goqu.I("mi.image"))).
		Limit(1)
	ds := mr.db.Select(
		"m.id",
		"m.name",
		"m.origin_name",
		"m.slug",
		"m.type",
		"m.release_date",
		"m.rating",
		goqu.Func("IFNULL", goqu.I("m.content"), "").As("content"),
		goqu.Func("IFNULL", goqu.I("m.runtime"), "").As("runtime"),
		goqu.Func("IFNULL", goqu.I("m.age"), "").As("age"),
		goqu.Func("IFNULL", goqu.I("m.trailer"), "").As("trailer"),
		thumbSubquery.As("thumb"),
	).
	From(goqu.T("movies").As("m")).
	Where(goqu.Ex{"slug": slug})

	found, err := ds.ScanStruct(&movie_raw)
	if err != nil {
		
		return nil, fmt.Errorf("Faile fetch movie:%v", err)
	}
	if !found {

		return nil, fmt.Errorf("Not found movie")
	}

	ds_genres := mr.db.From("genres").
	Join(
		goqu.T("movie_genres").As("mg"),
		goqu.On(goqu.I("genres.id").Eq(goqu.I("mg.genre_id"))),
	).
	Where(goqu.Ex{"mg.movie_id": movie_raw.Id}).
	Select("genres.name", "genres.slug")
	if er := ds_genres.ScanStructs(&genres); er != nil {
		return nil, fmt.Errorf("Faile fetch genres:%v", er)
	}	
	movie := v1dto.MapMovieDetailDTO(movie_raw)

	movie.Genres = genres
	server, err := mr.FindServer(movie_raw.Id)

	if err != nil {
		return nil, fmt.Errorf("Faile fetch server:%v", err)
	}

	movie.Servers = server
	
	return movie, nil
}

func (mr *SqlMovieRepository) FindServer(id int) ([]v1dto.ServerDTO, error) {
	var servers []struct {
        Id       int             `db:"id"`
        Name     string          `db:"name"`
        Episodes json.RawMessage `db:"episodes"`
    }
    err := mr.db.From("movie_servers").
        LeftJoin(goqu.T("episodes"), 
            goqu.On(goqu.I("movie_servers.id").Eq(goqu.I("episodes.server_id")))).
        Where(goqu.Ex{
            "episodes.movie_id": uint(id),
        }).
        Select(
            "movie_servers.id",
            "movie_servers.name",
            goqu.L(`JSON_ARRAYAGG(
                CASE 
                    WHEN episodes.server_id IS NULL THEN NULL
                    ELSE JSON_OBJECT(
                        'server_id', episodes.server_id,
                        'episode', episodes.episode,
                        'hls', episodes.hls
                    )
                END
            )`).As("episodes"),
        ).
        GroupBy("movie_servers.id").
        ScanStructs(&servers)

    if err != nil {
        return nil, err
    }
    result := make([]v1dto.ServerDTO, 0, len(servers))
    for _, s := range servers {
        server := v1dto.ServerDTO{
            Id:   s.Id,
            Name: s.Name,
        }
        // Parse JSON episodes
        if len(s.Episodes) > 0 && string(s.Episodes) != "null" {
            var episodes []v1dto.EpisodeDTO
            if err := json.Unmarshal(s.Episodes, &episodes); err != nil {
                return nil, fmt.Errorf("failed to unmarshal episodes: %v", err)
            }
            server.Episodes = episodes
        }

        result = append(result, server)
    }

    return result, nil
}