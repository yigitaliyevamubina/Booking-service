package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateDoctorAvailability(ctx context.Context, req *pb.CreateDoctorAvailabilityRequest) (resp *pb.DoctorAvailability, err error) {

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
	err = r.db.QueryRowContext(ctx, query, req.DoctorAvailability.DoctorId,
		req.DoctorAvailability.DepartmentId,
		req.DoctorAvailability.AvailabilityDate,
		req.DoctorAvailability.AvailabilityTime,
		req.DoctorAvailability.Status).Scan(
		&resp.Id, &resp.DoctorId, &resp.DepartmentId,
		&resp.AvailabilityDate, &resp.AvailabilityTime, &resp.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to create doctor availability: %v", err)
	}

	return resp, nil
}

func (r *bookingRepo) GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityRequest1) (*pb.DoctorAvailability, error) {
	var availability pb.DoctorAvailability
	err := r.db.QueryRowContext(ctx, `
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

func (r *bookingRepo) UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {

	_, err := r.db.ExecContext(ctx, `
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

	updatedAvailability, err := r.GetDoctorAvailability(ctx, &pb.GetDoctorAvailabilityRequest1{Id: req.DoctorAvailability.DoctorId})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch updated doctor availability: %v", err)
	}

	return updatedAvailability, nil
}

func (r *bookingRepo) DeleteDoctorAvailability(ctx context.Context, req *pb.DeleteDoctorAvailabilityRequest) (del *pb.IsDeleted, err error) {

	_, err = r.GetDoctorAvailability(ctx, &pb.GetDoctorAvailabilityRequest1{Id: req.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch doctor availability before deletion: %v", err)
	}

	_, err = r.db.ExecContext(ctx, `
		UPDATE doctor_availability
        SET deleted_at = $1
        WHERE patient_id = $2
	`, time.Now(), req.Id)
	if err != nil {
		del.IsDeleted = false
		return del, fmt.Errorf("failed to delete doctor availability: %v", err)
	}
	del.IsDeleted = true

	return del, nil
}
