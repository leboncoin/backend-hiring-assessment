package usecase

import (
	"context"
	"fmt"

	adsdao "github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/dao"
	"github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/model"
)

type (
	// ListAdsFunc returns all ads matching the filters.
	ListAdsFunc func(ctx context.Context, req ListAdsRequest) ([]model.Ad, error)

	// ListAdsRequest describes which ads to return.
	ListAdsRequest struct {
		OwnerID       model.UserID
		MinPriceCents int64
		MaxPriceCents int64
		Title         string
	}

	listAds struct {
		adDAO adsdao.DAO
	}
)

// NewListAdsFunc returns a new ListAdsFunc.
func NewListAdsFunc(adDAO adsdao.DAO) ListAdsFunc {
	return listAds{adDAO: adDAO}.listAds
}

func (uc listAds) listAds(ctx context.Context, req ListAdsRequest) ([]model.Ad, error) {
	ads, err := uc.adDAO.SearchAds(ctx, adsdao.SearchAdsRequest{
		OwnerID:       req.OwnerID,
		MinPriceCents: req.MinPriceCents,
		MaxPriceCents: req.MaxPriceCents,
		Title:         req.Title,
	})
	if err != nil {
		return nil, fmt.Errorf("unable to search ads: %w", err)
	}

	return ads, nil
}
