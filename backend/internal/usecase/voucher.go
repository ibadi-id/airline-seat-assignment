package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/ibadi-id/airline-seat-assignment/backend/internal/domain"
	"github.com/ibadi-id/airline-seat-assignment/backend/internal/repository"
)

type VoucherUsecase interface {
	Check(flightNumber, date string) (bool, error)
	Generate(v domain.Voucher) ([]string, error)
}

type voucherUsecase struct {
	Repo repository.VoucherRepository
}

func NewVoucherUsecase(r repository.VoucherRepository) VoucherUsecase {
	return &voucherUsecase{r}
}

func (u *voucherUsecase) Check(flightNumber, date string) (bool, error) {
	return u.Repo.Exists(flightNumber, date)
}

func (u *voucherUsecase) Generate(v domain.Voucher) ([]string, error) {
	if exists, _ := u.Repo.Exists(v.FlightNumber, v.FlightDate); exists {
		return nil, nil
	}

	seats := generateSeats(v.AircraftType)

	if len(seats) < 3 {
		return nil, errors.New("aircraft type not valid")
	}

	v.Seat1, v.Seat2, v.Seat3 = seats[0], seats[1], seats[2]

	err := u.Repo.Save(v)
	if err != nil {
		return nil, err
	}

	return seats, nil
}

func generateSeats(aircraftType string) []string {
	layout := map[string]struct {
		Rows  int
		Seats []string
	}{
		"ATR":            {18, []string{"A", "C", "D", "F"}},
		"Airbus 320":     {32, []string{"A", "B", "C", "D", "E", "F"}},
		"Boeing 737 Max": {32, []string{"A", "B", "C", "D", "E", "F"}},
	}

	conf, ok := layout[aircraftType]
	if !ok {
		return nil
	}

	allSeats := []string{}

	for i := 1; i <= conf.Rows; i++ {
		for _, s := range conf.Seats {
			allSeats = append(allSeats, formatSeat(i, s))
		}
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allSeats), func(i, j int) {
		allSeats[i], allSeats[j] = allSeats[j], allSeats[i]
	})

	return allSeats[:3]
}

func formatSeat(row int, seat string) string {
	return fmt.Sprintf("%d%s", row, seat)
}
