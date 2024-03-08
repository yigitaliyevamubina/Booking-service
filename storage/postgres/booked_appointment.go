package postgres

import (
	"context"
	"fmt"
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
			type, 
			duration, 
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
		&ba.Type, 
		&ba.Duration,
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

func (r *bookingRepo) GetBookedAppointmentsByPatientID(ctx context.Context, patientID *pb.PatientID) (*pb.GetBookedAppointmentsByPatientIDResponse, error) {
    appointments, err := r.db.QueryContext(ctx, `SELECT 
        id, 
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
    FROM booked_appointments
    WHERE patient_id = $1 AND deleted_at IS NULL`, patientID.PatientId)
    if err != nil {
		fmt.Println(err)
        return nil, err
    }
    defer appointments.Close()

    // Create a new response object
    bookedAppointments := &pb.GetBookedAppointmentsByPatientIDResponse{}

    for appointments.Next() {
        var appointment pb.BookedAppointment
        err := appointments.Scan(
            &appointment.Id,
            &appointment.DepartmentId,
            &appointment.DoctorId,
            &appointment.PatientId,
            &appointment.AppointmentDate,
            &appointment.AppointmentTime,
            &appointment.Type,
            &appointment.Duration,
            &appointment.ExpiresAt,
            &appointment.Token,
            &appointment.PatientStatus,
            &appointment.Status)
        if err != nil {
			fmt.Println(err)
            return nil, err
        }

        // Append each appointment to the response
        bookedAppointments.BookedAppointments = append(bookedAppointments.BookedAppointments, &appointment)
    }

    return bookedAppointments, nil
}


func (r *bookingRepo) GetBookedAppointmentsByDoctorID(ctx context.Context, doctorID *pb.GetBookedAppointmentRequest) (*pb.GetBookedAppointmentsByPatientIDResponse, error) {
    appointments, err := r.db.QueryContext(ctx, `SELECT 
        id, 
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
    FROM booked_appointments
    WHERE doctor_id = $1 AND deleted_at IS NULL`, doctorID.Id)
    if err != nil {
        return nil, err
    }
    defer appointments.Close()

    bookedAppointments := &pb.GetBookedAppointmentsByPatientIDResponse{}

    for appointments.Next() {
        var appointment pb.BookedAppointment
        err := appointments.Scan(
            &appointment.Id,
            &appointment.DepartmentId,
            &appointment.DoctorId,
            &appointment.PatientId,
            &appointment.AppointmentDate,
            &appointment.AppointmentTime,
            &appointment.Type,
            &appointment.Duration,
            &appointment.ExpiresAt,
            &appointment.Token,
            &appointment.PatientStatus,
            &appointment.Status)
        if err != nil {
            return nil, err
        }

        // Append each appointment to the response
        bookedAppointments.BookedAppointments = append(bookedAppointments.BookedAppointments, &appointment)
    }

    return bookedAppointments, nil
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

// func (r *bookingRepo) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (del *pb.IsDeleted, err error) {
// 	fmt.Println("MANA>>>>",req.Id)

// 	_, err = r.db.Exec(`
//         UPDATE booked_appointments
//         SET deleted_at = $1
//         WHERE patient_id = $2
//     `, time.Now(), req.Id)
// 	if err != nil {
// 		fmt.Println("MANA>>>>", err)
// 		return nil, err
// 	}
// 	fmt.Println("MANA>>>>",del)
// 	//del.IsDeleted = true
// 	return del, nil
// }

func (r *bookingRepo) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (del *pb.Status, err error) {
	_, err = r.db.Exec(`
		UPDATE booked_appointments
        SET deleted_at = $1
        WHERE patient_id = $2
	`, time.Now(), req.Id)
	if err != nil {
		return nil,err
	}

	// Soft delete successful, return nil
	return del,nil
}