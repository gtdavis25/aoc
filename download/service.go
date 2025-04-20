package download

import (
	"context"
	"fmt"
	"io"
	"os"
)

type AOCClient interface {
	GetPuzzleInput(context.Context, int, int) (io.ReadCloser, error)
}

type Service struct {
	client AOCClient
}

func NewService(client AOCClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) DownloadPuzzleInput(ctx context.Context, year, day int, filePath string) error {
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
