package products

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ------------------------------------------------------------
// UNIT TESTS
// ------------------------------------------------------------

// TestRepository_GetAllBySeller is a test for GetAllBySeller method.
func TestRepository_GetAllBySeller(t *testing.T) {
	// Arrange
	repo := NewRepository()
	sellerID := "FEX112AC"
	expectedProducts := []Product{
		{
			ID:          "mock",
			SellerID:    "FEX112AC",
			Description: "generic product",
			Price:       123.55,
		},
	}

	// Act
	prodList, err := repo.GetAllBySeller(sellerID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedProducts, prodList)
}
