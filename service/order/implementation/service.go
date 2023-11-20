package implementation

import (
	"context"
	"database/sql"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

// service implements Service
type service struct {
	repository ordersvc.Repository
	logger     log.Logger
}

// NewSerive creates and returns a new Order service instance
func NewService(rep ordersvc.Repository, logger log.Logger) ordersvc.Service {
	return &service{
		repository: rep,
		logger:     logger,
	}
}

// Create makes an order
func (s *service) Create(ctx context.Context, order ordersvc.Order) (string, error) {
	logger := log.With(s.logger, "method", "Create")
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	order.Status = "Pending"
	order.CreatedOn = time.Now().Unix()

	if err := s.repository.CreateOrder(ctx, order); err != nil {
		level.Error(logger).Log("err", err)
		return "", ordersvc.ErrCmdRepository
	}
	return id, nil
}

// GetByID returns an order given by id
func (s *service) GetByID(ctx context.Context, id string) (ordersvc.Order, error) {
	logger := log.With(s.logger, "method", "GetByID")
	order, err := s.repository.GetOrderByID(ctx, id)
	if err != nil {
		level.Error(logger).log("err", err)
		if err == sql.ErrNoRows {
			return order, ordersvc.ErrQueryRepository
		}
		return order, nil
	}
}
