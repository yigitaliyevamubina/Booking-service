package postgres

import (
	"context"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        INSERT INTO booked_appointments (department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status,status)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id, department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status, status
    `, req.BookedAppointment.DepartmentId, req.BookedAppointment.DoctorId, req.BookedAppointment.PatientId, req.BookedAppointment.AppointmentDate, req.BookedAppointment.AppointmentTime, req.BookedAppointment.Type, req.BookedAppointment.Duration, req.BookedAppointment.ExpiresAt, req.BookedAppointment.Token, req.BookedAppointment.PatientStatus, req.BookedAppointment.Status).Scan(
		&ba.Id, &ba.DepartmentId, &ba.DoctorId, &ba.PatientId, &ba.AppointmentDate, &ba.AppointmentTime, &ba.Type, &ba.Duration, &ba.ExpiresAt, &ba.Token, &ba.PatientStatus, &ba.Status)
	if err != nil {
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) GetBookedAppointment(ctx context.Context, req *pb.GetBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        SELECT id, department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status, status
        FROM booked_appointments
        WHERE patient_id = $1
    `, req.Id).Scan(
		&ba.Id, &ba.DepartmentId, &ba.DoctorId, &ba.PatientId, &ba.AppointmentDate, &ba.AppointmentTime, &ba.Type, &ba.Duration, &ba.ExpiresAt, &ba.Token, &ba.PatientStatus, &ba.Status)
	if err != nil {
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        UPDATE booked_appointments
        SET id = $2, department_id = $3, doctor_id = $4, patient_id = $5, appointment_date = $6, appointment_time = $7, type = $8, duration = $9, expires_at = $10, token = $11, patient_status = $12, status = $13
        WHERE patient_id = $1
        RETURNING id, department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status, status
    `, req.Id, req.BookedAppointment.Id, req.BookedAppointment.DepartmentId, req.BookedAppointment.DoctorId, req.BookedAppointment.PatientId, req.BookedAppointment.AppointmentDate, req.BookedAppointment.AppointmentTime, req.BookedAppointment.Type, req.BookedAppointment.Duration, req.BookedAppointment.ExpiresAt, req.BookedAppointment.Token, req.BookedAppointment.PatientStatus, req.BookedAppointment.Status).Scan(
		&ba.Id, &ba.DepartmentId, &ba.DoctorId, &ba.PatientId, &ba.AppointmentDate, &ba.AppointmentTime, &ba.Type, &ba.Duration, &ba.ExpiresAt, &ba.Token, &ba.PatientStatus, &ba.Status)
	if err != nil {
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	// Fetch the data before deleting
	var deletedData pb.BookedAppointment
	err := r.db.QueryRowContext(ctx, `
        DELETE FROM booked_appointments
        WHERE patient_id = $1
        RETURNING id, department_id, doctor_id, patient_id, appointment_date, appointment_time, type, duration, expires_at, token, patient_status, status
    `, req.Id).Scan(
		&deletedData.Id, &deletedData.DepartmentId, &deletedData.DoctorId, &deletedData.PatientId, &deletedData.AppointmentDate, &deletedData.AppointmentTime, &deletedData.Type, &deletedData.Duration, &deletedData.ExpiresAt, &deletedData.Token, &deletedData.PatientStatus, &deletedData.Status)
	if err != nil {
		return nil, err
	}

	// Return the deleted data
	return &deletedData, nil
}
