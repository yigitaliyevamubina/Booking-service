package postgres_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	pb "Booking-service/genproto/booking-service"
	p "Booking-service/storage/postgres"
)

func TestDoctorAvailabilitySuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestCreateDoctorAvailability", func(t *testing.T) {

		mock.ExpectExec(`INSERT INTO doctor_availability`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		req := &pb.CreateDoctorAvailabilitys{
			Id:               "fd76880a-d2c2-4e43-b7eb-c6f555709323",
			DoctorId:         "f5f17b11-e9ff-4bd2-9ea8-9e311cda845a",
			DepartmentId:     "e1b673e5-973e-411b-b146-6c1c7d1aa9dd",
			AvailabilityDate: time.Now().Format("2006-01-02"),
			AvailabilityTime: time.Now().Format("15:04:05"),
			Status:           true,
		}

		_, err = repo.CreateDoctorAvailability(req)
	})

	t.Run("TestGetDoctorAvailability", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "doctor_id", "department_id", "availability_date", "availability_time", "status", "created_at", "updated_at"}).
			AddRow("test-id", "doctor-id", "department-id", time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), true, time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("f5f17b11-e9ff-4bd2-9ea8-9e311cda845a").
			WillReturnRows(rows)

		resp, err := repo.GetDoctorAvailability(&pb.GetDoctorAvailabilityById{Id: "f5f17b11-e9ff-4bd2-9ea8-9e311cda845a"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeleteDoctorAvailability", func(t *testing.T) {

		mock.ExpectExec(`UPDATE doctor_availability`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		success, err := repo.DeleteDoctorAvailability(&pb.GetDoctorAvailabilityById{Id: "f5f17b11-e9ff-4bd2-9ea8-9e311cda845a"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
