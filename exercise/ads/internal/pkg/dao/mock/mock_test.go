package mock_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	adsdao "github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/dao"
	"github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/dao/mock"
	"github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/model"
)

var ctx = context.Background()

func TestMock_SearchAds(t *testing.T) {
	t.Run("returns configured ads", func(t *testing.T) {
		m := mock.New(t)
		req := adsdao.SearchAdsRequest{Title: "bike"}
		want := []model.Ad{{ID: 1, Title: "bike", Price: 5000}}

		m.On("SearchAds", ctx, req).Return(want, nil)

		got, err := m.SearchAds(ctx, req)

		require.Error(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("returns nil slice on error", func(t *testing.T) {
		m := mock.New(t)
		req := adsdao.SearchAdsRequest{Title: "bike"}
		boom := errors.New("db down")

		m.On("SearchAds", ctx, req).Return(nil, boom)

		got, err := m.SearchAds(ctx, req)

		assert.ErrorIs(t, err, boom)
		assert.Nil(t, got)
	})
}

func TestMock_GetAdByID(t *testing.T) {
	t.Run("returns configured ad", func(t *testing.T) {
		m := mock.New(t)
		want := model.Ad{ID: 42, Title: "surfboard", Price: 15000}

		m.On("GetAdByID", ctx, model.AdID(42)).Return(want, nil)

		got, err := m.GetAdByID(ctx, 42)

		require.NoError(t, err)
		assert.Equal(t, want, got)
	})

	t.Run("returns ErrNotFound", func(t *testing.T) {
		m := mock.New(t)

		m.On("GetAdByID", ctx, model.AdID(99)).Return(model.Ad{}, adsdao.ErrNotFound)

		_, err := m.GetAdByID(ctx, 99)

		assert.ErrorIs(t, err, adsdao.ErrNotFound)
	})
}

func TestMock_CreateAd(t *testing.T) {
	t.Run("returns generated ID", func(t *testing.T) {
		m := mock.New(t)
		req := adsdao.CreateAdRequest{
			Title:      "bike",
			PriceCents: 5000,
			PhotoURL:   "https://example.com/bike.jpg",
			OwnerID:    "user-1",
		}

		m.On("CreateAd", ctx, req).Return(model.AdID(7), nil)

		got, err := m.CreateAd(ctx, req)

		require.NoError(t, err)
		assert.Equal(t, model.AdID(7), got)
	})

	t.Run("returns ErrUnknownOwner when owner does not exist", func(t *testing.T) {
		m := mock.New(t)
		req := adsdao.CreateAdRequest{
			Title:      "bike",
			PriceCents: 5000,
			PhotoURL:   "https://example.com/bike.jpg",
			OwnerID:    "ghost",
		}

		m.On("CreateAd", ctx, req).Return(model.AdID(0), adsdao.ErrUnknownOwner)

		_, err := m.CreateAd(ctx, req)

		assert.ErrorIs(t, err, adsdao.ErrUnknownOwner)
	})
}
