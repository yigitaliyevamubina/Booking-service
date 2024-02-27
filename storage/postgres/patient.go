package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

// CreatePatient creates a new patient record.
func (r *bookingRepo) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.Patient, error) {

	_, err := r.db.ExecContext(ctx, `
		INSERT INTO patients (id, first_name, last_name, birth_date, gender, city, phone_number)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, req.Patient.Id, req.Patient.FirstName, req.Patient.LastName, req.Patient.BirthDate, req.Patient.Gender, req.Patient.City, req.Patient.PhoneNumber)
	if err != nil {
		log.Println(err)
		return nil, fmt.Errorf("failed to create patient: %v", err)
	}

	return req.Patient, nil
}

func (r *bookingRepo) GetPatient(ctx context.Context, req *pb.GetPatientRequest) (*pb.Patient, error) {

	var patient pb.Patient
	err := r.db.QueryRowContext(ctx, `
		SELECT id, first_name, last_name, birth_date, gender, city, phone_number
		FROM patients
		WHERE id = $1
	`, req.Id).Scan(&patient.Id, &patient.FirstName, &patient.LastName, &patient.BirthDate, &patient.Gender, &patient.City, &patient.PhoneNumber)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("patient not found: %v", err)
		}
		return nil, fmt.Errorf("failed to get patient: %v", err)
	}

	return &patient, nil
}

// UpdatePatient updates an existing patient record.
func (r *bookingRepo) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.Patient, error) {
	_, err := r.db.ExecContext(ctx, `
		UPDATE patients
		SET first_name = $1, last_name = $2, birth_date = $3, gender = $4, city = $5, phone_number = $6
		WHERE id = $7
	`, req.Patient.FirstName, req.Patient.LastName, req.Patient.BirthDate, req.Patient.Gender, req.Patient.City, req.Patient.PhoneNumber, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update patient: %v", err)
	}

	return r.GetPatient(ctx, &pb.GetPatientRequest{Id: req.Id})
}

// DeletePatient deletes a patient record based on the provided patient ID.
func (r *bookingRepo) DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.Patient, error) {

	patient, err := r.GetPatient(ctx, &pb.GetPatientRequest{Id: req.Id})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch patient before deletion: %v", err)
	}
	_, err = r.db.ExecContext(ctx, `
		DELETE FROM patients
		WHERE id = $1
	`, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete patient: %v", err)
	}
	return patient, nil
}
