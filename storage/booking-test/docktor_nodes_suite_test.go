package postgres_test

import (
	"testing"
	"time"

	pb "Booking-service/genproto/booking-service"
	p "Booking-service/storage/postgres"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestDoctorNotesSuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestCreateDoctorNote", func(t *testing.T) {

		mock.ExpectQuery(`INSERT INTO doctor_notes`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("test-note-id"))

			req := &pb.CreateDoctorNoteReq{
			Id:            "test-note-id",
			AppointmentId: "appointment-id",
			DoctorId:      "doctor-id",
			PatientId:     "patient-id",
			NoteType:      "prescription",
			NoteText:      "Test prescription note",
		}

		_, err = repo.CreateDoctorNote(req)

	})

	t.Run("TestGetDoctorNote", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "appointment_id", "doctor_id", "patient_id", "note_type", "note_text", "created_at", "updated_at"}).
			AddRow("test-note-id", "appointment-id", "doctor-id", "patient-id", "prescription", "Test prescription note", time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("test-note-id").
			WillReturnRows(rows)

		resp, err := repo.GetDoctorNote(&pb.GetDoctorNoteReq{Id: "test-note-id"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeleteDoctorNote", func(t *testing.T) {
		mock.ExpectExec(`UPDATE doctor_notes`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		success, err := repo.DeleteDoctorNote(&pb.GetDoctorNoteReq{Id: "test-note-id"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
