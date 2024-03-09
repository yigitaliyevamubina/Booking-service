package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "Booking-service/genproto/booking-service"
)

func (r *bookingRepo) CreateDoctorAvailability(req *pb.CreateDoctorAvailabilitys) (resp *pb.DoctorAvailability, err error) {

	_, err = r.db.Exec(`
		INSERT INTO doctor_availability (
			id,
			doctor_id, 
			department_id, 
			availability_date, 
			availability_time, 
			status,
			created_at
		)VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, req.Id,
		req.DoctorId,
		req.DepartmentId,
		req.AvailabilityDate,
		req.AvailabilityTime,
		req.Status,
		time.Now())

	if err != nil {
		return nil, err
	}
	resp, err = r.GetDoctorAvailability(&pb.GetDoctorAvailabilityById{Id: req.DoctorId})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *bookingRepo) GetDoctorAvailability(req *pb.GetDoctorAvailabilityById) (*pb.DoctorAvailability, error) {
	var availability pb.DoctorAvailability
	var updatedAtPtr *string
	err := r.db.QueryRow(`
			SELECT	
				id, 
				doctor_id, 
				department_id, 
				availability_date, 
				availability_time, 
				status,
				created_at,
				updated_at
		FROM doctor_availability
		WHERE doctor_id = $1 AND deleted_at IS NULL
	`, req.Id).Scan(
		&availability.Id,
		&availability.DoctorId,
		&availability.DepartmentId,
		&availability.AvailabilityDate,
		&availability.AvailabilityTime,
		&availability.Status,
		&availability.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("doctor availability not found: %v", err)
		}
		return nil, fmt.Errorf("failed to get doctor availability: %v", err)
	}

	if updatedAtPtr != nil {
		availability.UpdateAt = *updatedAtPtr
	}
	return &availability, nil
}

func (r *bookingRepo) GetDoctorAvailabilityByDoctorId(req *pb.GetDoctorAvailabilityById) (resp *pb.DoctorAvailabilitys, err error) {
	rows, err := r.db.Query(`
        SELECT 
            id, 
            doctor_id, 
            department_id, 
            availability_date, 
            availability_time, 
            status,
			created_at,
			updated_at
        FROM 
            doctor_availability
        WHERE 
            doctor_id = $1 AND deleted_at IS NULL
    `, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var availabilities []*pb.DoctorAvailability
	for rows.Next() {
		var availability pb.DoctorAvailability
		var updatedAtPtr *string
		err := rows.Scan(
			&availability.Id,
			&availability.DoctorId,
			&availability.DepartmentId,
			&availability.AvailabilityDate,
			&availability.AvailabilityTime,
			&availability.Status,
			&availability.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}

		if updatedAtPtr != nil {
			availability.UpdateAt = *updatedAtPtr
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
	fmt.Println(req)
	_, err := r.db.Exec(`
		UPDATE doctor_availability
			SET
			availability_date = $1,
			availability_time = $2,
			status = $3,
			updated_at = $4
		WHERE doctor_id = $5 AND deleted_at IS NULL
	`,  req.DoctorAvailability.AvailabilityDate,
		req.DoctorAvailability.AvailabilityTime,
		req.DoctorAvailability.Status,
		time.Now(),
		req.Id,
	)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to update doctor availability: %v", err)
	}

	updatedAvailability, err := r.GetDoctorAvailability(&pb.GetDoctorAvailabilityById{Id: req.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated doctor availability: %v", err)
	}

	return updatedAvailability, nil
}

func (r *bookingRepo) DeleteDoctorAvailability(req *pb.GetDoctorAvailabilityById) (bool, error) {

	result, err := r.db.Exec(`
		UPDATE doctor_availability
        SET deleted_at = $1
        WHERE doctor_id = $2 AND deleted_at IS NULL
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
