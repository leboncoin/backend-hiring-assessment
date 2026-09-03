package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/mock"
	adsdao "github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/dao"
	"github.mpi-internal.com/leboncoin/backend-hiring-assessment/ads/internal/pkg/model"
)

// Mock is a testify-based implementation of dao.DAO.
type Mock struct {
	mock.Mock
}

// New returns a mock bound to the given test. Every expectation set on it is verified when the test finishes.
func New(t *testing.T) *Mock {
	t.Helper()

	m := &Mock{}
	m.Test(t)
	t.Cleanup(func() { m.AssertExpectations(t) })

	return m
}

// SearchAds returns the ads and the error set on the mock.
func (m *Mock) SearchAds(ctx context.Context, req adsdao.SearchAdsRequest) ([]model.Ad, error) {
	args := m.Called(ctx, req)

	ads, _ := args.Get(0).([]model.Ad)

	return ads, args.Error(1)
}

// GetAdByID returns the ad and the error set on the mock.
func (m *Mock) GetAdByID(ctx context.Context, id model.AdID) (model.Ad, error) {
	args := m.Called(ctx, id)

	ad, _ := args.Get(0).(model.Ad)

	return ad, args.Error(1)
}

// CreateAd returns the ad and the error set on the mock.
func (m *Mock) CreateAd(ctx context.Context, req adsdao.CreateAdRequest) (model.AdID, error) {
	args := m.Called(ctx, req)

	adID, _ := args.Get(0).(model.AdID)

	return adID, args.Error(1)
}
