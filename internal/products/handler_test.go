package products

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ------------------------------------------------------------
// UNIT TESTS
// ------------------------------------------------------------

func TestNewHandler(t *testing.T) {
	// Arrange
	repo := NewRepository()
	svc := NewService(repo)
	expectedHandler := &Handler{
		svc: svc,
	}

	// Act
	handler := NewHandler(svc)

	// Assert
	assert.Equal(t, expectedHandler, handler)
}

// ------------------------------------------------------------
// INTEGRATION TESTS
// ------------------------------------------------------------

func TestHandler_GetProducts_OK(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Arrange
		repoMock := NewRepositoryMock()

		var expectedRepoError error
		expectedProductList := []Product{
			{
				ID:          "mock",
				SellerID:    "FEX112AC",
				Description: "mock product",
				Price:       123.55,
			},
		}

		repoMock.Mock.On("GetAllBySeller", "FEX112AC").Return(expectedProductList, expectedRepoError)
		testService := NewService(repoMock)
		testHandler := NewHandler(testService)

		router := gin.New()
		router.GET("/products", testHandler.GetProducts)

		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedResponse := `[{"id":"mock","seller_id":"FEX112AC","description":"mock product","price":123.55}]`

		// Act
		url := "/products" + "?seller_id=FEX112AC"
		request := httptest.NewRequest("GET", url, nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})

	t.Run("missing query param", func(t *testing.T) {
		// Arrange
		repoMock := NewRepositoryMock()
		testService := NewService(repoMock)
		testHandler := NewHandler(testService)

		router := gin.New()
		router.GET("/products", testHandler.GetProducts)

		expectedResponse := `{"error":"seller_id query param is required"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		// Act
		url := "/products"
		request := httptest.NewRequest("GET", url, nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusBadRequest, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})

	t.Run("repo error", func(t *testing.T) {
		// Arrange
		repoMock := NewRepositoryMock()

		var expectedProductList []Product
		expectedRepoError := errors.New("repository error")

		repoMock.Mock.On("GetAllBySeller", "FEX112AC").Return(expectedProductList, expectedRepoError)
		testService := NewService(repoMock)
		testHandler := NewHandler(testService)

		router := gin.New()
		router.GET("/products", testHandler.GetProducts)

		expectedResponse := `{"error":"repository error"}`
		expectedHeader := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}

		// Act
		url := "/products" + "?seller_id=FEX112AC"
		request := httptest.NewRequest("GET", url, nil)
		request.Header.Set("Content-Type", "application/json")
		response := httptest.NewRecorder()
		router.ServeHTTP(response, request)

		// Assert
		assert.Equal(t, http.StatusInternalServerError, response.Code)
		assert.Equal(t, expectedHeader, response.Header())
		assert.Equal(t, expectedResponse, response.Body.String())
	})
}
