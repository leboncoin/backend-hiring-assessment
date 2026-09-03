package usecase

import (
	"context"
	"fmt"

	adsdao "github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/dao"
	"github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/model"
)

type (
	// GetAdByIDFunc returns a single ad. It fails with ErrorCodeNotFound when no ad matches the identifier.
	GetAdByIDFunc func(ctx context.Context, id model.AdID) (model.Ad, error)

	getAdByID struct {
		adDAO adsdao.DAO
	}
)

// NewGetAdByIDFunc returns a GetAdByIDFunc backed by the given DAO.
func NewGetAdByIDFunc(adDAO adsdao.DAO) GetAdByIDFunc {
	return getAdByID{adDAO: adDAO}.getAdByID
}

func (uc getAdByID) getAdByID(ctx context.Context, id model.AdID) (model.Ad, error) {
	found, err := uc.adDAO.GetAdByID(ctx, 0)
	if err != nil {
		return model.Ad{}, fmt.Errorf("unable to get ad %s: %w", id, err)
	}

	return found, nil
}
