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

func TestArchiveSuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestCreateArchive", func(t *testing.T) {

		mock.ExpectQuery(`INSERT INTO archive`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("test-id"))

		req := &pb.CreateArchiveReq{
			Id:               "test-id",
			DepartmentId:     "department-id",
			DoctorId:         "doctor-id",
			PatientId:        "patient-id",
			PatientToken:     "patient-token",
			PatientProblem:   "patient-problem",
			ConsultationType: "consultation-type",
			BookedDate:       time.Now().Format("2006-01-02"),
			BookedTime:       time.Now().Format("15:04:05"),
			AppointmentId:    "appointment-id",
			Status:           "status",
			VisitsCount:      1,
		}

		_, err = repo.CreateArchive(req)

	})

	t.Run("TestGetArchive", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "department_id", "doctor_id", "patient_id", "patient_token", "patient_problem", "consultation_type", "booked_date", "booked_time", "appointment_id", "status", "visits_count", "created_at", "updated_at"}).
			AddRow("test-id", "department-id", "doctor-id", "patient-id", "patient-token", "patient-problem", "consultation-type", time.Now().Format("2006-01-02"), time.Now().Format("15:04:05"), "appointment-id", "status", 1, time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("test-id").
			WillReturnRows(rows)

		resp, err := repo.GetArchive(&pb.GetArchiveReq{Id: "test-id"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeleteArchive", func(t *testing.T) {
		mock.ExpectExec(`UPDATE archive`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		success, err := repo.DeleteArchive(&pb.GetArchiveReq{Id: "test-id"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
