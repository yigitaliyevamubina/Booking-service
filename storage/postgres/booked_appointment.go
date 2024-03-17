package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateBookedAppointment(req *pb.CreateBookedAppointments) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRow(`
        INSERT INTO booked_appointments (
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
			status,
			created_at
		)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
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
			status,
			created_at
	 `,	req.Id,
	 	req.DepartmentId,
		req.DoctorId,
		req.PatientId,
		req.AppointmentDate,
		req.AppointmentTime,
		req.Type,
		req.Duration,
		req.ExpiresAt,
		req.Token,
		req.PatientStatus,
		req.Status,
		time.Now()).Scan(
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
		&ba.CreateAt,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ba, nil
}

func (r *bookingRepo) GetBookedAppointment(req *pb.GetRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	var updatedAtPtr *string
	err := r.db.QueryRow(`
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
			status,
			created_at,
			updated_at
        FROM booked_appointments
        	WHERE patient_id = $1 AND deleted_at IS NULL
    `,  req.Id).Scan(
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
		&ba.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if updatedAtPtr != nil {
		ba.UpdateAt = *updatedAtPtr
	}
	return &ba, nil
}


func (r *bookingRepo) GetBookedAppointmentsByDoctorID(doctorID *pb.GetRequest) (*pb.GetBookedAppointments, error) {
	appointments, err := r.db.Query(`SELECT 
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
        status,
		created_at,
		updated_at
    FROM booked_appointments
    WHERE doctor_id = $1 AND deleted_at IS NULL`, doctorID.Id)
	if err != nil {
		return nil, err
	}
	defer appointments.Close()

	bookedAppointments := &pb.GetBookedAppointments{}

	for appointments.Next() {
		var appointment pb.BookedAppointment
		var updatedAtPtr *string
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
			&appointment.Status,
			&appointment.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}

		if updatedAtPtr != nil {
			appointment.UpdateAt = *updatedAtPtr
		}

		bookedAppointments.BookedAppointments = append(bookedAppointments.BookedAppointments, &appointment)
	}

	return bookedAppointments, nil
}

func (r *bookingRepo) GetBookedAppointmentsByPatientID(doctorID *pb.GetRequest) (*pb.GetBookedAppointments, error) {
	appointments, err := r.db.Query(`SELECT 
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
        status,
		created_at,
		updated_at
    FROM booked_appointments
    WHERE patient_id = $1 AND deleted_at IS NULL`, doctorID.Id)
	if err != nil {
		return nil, err
	}
	defer appointments.Close()

	bookedAppointments := &pb.GetBookedAppointments{}

	for appointments.Next() {
		var appointment pb.BookedAppointment
		var updatedAtPtr *string
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
			&appointment.Status,
			&appointment.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}
		if updatedAtPtr != nil {
			appointment.UpdateAt = *updatedAtPtr
		}
		
		bookedAppointments.BookedAppointments = append(bookedAppointments.BookedAppointments, &appointment)
	}

	return bookedAppointments, nil
}

func (r *bookingRepo) UpdatePatientStatusByToken(req *pb.UpdRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRow(`
        UPDATE booked_appointments
        	SET 
				appointment_date = $1, 
				appointment_time = $2, 
				type = $3, 
				duration = $4, 
				expires_at = $5, 
				token = $6, 
				patient_status = $7, 
				status = $8,
				updated_at = $9
        	WHERE token = $10 AND deleted_at IS NULL
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
				status,
				created_at,
				updated_at
     `, 
		req.BookedAppointment.AppointmentDate,
		req.BookedAppointment.AppointmentTime,
		req.BookedAppointment.Type,
		req.BookedAppointment.Duration,
		req.BookedAppointment.ExpiresAt,
		req.BookedAppointment.Token,
		req.BookedAppointment.PatientStatus,
		req.BookedAppointment.Status,
		time.Now(),
		req.Token,
		).Scan(
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
			&ba.CreateAt,
			&ba.UpdateAt,
		)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) UpdateBookedAppointment(req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	var ba pb.BookedAppointment
	err := r.db.QueryRow(`
        UPDATE booked_appointments
        	SET 
				appointment_date = $1, 
				appointment_time = $2, 
				type = $3, 
				duration = $4, 
				expires_at = $5, 
				token = $6, 
				patient_status = $7, 
				status = $8,
				updated_at = $9
        	WHERE patient_id = $10 AND deleted_at IS NULL
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
				status,
				created_at,
				updated_at
     `, 
		req.BookedAppointment.AppointmentDate,
		req.BookedAppointment.AppointmentTime,
		req.BookedAppointment.Type,
		req.BookedAppointment.Duration,
		req.BookedAppointment.ExpiresAt,
		req.BookedAppointment.Token,
		req.BookedAppointment.PatientStatus,
		req.BookedAppointment.Status,
		time.Now(),
		req.Id,
		).Scan(
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
			&ba.CreateAt,
			&ba.UpdateAt,
		)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &ba, nil
}

func (r *bookingRepo) DeleteBookedAppointment(req *pb.GetRequest) (bool, error) {
	result, err := r.db.Exec(`
		UPDATE booked_appointments
        SET deleted_at = $1
        WHERE patient_id = $2 AND deleted_at IS NULL
	`, time.Now(), req.Id)
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


