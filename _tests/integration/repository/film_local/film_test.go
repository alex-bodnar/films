package film_redis

import (
	"context"
	"testing"
	"time"

	"gotest.tools/assert"

	"films-api/internal/api/domain/film"
	filmLocal "films-api/internal/api/repository/film_local"
	"films-api/internal/config"
)

func TestRepository_SetByName(t *testing.T) {
	type args struct {
		ctx  context.Context
		name string
		data film.FilmList
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:  context.Background(),
				name: "academy dinosaur",
				data: film.FilmList{
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
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := filmLocal.NewRepository(config.LocalCache{
				TimeLive:        time.Second * 10,
				NumberOfRecords: 10,
			})

			if err := r.SetByName(tt.args.ctx, tt.args.name, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("Repository.SetByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_GetByName(t *testing.T) {
	type args struct {
		ctx      context.Context
		name     string
		waitTime time.Duration
		liveTime time.Duration
	}
	tests := []struct {
		name    string
		args    args
		want    film.FilmList
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				name:     "academy dinosaur",
				waitTime: time.Second * 1,
				liveTime: time.Second * 20,
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
		{
			name: "fail timeout clean",
			args: args{
				ctx:      context.Background(),
				name:     "academy dinosaur",
				waitTime: time.Second * 2,
				liveTime: time.Second * 1,
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := filmLocal.NewRepository(config.LocalCache{
				TimeLive:        tt.args.liveTime,
				NumberOfRecords: 10,
			})

			if err := r.SetByName(tt.args.ctx, tt.args.name, tt.want); err != nil {
				t.Errorf("Repository.SetByName() error = %v", err)
			}

			time.Sleep(tt.args.waitTime)

			got, err := r.GetByName(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				assert.DeepEqual(t, got, tt.want)
			}
		})
	}
}
