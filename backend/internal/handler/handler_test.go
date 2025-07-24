package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ibadi-id/airline-seat-assignment/backend/internal/domain"
	"github.com/ibadi-id/airline-seat-assignment/backend/internal/usecase"
	"github.com/ibadi-id/airline-seat-assignment/backend/pkg/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCheckHandler(t *testing.T) {
	e := echo.New()
	e.Validator = validator.NewValidator()

	mockUsecase := usecase.NewMockVoucherUsecase(t)
	h := NewHandler(mockUsecase)

	reqBody := CheckRequest{
		FlightNumber: "GA123",
		Date:         "2025-07-12",
	}

	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/check", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	mockUsecase.On("Check", "GA123", "2025-07-12").Return(true, nil)

	err := h.Check(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	expected := `{"success":true,"exists":true}`
	assert.JSONEq(t, expected, rec.Body.String())
	mockUsecase.AssertExpectations(t)
}

func TestGenerateHandler_Success(t *testing.T) {
	e := echo.New()
	e.Validator = validator.NewValidator()

	mockUsecase := usecase.NewMockVoucherUsecase(t)
	h := NewHandler(mockUsecase)

	reqBody := GenerateRequest{
		Name:         "John Doe",
		ID:           "C001",
		FlightNumber: "GA123",
		Date:         "2025-12-01",
		Aircraft:     "ATR",
	}
	jsonBody, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/generate", bytes.NewReader(jsonBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	expectedVoucher := domain.Voucher{
		CrewName:     "John Doe",
		CrewID:       "C001",
		FlightNumber: "GA123",
		FlightDate:   "2025-12-01",
		AircraftType: "ATR",
	}

	seats := []string{"1A", "2B", "3C"}

	mockUsecase.
		On("Generate", expectedVoucher).
		Return(seats, nil)

	err := h.Generate(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp GenerateResponse
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.True(t, resp.Success)
	assert.Equal(t, seats, resp.Seats)

	mockUsecase.AssertExpectations(t)
}
