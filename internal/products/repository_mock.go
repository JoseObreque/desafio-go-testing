package products

import "github.com/stretchr/testify/mock"

// RepositoryMock is a mock implementation of the Repository interface, using testify/mock.
type RepositoryMock struct {
	Mock mock.Mock
}

// GetAllBySeller is a mock implementation of the Repository interface's method.
func (m *RepositoryMock) GetAllBySeller(sellerID string) ([]Product, error) {
	args := m.Mock.Called(sellerID)
	return args.Get(0).([]Product), args.Error(1)
}

// NewRepositoryMock returns a new mock implementation of the Repository interface.
func NewRepositoryMock() *RepositoryMock {
	return &RepositoryMock{}
}
