package products

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

// --------------------------------------------------------------------------------------------
// UNIT TESTS
// --------------------------------------------------------------------------------------------

// TestService_GetAllBySeller tests the service's GetAllBySeller method.
func TestService_GetAllBySeller(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		// Arrange
		repo := NewRepositoryMock()
		var expectedError error
		expectedProducts := []Product{
			{
				ID:          "mockID",
				SellerID:    "FEX112AC",
				Description: "mock product",
				Price:       123.55,
			},
		}
		repo.Mock.On("GetAllBySeller", "FEX112AC").Return(expectedProducts, expectedError)
		service := NewService(repo)

		// Act
		products, err := service.GetAllBySeller("FEX112AC")

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, products)
	})

	t.Run("Repository error", func(t *testing.T) {
		// Arrange
		repo := NewRepositoryMock()
		expectedError := errors.New("mock repository error")
		var expectedProducts []Product

		repo.Mock.On("GetAllBySeller", "Some_ID").Return(expectedProducts, expectedError)

		service := NewService(repo)

		// Act
		products, err := service.GetAllBySeller("Some_ID")

		// Assert
		assert.Error(t, err)
		assert.Equal(t, expectedProducts, products)
	})
}

// --------------------------------------------------------------------------------------------
// INTEGRATION TESTS
// --------------------------------------------------------------------------------------------

// TestService_GetAllBySeller_Integration tests the service's GetAllBySeller method.
func TestService_GetAllBySeller_Integration(t *testing.T) {
	// Arrange
	repo := NewRepository()
	testService := NewService(repo)
	sellerID := "FEX112AC"

	// Act
	products, err := testService.GetAllBySeller(sellerID)

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}
