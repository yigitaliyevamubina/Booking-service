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

func TestPaymentSuite(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock: %s", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := p.NewBookingRepo(sqlxDB)

	t.Run("TestMakePayment", func(t *testing.T) {

		mock.ExpectQuery(`INSERT INTO patient_payment`).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("test-payment-id"))

		req := &pb.PatientPayment{
			Id:            "test-payment-id",
			AppointmentId: "appointment-id",
			PatientId:     "patient-id",
			Type:          "type",
			Amount:        100.0,
			Status:        "status",
			Ispaid:        true,
		}

		_, err = repo.MakePayment(req)

	})

	t.Run("TestGetPayment", func(t *testing.T) {

		rows := sqlmock.NewRows([]string{"id", "appointment_id", "patient_id", "type", "amount", "status", "paid", "created_at", "updated_at"}).
			AddRow("test-payment-id", "appointment-id", "patient-id", "type", 100.0, "status", true, time.Now(), nil)
		mock.ExpectQuery(`SELECT`).
			WithArgs("patient-id").
			WillReturnRows(rows)

			resp, err := repo.GetPayment(&pb.GetPaymentReq{Id: "patient-id"})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("TestDeletePayment", func(t *testing.T) {

		mock.ExpectExec(`UPDATE patient_payment`).
			WillReturnResult(sqlmock.NewResult(0, 1))


		success, err := repo.DeletePayment(&pb.GetPaymentReq{Id: "test-payment-id"})

		assert.NoError(t, err)
		assert.True(t, success)
	})
}
