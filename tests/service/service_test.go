package service_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"testing"
	"time"

	"films-api/pkg/log"

	_ "github.com/jackc/pgx/stdlib"

	"films-api/internal/api/delivery/http/statistics"
	"films-api/internal/config"
	"films-api/pkg/database"
)

func TestService(t *testing.T) {
	logger := log.New()
	ctx := context.Background()

	config, err := config.New("films-api", "../../volume/config.yaml")
	if err != nil {
		logger.Fatal(err)
	}

	db, err := database.New(config.Storage.Postgres, logger)
	if err != nil {
		logger.Fatal(err)
	}

	query := `SELECT title FROM film`

	var filmName []string
	if err := db.SelectContext(ctx, &filmName, query); err != nil {
		logger.Fatal(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(filmName))

	for index, name := range filmName {
		fmt.Printf("%d: %s\n", index, name)
		time.Sleep(time.Millisecond * 200)

		go func(name string) {
			defer wg.Done()

			if err := getFilm(config.Delivery.HTTPServer.ListenAddress, name); err != nil {
				logger.Error(err)
			}

			time.Sleep(time.Millisecond * 10)

			if err := getFilm(config.Delivery.HTTPServer.ListenAddress, name); err != nil {
				logger.Error(err)
			}

			time.Sleep(config.Extra.LocalCache.TimeLive)

			if err := getFilm(config.Delivery.HTTPServer.ListenAddress, name); err != nil {
				logger.Error(err)
			}

		}(name)
	}

	wg.Wait()

	getStatistic(config.Delivery.HTTPServer.ListenAddress, len(filmName))
}

func getFilm(listenAddress, name string) error {
	link := fmt.Sprintf("http://%s/v1/films-api/film/%s",
		listenAddress,
		url.QueryEscape(name),
	)

	resp, err := http.Get(link)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("for film %s status code is not 200, got %d", name, resp.StatusCode)
	}

	return nil
}

func getStatistic(listenAddress string, limit int) error {
	link := fmt.Sprintf("http://%s/v1/films-api/statistics?limit=%v", listenAddress, limit)

	resp, err := http.Get(link)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("for statistic status code is not 200, got %d", resp.StatusCode)
	}

	var response statistics.StatisticList
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	for _, statistic := range response.FilmStatistics {
		fmt.Printf("Request: %s, timeDB %d, timeRedis %d, timeCache%d\n",
			statistic.Request,
			statistic.TimeDB,
			statistic.TimeRedis,
			statistic.TimeMemory,
		)
	}

	return nil
}
