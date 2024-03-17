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

func TestBookedAppointmentSuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestCreateBookedAppointment", func(t *testing.T) {

		mock.ExpectQuery(`INSERT INTO booked_appointments`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("test-id"))

		req := &pb.CreateBookedAppointments{
			Id:              "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444",
			DepartmentId:    "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444",
			DoctorId:        "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444",
			PatientId:       "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444",
			AppointmentDate: time.Now().Format("2006-01-02"),
			AppointmentTime: time.Now().Format("15:04:05"),
			Type:            "type",
			Duration:        "duration",
			ExpiresAt:       time.Now().Add(time.Hour).Format(time.RFC3339),
			Token:           "token",
			PatientStatus:   true,
			Status:          "completed",
		}

		_, err = repo.CreateBookedAppointment(req)

	})

	t.Run("TestGetBookedAppointment", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "department_id", "doctor_id", "patient_id", "appointment_date", "appointment_time", "type", "duration", "expires_at", "token", "patient_status", "status", "created_at", "updated_at"}).
			AddRow("test-id", "department-id", "doctor-id", "patient-id", time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), "type", "duration", time.Now().Add(time.Hour).Format(time.RFC3339), "token", true, "completed", time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("5e518f7a-e14d-4b7b-8b1d-9cfb5a128444").
			WillReturnRows(rows)

		resp, err := repo.GetBookedAppointment(&pb.GetRequest{Id: "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeleteBookedAppointment", func(t *testing.T) {

		mock.ExpectExec(`UPDATE booked_appointments`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		success, err := repo.DeleteBookedAppointment(&pb.GetRequest{Id: "5e518f7a-e14d-4b7b-8b1d-9cfb5a128444"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
