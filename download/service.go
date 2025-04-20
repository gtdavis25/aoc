package download

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

type AOCClient interface {
	GetPuzzleInput(context.Context, int, int) (io.ReadCloser, error)
	GetDaysForYear(context.Context, int) ([]int, error)
	GetYears(context.Context) ([]int, error)
}

type Service struct {
	client AOCClient
}

func NewService(client AOCClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) DownloadAll(ctx context.Context, directory string) error {
	if err := createDirectoryIfNotExists(directory); err != nil {
		return err
	}

	years, err := s.client.GetYears(ctx)
	if err != nil {
		return fmt.Errorf("fetching years: %w", err)
	}

	group, groupCtx := errgroup.WithContext(ctx)
	for _, year := range years {
		year := year
		group.Go(func() error {
			directory := fmt.Sprintf("%s/%d", directory, year)
			if err := s.DownloadYear(groupCtx, year, directory); err != nil {
				return fmt.Errorf("year %d: %w", year, err)
			}

			return nil
		})
	}

	return group.Wait()
}

func (s *Service) DownloadYear(ctx context.Context, year int, directory string) error {
	if err := createDirectoryIfNotExists(directory); err != nil {
		return err
	}

	days, err := s.client.GetDaysForYear(ctx, year)
	if err != nil {
		return fmt.Errorf("fetching days for year %d: %w", year, err)
	}

	group, groupCtx := errgroup.WithContext(ctx)
	for _, day := range days {
		day := day
		group.Go(func() error {
			if err := s.DownloadDay(groupCtx, year, day, fmt.Sprintf("%s/%02d.txt", directory, day)); err != nil {
				return fmt.Errorf("downloading puzzle input for %d day %d: %w", year, day, err)
			}

			return nil
		})
	}

	return group.Wait()
}

func (s *Service) DownloadDay(ctx context.Context, year, day int, filePath string) error {
	i := strings.LastIndex(filePath, "/")
	if i != -1 {
		directory := filePath[:i]
		if err := createDirectoryIfNotExists(directory); err != nil {
			return fmt.Errorf("creating output directory %s: %w", directory, err)
		}
	}

	r, err := s.client.GetPuzzleInput(ctx, year, day)
	if err != nil {
		return fmt.Errorf("downloading puzzle input for %d day %d: %w", year, day, err)
	}

	defer r.Close()
	f, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("creating output file %s: %w", filePath, err)
	}

	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return fmt.Errorf("writing output file %s: %w", filePath, err)
	}

	return nil
}

func createDirectoryIfNotExists(directory string) error {
	info, err := os.Stat(directory)
	switch {
	case errors.Is(err, os.ErrNotExist):
		if err := os.MkdirAll(directory, 0777); err != nil {
			return fmt.Errorf("creating output directory %s: %w", directory, err)
		}

	case err != nil:
		return fmt.Errorf("stat %s: %w", directory, err)

	case !info.IsDir():
		return fmt.Errorf("not a directory: %s", directory)
	}

	return nil
}
