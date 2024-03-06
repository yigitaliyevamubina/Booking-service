package postgres

import (
	"context"
	"log"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        INSERT INTO booked_appointments (
			department_id, 
			doctor_id, 
			patient_id, 
			appointment_date, 
			appointment_time, 
			type, 
			duration, 
			expires_at, 
			token,
			patient_status,
			status
		)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
		RETURNING id, 
			department_id, 
			doctor_id, 
			patient_id, 
			appointment_date, 
			appointment_time, 
			type, 
			duration, 
			expires_at, 
			token, 
			patient_status, 
			status
	 `, req.BookedAppointment.DepartmentId,
		req.BookedAppointment.DoctorId,
		req.BookedAppointment.PatientId,
		req.BookedAppointment.AppointmentDate,
		req.BookedAppointment.AppointmentTime,
		req.BookedAppointment.Type,
		req.BookedAppointment.Duration,
		req.BookedAppointment.ExpiresAt,
		req.BookedAppointment.Token,
		req.BookedAppointment.PatientStatus,
		req.BookedAppointment.Status).Scan(
		&ba.Id, &ba.DepartmentId,
		&ba.DoctorId, &ba.PatientId,
		&ba.AppointmentDate,
		&ba.AppointmentTime,
		&ba.Type,
		&ba.Duration,
		&ba.ExpiresAt,
		&ba.Token,
		&ba.PatientStatus,
		&ba.Status)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	
	return &ba, nil
}

func (r *bookingRepo) GetBookedAppointment(ctx context.Context, req *pb.GetBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        SELECT 
			id, 
			department_id, 
			doctor_id, 
			patient_id, 
			appointment_date, 
			appointment_time, 
			type, duration, 
			expires_at, 
			token, 
			patient_status, 
			status
        FROM booked_appointments
        	WHERE patient_id = $1 AND deleted_at IS NULL
    `, req.Id).Scan(
		&ba.Id,
		&ba.DepartmentId,
		&ba.DoctorId,
		&ba.PatientId,
		&ba.AppointmentDate,
		&ba.AppointmentTime,
		&ba.Type, &ba.Duration,
		&ba.ExpiresAt,
		&ba.Token,
		&ba.PatientStatus,
		&ba.Status,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        UPDATE booked_appointments
        	SET 
				id = $2, 
				department_id = $3,
				doctor_id = $4, 
				patient_id = $5, 
				appointment_date = $6, 
				appointment_time = $7, 
				type = $8, 
				duration = $9, 
				expires_at = $10, 
				token = $11, 
				patient_status = $12, 
				status = $13
        	WHERE patient_id = $1 AND deleted_at IS NULL
        	RETURNING 
				id, 
				department_id, 
				doctor_id, patient_id, 
				appointment_date, 
				appointment_time, 
				type, 
				duration, 
				expires_at, 
				token, 
				patient_status, 
				status
     `, req.Id,
		req.BookedAppointment.Id,
		req.BookedAppointment.DepartmentId,
		req.BookedAppointment.DoctorId,
		req.BookedAppointment.PatientId,
		req.BookedAppointment.AppointmentDate,
		req.BookedAppointment.AppointmentTime,
		req.BookedAppointment.Type,
		req.BookedAppointment.Duration,
		req.BookedAppointment.ExpiresAt,
		req.BookedAppointment.Token,
		req.BookedAppointment.PatientStatus,
		req.BookedAppointment.Status).
		Scan(
			&ba.Id,
			&ba.DepartmentId,
			&ba.DoctorId,
			&ba.PatientId,
			&ba.AppointmentDate,
			&ba.AppointmentTime,
			&ba.Type, &ba.Duration,
			&ba.ExpiresAt,
			&ba.Token,
			&ba.PatientStatus,
			&ba.Status,
		)
	if err != nil {
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (del *pb.IsDeleted, err error) {

	_, err = r.db.ExecContext(ctx, `
        UPDATE booked_appointments
        SET deleted_at = $1
        WHERE patient_id = $2
    `, time.Now(), req.Id)
	if err != nil {
		del.IsDeleted = false
		return del, err
	}
	del.IsDeleted = true

	return del, nil
}