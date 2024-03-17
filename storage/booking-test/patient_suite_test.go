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

func TestCreatePatient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := p.NewBookingRepo(sqlxDB)

	currentTime := time.Now().Format("15:04:05")

	req := &pb.Patient{
		Id:          "test-id",
		FirstName:   "John",
		LastName:    "Doe",
		BirthDate:   "1990-01-01",
		Gender:      "male",
		City:        "New York",
		PhoneNumber: "123456789",
		CreateAt:    currentTime,
	}

	mock.ExpectExec(`INSERT INTO patients`).WithArgs(
		req.Id, req.FirstName, req.LastName, req.BirthDate, req.Gender, req.City, req.PhoneNumber, req.CreateAt).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectQuery(`SELECT`).WithArgs(req.Id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date", "gender", "city", "phone_number", "created_at", "updated_at"}).
			AddRow(req.Id, req.FirstName, req.LastName, req.BirthDate, req.Gender, req.City, req.PhoneNumber, currentTime, nil))

	_, err = repo.CreatePatient(req)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPatient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := p.NewBookingRepo(sqlxDB)

	req := &pb.GetPatientRequest{
		Id: "123",
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date", "gender", "city", "phone_number", "created_at", "updated_at"}).
		AddRow("123", "John", "Doe", "1990-01-01", "Male", "New York", "123456789", time.Now(), nil)

	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)

	patient, err := repo.GetPatient(req)
	assert.NoError(t, err)
	assert.NotNil(t, patient)
	assert.Equal(t, req.Id, patient.Id)
}

func TestGetPatients(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := p.NewBookingRepo(sqlxDB)

	req := &pb.GetPatientsRequest{
		Limit: "10",
		Page:  "1",
	}

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date", "gender", "city", "phone_number", "created_at", "updated_at"}).
		AddRow("test-id", "John", "Doe", "1990-01-01", "male", "New York", "123456789", time.Now(), nil)

	mock.ExpectQuery(`SELECT`).
		WillReturnRows(rows)

	resp, err := repo.GetPatients(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 1, len(resp.Patient))
}

func TestUpdatePatient(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := p.NewBookingRepo(sqlxDB)

	req := &pb.UpdatePatientRequest{
		Id: "6a55e7da-7951-4cae-8c7d-534fbf9d280d",
		Patient: &pb.PatientUpdate{
			FirstName:   "John",
			LastName:    "Doe",
			BirthDate:   "1990-01-01",
			Gender:      "male",
			City:        "New York",
			PhoneNumber: "123456789",
		},
	}

	mock.ExpectExec(`UPDATE patients`).
		WillReturnResult(sqlmock.NewResult(0, 1))

	rows := sqlmock.NewRows([]string{"id", "first_name", "last_name", "birth_date", "gender", "city", "phone_number", "created_at", "updated_at"}).
		AddRow(req.Id, req.Patient.FirstName, req.Patient.LastName, req.Patient.BirthDate, req.Patient.Gender, req.Patient.City, req.Patient.PhoneNumber, time.Now(), nil)

	mock.ExpectQuery(`SELECT`).
		WillReturnRows(rows)

	resp, err := repo.UpdatePatient(req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, req.Id, resp.Id)
}


func TestDeletePatient_NoRowsAffected(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	repo := p.NewBookingRepo(sqlxDB)

	req := &pb.GetPatientRequest{
		Id: "6a55e7da-7951-4cae-8c7d-534fbf9d280d",
	}

	mock.ExpectExec(`UPDATE patients`).
		WillReturnResult(sqlmock.NewResult(0, 0))

	success, err := repo.DeletePatient(req)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NoError(t, err)
	assert.False(t, success)
}
