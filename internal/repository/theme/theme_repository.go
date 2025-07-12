package themerepository

import (
	"fmt"
	"log"
	"sync"

	v1dto "github.com/dangLuan01/rebuild-api-movie28/internal/dto/v1"
	"github.com/dangLuan01/rebuild-api-movie28/internal/utils"
	"github.com/doug-martin/goqu/v9"
)

type SqlThemeRepository struct {
	db *goqu.Database
}

func NewSqlThemeRepository(DB *goqu.Database) ThemeRepository {
	return &SqlThemeRepository{
		db: DB,
	}
}
func (tr *SqlThemeRepository) buildMovieQueryFromTheme(theme v1dto.ThemeDTO) *goqu.SelectDataset {
    // Subquery cho genre_name
    genreSubquery := tr.db.From(goqu.T("genres").As("g")).
        Join(goqu.T("movie_genres").As("mg"), goqu.On(goqu.I("g.id").Eq(goqu.I("mg.genre_id")))).
        Where(goqu.I("mg.movie_id").Eq(goqu.I("m.id"))).
        Select(goqu.I("g.name")).
        Limit(1)

    // Subquery cho poster
    posterSubquery := tr.db.From(goqu.T("movie_images").As("mi")).
        Where(
            goqu.I("mi.movie_id").Eq(goqu.I("m.id")),
            goqu.I("mi.is_thumbnail").Eq(0),
        ).
        Select(goqu.Func("CONCAT", goqu.I("mi.path"), goqu.I("mi.image"))).
        Limit(1)

    // Truy vấn chính
    query := tr.db.From(goqu.T("movies").As("m")).
        Select(
            goqu.I("m.id"),
            goqu.I("m.name"),
            goqu.I("m.origin_name"),
            goqu.I("m.slug"),
            goqu.I("m.type"),
            goqu.I("m.release_date"),
            goqu.I("m.rating"),
            posterSubquery.As("poster"),
            genreSubquery.As("genre"),
        ).
        Where(goqu.I("m.hot").Eq(0))
    // Điều kiện genre_id
    if theme.Genre_id != nil {
        query = query.Join(
            goqu.T("movie_genres").As("mg"),
            goqu.On(goqu.I("mg.movie_id").Eq(goqu.I("m.id"))),
        ).Where(goqu.I("mg.genre_id").Eq(*theme.Genre_id))
    }

    // Điều kiện country_id
    if theme.Country_id != nil {
        query = query.Join(
            goqu.T("movie_countries").As("mc"),
            goqu.On(goqu.I("mc.movie_id").Eq(goqu.I("m.id"))),
        ).Where(goqu.I("mc.country_id").Eq(*theme.Country_id))
    }

    // Điều kiện type
    if theme.Type != nil {
        query = query.Where(goqu.I("m.type").Eq(*theme.Type))
    }

    // Điều kiện year
    if theme.Year != nil {
       query = query.Where(goqu.I("m.release_year").Eq(*theme.Year))
    }
    query = query.Order(goqu.I("m.updated_at").Desc())
    return query
}

func (tr *SqlThemeRepository) FindMoviesByTheme(theme v1dto.ThemeDTO, page, pageSize int64) (*v1dto.MoviesDTOWithPaginate, error) {
   
    offset      := (page - 1) * pageSize
    baseQuery   := tr.buildMovieQueryFromTheme(theme)
    totalSize, err := baseQuery.Count()
    if err != nil {
        return nil, fmt.Errorf("Failed to count movies: %v", err)
    }
    var movies []v1dto.MovieRawDTO
    err = baseQuery.Offset(uint(offset)).
        Limit(uint(pageSize)).
        ScanStructs(&movies)
    
    if err != nil {
        return nil, fmt.Errorf("Failed to scant movies: %v", err)
    }
    
	movie := v1dto.MapMovieDTOWithPanigate(movies, v1dto.Paginate{
		Page: page,
		PageSize: pageSize,
		TotalPages: utils.TotalPages(totalSize, pageSize),	
	})

    return movie, nil
}
func (tr *SqlThemeRepository) FindAll(id, pageTheme, pageMovie, pageSize int64) (*v1dto.ThemesWithPaginateDTO, error) {
	
    offset := (pageTheme - 1) * pageSize

    // Lấy danh sách theme
    var themes []v1dto.ThemeDTO
    ds := tr.db.From(goqu.T("themes").As("t")).
	Where(
		goqu.I("t.status").Eq(1),
	)
    if id != 0 {
        ds = ds.Where(
			goqu.I("t.id").Eq(id),
		)
    }
    totalSize, err := ds.Count()
    if err != nil {
        return nil, fmt.Errorf("Failed to count themes: %v", err)
    }
    if err := ds.Order(goqu.I("t.priority").Asc()).Offset(uint(offset)).Limit(uint(pageSize)).ScanStructs(&themes); err != nil {
        return nil, fmt.Errorf("Failed to scan themes: %v", err)
    }

    // Kênh để thu thập kết quả và lỗi
    type resultStruct struct {
        index           int
        themeWithMovies v1dto.ThemeWithMovieDTO
        err             error
    }
    resultsChan := make(chan resultStruct, len(themes))
    var wg sync.WaitGroup

    // Chạy goroutines cho mỗi theme
    // start := time.Now()
    for i, theme := range themes {
        wg.Add(1)
        go func(idx int, t v1dto.ThemeDTO) {
            defer wg.Done()
            movies, err := tr.FindMoviesByTheme(t, pageMovie, t.Limit)
            if err != nil {
                resultsChan <- resultStruct{
                    index: idx, 
                    err: fmt.Errorf("failed to get movies for theme %s: %v", t.Name, err),
                }
                return
            }
            resultsChan <- resultStruct{
                index: idx,
                themeWithMovies: v1dto.ThemeWithMovieDTO{
                    Theme:       t,
                    Movies: *movies,
                },
            }
        }(i, theme)
    }

    // Đóng kênh sau khi tất cả goroutines hoàn thành
    go func() {
        wg.Wait()
        close(resultsChan)
    }()

    // Thu thập kết quả theo thứ tự
    resultSlice := make([]v1dto.ThemeWithMovieDTO, len(themes))
    var errors []error
    for res := range resultsChan {
        if res.err != nil {
            errors = append(errors, res.err)
            continue
        }
        resultSlice[res.index] = res.themeWithMovies
    }

    // Kiểm tra lỗi
    if len(errors) > 0 {
        log.Printf("Encountered %d errors while fetching movies: %v", len(errors), errors)
        // Nếu tất cả theme đều lỗi, trả về lỗi
        if len(errors) == len(themes) {
            return nil, fmt.Errorf("Failed to fetch movies for all themes: %v", errors)
        }
    }

    // Loại bỏ các theme rỗng (nếu có)
    result := make([]v1dto.ThemeWithMovieDTO, 0, len(themes))
    for _, r := range resultSlice {
        if r.Theme.Id != 0 { // Chỉ thêm các theme có dữ liệu hợp lệ
            result = append(result, r)
        }
    }
    t := v1dto.MapThemeDTOWithPaginate(result, v1dto.Paginate{
        Page:            pageTheme,
        PageSize:        pageSize,
        TotalPages:      utils.TotalPages(totalSize, pageSize),
	})
	
    return t, nil
}