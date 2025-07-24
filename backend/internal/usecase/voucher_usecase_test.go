package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ibadi-id/airline-seat-assignment/backend/internal/domain"
	"github.com/ibadi-id/airline-seat-assignment/backend/internal/repository"
)

func TestCheck(t *testing.T) {
	mockRepo := repository.NewMockVoucherRepository(t)
	u := NewVoucherUsecase(mockRepo)

	mockRepo.On("Exists", "GA123", "2025-12-01").Return(true, nil)

	exists, err := u.Check("GA123", "2025-12-01")

	assert.NoError(t, err)
	assert.True(t, exists)
	mockRepo.AssertExpectations(t)
}

func TestGenerate_SaveError(t *testing.T) {
	mockRepo := repository.NewMockVoucherRepository(t)
	u := NewVoucherUsecase(mockRepo)

	v := domain.Voucher{
		FlightNumber: "GA123",
		FlightDate:   "2025-12-01",
		AircraftType: "Airbus 320",
	}

	// Exists return false ➜ lanjut ke Save
	mockRepo.On("Exists", v.FlightNumber, v.FlightDate).Return(false, nil)

	// Save return error ➜ target line ter-cover
	mockRepo.On("Save", mock.Anything).Return(errors.New("db error"))

	// Call function
	seats, err := u.Generate(v)

	// ✅ Now assert error returned from Save
	assert.Nil(t, seats)
	assert.Error(t, err)
	assert.EqualError(t, err, "db error")

	mockRepo.AssertExpectations(t)
}

func TestGenerate_InvalidAircraftType(t *testing.T) {
	mockRepo := repository.NewMockVoucherRepository(t)
	u := NewVoucherUsecase(mockRepo)

	v := domain.Voucher{
		FlightNumber: "GA123",
		FlightDate:   "2025-12-01",
		AircraftType: "UNKNOWN",
	}

	mockRepo.On("Exists", v.FlightNumber, v.FlightDate).Return(false, nil)

	seats, err := u.Generate(v)

	assert.Nil(t, seats)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "aircraft type not valid")
	mockRepo.AssertExpectations(t)
}

func TestGenerate_AlreadyExists(t *testing.T) {
	mockRepo := repository.NewMockVoucherRepository(t)
	u := NewVoucherUsecase(mockRepo)

	v := domain.Voucher{
		FlightNumber: "GA123",
		FlightDate:   "2025-12-01",
		AircraftType: "ATR",
	}

	mockRepo.On("Exists", v.FlightNumber, v.FlightDate).Return(true, nil)

	seats, err := u.Generate(v)

	assert.NoError(t, err)
	assert.Nil(t, seats)
	mockRepo.AssertExpectations(t)
}

func TestGenerate_NewVoucher(t *testing.T) {
	mockRepo := repository.NewMockVoucherRepository(t)
	u := NewVoucherUsecase(mockRepo)

	v := domain.Voucher{
		FlightNumber: "GA123",
		FlightDate:   "2025-12-01",
		AircraftType: "ATR",
	}

	mockRepo.On("Exists", v.FlightNumber, v.FlightDate).Return(false, nil)
	mockRepo.On("Save", mock.MatchedBy(func(v domain.Voucher) bool {
		return v.Seat1 != "" && v.Seat2 != "" && v.Seat3 != "" &&
			v.Seat1 != v.Seat2 && v.Seat2 != v.Seat3 && v.Seat1 != v.Seat3
	})).Return(nil)

	seats, err := u.Generate(v)

	assert.NoError(t, err)
	assert.Len(t, seats, 3)
	assert.NotEqual(t, seats[0], seats[1])
	assert.NotEqual(t, seats[1], seats[2])
	mockRepo.AssertExpectations(t)
}
