package services

import (
	"context"
	"fmt"

	"github.com/mgrinsted/pylos/internal/repository"

	models "github.com/mgrinsted/pylos/internal/models/estate"
)

type EstateCustomerService struct {
	repo repository.EstateCustomerRepository
}

func NewEstateCustomerService(repo repository.EstateCustomerRepository) *EstateCustomerService {
	return &EstateCustomerService{repo: repo}
}

// GetPaginated retrieves paginated customer summaries from the database.
func (s *EstateCustomerService) GetPaginated(limit, offset int) ([]models.CustomerSummaryDTO, int, error) {
	fmt.Printf("LIMIT: %d, OFFSET: %d\n", limit, offset)
	ctx := context.Background()

	customers, err := s.repo.GetAllSummary(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountAll(ctx)
	if err != nil {
		return nil, 0, err
	}

	return customers, total, nil
}

// GetByID retrieves a single customer by ID
func (s *EstateCustomerService) GetByID(id int) (*models.CustomerDetailDTO, error) {
	ctx := context.Background()

	customer, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
