package postgres

import (
	pb "Booking-service/genproto/booking-service"
	"database/sql"
	"errors"
	"time"
)

func (r *bookingRepo) MakePayment(request *pb.PatientPayment) (*pb.PatientPayment, error) {
	query := `
		INSERT INTO patient_payment (
			id, 
			appointment_id, 
			patient_id, 
			type, 
			amount, 
			status, 
			paid,
		    created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING 
			id, 
			appointment_id, 
			patient_id, 
			type, 
			amount,
			status, 
			paid,
			created_at
	`
	row := r.db.QueryRow(
		query,
		request.Id,
		request.AppointmentId,
		request.PatientId,
		request.Type,
		request.Amount,
		request.Status,
		request.Ispaid,
		time.Now(),
	)

	var payment pb.PatientPayment
	err := row.Scan(
		&payment.Id,
		&payment.AppointmentId,
		&payment.PatientId,
		&payment.Type,
		&payment.Amount,
		&payment.Status,
		&payment.Ispaid,
		&payment.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *bookingRepo) GetPayment(request *pb.GetPaymentReq) (*pb.PatientPayment, error) {
	var updatedAtPtr *string
	query := `
		SELECT 
			id, 
			appointment_id, 
			patient_id, 
			type, 
			amount, 
			status,
			paid,
			created_at, 
			updated_at
		FROM patient_payment 
		WHERE patient_id = $1 AND deleted_at IS NULL
	`
	row := r.db.QueryRow(query, request.Id)

	var payment pb.PatientPayment
	err := row.Scan(
		&payment.Id,
		&payment.AppointmentId,
		&payment.PatientId,
		&payment.Type,
		&payment.Amount,
		&payment.Status,
		&payment.Ispaid,
		&payment.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		return nil, err
	}
	if updatedAtPtr != nil {
		payment.UpdateAt = *updatedAtPtr
	}
	return &payment, nil
}

func (r *bookingRepo) GetPaymentsByPatienId(request *pb.GetPaymentReq) (*pb.GetPaymentsResp, error) {

	query := `
		SELECT 
			id, 
			appointment_id, 
			patient_id, 
			type, 
			amount, 
			status, 
			paid,
			created_at, 
			updated_at
		FROM patient_payment
		WHERE patient_id = $1 AND deleted_at IS NULL
	`
	rows, err := r.db.Query(query, request.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var payments []*pb.PatientPayment // Define payments as a slice

	for rows.Next() {
		var payment pb.PatientPayment
		var updatedAtPtr *string
		err := rows.Scan(
			&payment.Id,
			&payment.AppointmentId,
			&payment.PatientId,
			&payment.Type,
			&payment.Amount,
			&payment.Status,
			&payment.Ispaid,
			&payment.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}
		if updatedAtPtr != nil {
			payment.UpdateAt = *updatedAtPtr
		}
		payments = append(payments, &payment)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	response := &pb.GetPaymentsResp{
		PatientPayment: payments,
	}

	return response, nil
}

func (r *bookingRepo) UpdatePayment(request *pb.UpdatePaymentRequest) (*pb.PatientPayment, error) {

	query := `
		UPDATE patient_payment
		SET 
			amount = $1, 
			status = $2, 
			paid = $3,
			updated_at = $4
		WHERE patient_id = $5 AND deleted_at IS NULL
		RETURNING 
			id, 
			appointment_id, 
			patient_id, 
			type, 
			amount, 
			status, 
			paid,
			created_at, 
			updated_at
	`
	row := r.db.QueryRow(
		query,
		request.Payment.Amount,
		request.Payment.Status,
		request.Payment.Ispaid,
		time.Now(),
		request.Id,
	)

	var payment pb.PatientPayment
	err := row.Scan(
		&payment.Id,
		&payment.AppointmentId,
		&payment.PatientId,
		&payment.Type,
		&payment.Amount,
		&payment.Status,
		&payment.Ispaid,
		&payment.CreateAt,
		&payment.UpdateAt,
	)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (r *bookingRepo) DeletePayment(request *pb.GetPaymentReq) (bool, error) {

	result, err := r.db.Exec(`
		UPDATE patient_payment
		SET deleted_at = $1
		WHERE patient_id = $2 AND deleted_at IS NULL
	`, time.Now(), request.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, err
		}
		return false, err
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return false, err
	}

	return true, nil
}
