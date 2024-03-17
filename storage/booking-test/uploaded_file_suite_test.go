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

func TestFileSuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestUploadFile", func(t *testing.T) {

		mock.ExpectQuery(`INSERT INTO uploaded_files`).
			WillReturnRows(sqlmock.NewRows([]string{"file_id"}).AddRow("test-file-id"))

		req := &pb.CreateUploadedFile{
			FileId:    "test-file-id",
			PatientId: "patient-id",
			RequestId: "request-id",
			File:      []byte{1, 2, 3}, 
		}

		_, err = repo.UploadFile(req)


	})

	t.Run("TestGetFileByPatientID", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"file_id", "patient_id", "request_id", "file", "created_at", "updated_at"}).
			AddRow("test-file-id", "patient-id", "request-id", []byte{1, 2, 3}, time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("patient-id").
			WillReturnRows(rows)

		resp, err := repo.GetFileByPatientID(&pb.GetFileRequest{FileId: "patient-id"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeleteFile", func(t *testing.T) {

		mock.ExpectExec(`UPDATE uploaded_files`).
			WillReturnResult(sqlmock.NewResult(0, 1))

		success, err := repo.DeleteFile(&pb.GetFileRequest{FileId: "test-file-id"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
