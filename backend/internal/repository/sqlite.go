package repository

import (
	"database/sql"
	"time"

	"github.com/ibadi-id/backend-bc/internal/domain"
)

type VoucherRepository interface {
	Exists(flightNumber, date string) (bool, error)
	Save(v domain.Voucher) error
}

type voucherRepo struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepo{db}
}

func (r *voucherRepo) Exists(flightNumber, date string) (bool, error) {
	var count int
	err := r.db.QueryRow("SELECT COUNT(*) FROM vouchers WHERE flight_number = ? AND flight_date = ?", flightNumber, date).Scan(&count)
	return count > 0, err
}

func (r *voucherRepo) Save(v domain.Voucher) error {
	query := `
		INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		v.CrewName, v.CrewID, v.FlightNumber, v.FlightDate, v.AircraftType,
		v.Seat1, v.Seat2, v.Seat3, time.Now().Format(time.RFC3339),
	)
	return err
}
