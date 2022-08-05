package film_postgres

import (
	"context"
	"testing"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"gotest.tools/assert"

	"films-api/internal/api/domain/film"
	filmPostgres "films-api/internal/api/repository/film_postgres"
	"films-api/pkg/log"
)

func TestRepository_GetByName(t *testing.T) {
	logger := log.New()

	DB, err := sqlx.Connect("pgx", "host=localhost port=5432 user=postgres dbname=dvdrental password=postgres sslmode=disable")
	if err != nil {
		logger.Fatalf("test init db failed: %s", err)
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		args    args
		want    film.FilmList
		wantErr bool
	}{
		{
			name: "success get film by name",
			args: args{
				ctx:  context.Background(),
				name: "Academy Dinosaur",
			},
			want: film.FilmList{
				{
					ID:              1,
					Name:            "Academy Dinosaur",
					Description:     "A Epic Drama of a Feminist And a Mad Scientist who must Battle a Teacher in The Canadian Rockies",
					ReleaseYear:     2006,
					Language:        "English             ",
					RentalDuration:  6,
					RentalRate:      0.99,
					Length:          86,
					ReplacementCost: 20.99,
					Rating:          "PG",
					SpecialFeatures: []string{"Deleted Scenes", "Behind the Scenes"},
					LastUpdate:      time.Unix(1369579858, 951000000).UTC(),
					Actors: []film.Actor{
						{ID: 1, Name: "Penelope", LastName: "Guiness"},
						{ID: 10, Name: "Christian", LastName: "Gable"},
						{ID: 20, Name: "Lucille", LastName: "Tracy"},
						{ID: 30, Name: "Sandra", LastName: "Peck"},
						{ID: 40, Name: "Johnny", LastName: "Cage"},
						{ID: 53, Name: "Mena", LastName: "Temple"},
						{ID: 108, Name: "Warren", LastName: "Nolte"},
						{ID: 162, Name: "Oprah", LastName: "Kilmer"},
						{ID: 188, Name: "Rock", LastName: "Dukakis"},
						{ID: 198, Name: "Mary", LastName: "Keitel"},
					},
					Categories: []film.Category{
						{ID: 6, Name: "Documentary"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := filmPostgres.NewRepository(time.Second*100, DB)

			got, err := r.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.DeepEqual(t, got, tt.want)
		})
	}
}
