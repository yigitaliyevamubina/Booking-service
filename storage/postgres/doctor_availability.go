package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateDoctorAvailability(req *pb.CreateDoctorAvailabilitys) (resp *pb.DoctorAvailability, err error) {

	query := `
		INSERT INTO doctor_availability (
			doctor_id, 
			department_id, 
			availability_date, 
			availability_time, 
			status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING 
			id, 
			doctor_id, 
			department_id, 
			availability_date, 
			availability_time, 
			status
	`
	resp = &pb.DoctorAvailability{}
	err = r.db.QueryRow(query, req.DoctorId,
		req.DepartmentId,
		req.AvailabilityDate,
		req.AvailabilityTime,
		req.Status).Scan(
		&resp.Id, &resp.DoctorId, &resp.DepartmentId,
		&resp.AvailabilityDate, &resp.AvailabilityTime, &resp.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to create doctor availability: %v", err)
	}

	return resp, nil
}

func (r *bookingRepo) GetDoctorAvailability(req *pb.GetDoctorAvailabilityById) (*pb.DoctorAvailability, error) {
	var availability pb.DoctorAvailability
	err := r.db.QueryRow(`
		SELECT	id, 
				doctor_id, 
				department_id, 
				availability_date, 
				availability_time, 
				status
		FROM doctor_availability
		WHERE doctor_id = $1 AND deleted_at IS NULL
	`, req.Id).Scan(
		&availability.Id,
		&availability.DoctorId,
		&availability.DepartmentId,
		&availability.AvailabilityDate,
		&availability.AvailabilityTime, &availability.Status)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("doctor availability not found: %v", err)
		}
		return nil, fmt.Errorf("failed to get doctor availability: %v", err)
	}

	return &availability, nil
}

func (r *bookingRepo) GetDoctorAvailabilityByDoctorId(req *pb.GetDoctorAvailabilityById) (*pb.DoctorAvailabilitys, error) {
	rows, err := r.db.Query(`
        SELECT 
            id, 
            doctor_id, 
            department_id, 
            availability_date, 
            availability_time, 
            status
        FROM 
            doctor_availability
        WHERE 
            doctor_id = $1
    `, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var availabilities []*pb.DoctorAvailability
	for rows.Next() {
		var availability pb.DoctorAvailability
		err := rows.Scan(
			&availability.Id,
			&availability.DoctorId,
			&availability.DepartmentId,
			&availability.AvailabilityDate,
			&availability.AvailabilityTime,
			&availability.Status,
		)
		if err != nil {
			return nil, err
		}
		availabilities = append(availabilities, &availability)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := &pb.DoctorAvailabilitys{
		DoctorAvailability: availabilities,
	}
	return response, nil
}

func (r *bookingRepo) UpdateDoctorAvailability(req *pb.UpdateDoctorAvailabilityById) (*pb.DoctorAvailability, error) {

	_, err := r.db.Exec(`
		UPDATE doctor_availability
			SET id = $2, 
			doctor_id = $3, 
			department_id = $4, 
			availability_date = $5,
			availability_time = $6,
			status = $7
		WHERE doctor_id = $1 AND deleted_at IS NULL
	`, req.Id,
		req.DoctorAvailability.Id,
		req.DoctorAvailability.DoctorId,
		req.DoctorAvailability.DepartmentId,
		req.DoctorAvailability.AvailabilityDate,
		req.DoctorAvailability.AvailabilityTime,
		req.DoctorAvailability.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to update doctor availability: %v", err)
	}

	updatedAvailability, err := r.GetDoctorAvailability(&pb.GetDoctorAvailabilityById{Id: req.DoctorAvailability.DoctorId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated doctor availability: %v", err)
	}

	return updatedAvailability, nil
}

func (r *bookingRepo) DeleteDoctorAvailability(req *pb.GetDoctorAvailabilityById) (bool, error) {

	result, err := r.db.Exec(`
		UPDATE doctor_availability
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
