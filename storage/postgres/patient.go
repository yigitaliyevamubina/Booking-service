package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreatePatient(req *pb.Patient) (resp *pb.Patient, err error) {

	_, err = r.db.Exec(`
		INSERT INTO patients (
			id, 
			first_name, 
			last_name, 
			birth_date,
			gender,
			city, 
			phone_number,
			created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) 
	`, req.Id,
		req.FirstName,
		req.LastName,
		req.BirthDate,
		req.Gender,
		req.City,
		req.PhoneNumber,
		time.Now())
	if err != nil {
		return nil, err
	}

	resp, err = r.GetPatient(&pb.GetPatientRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *bookingRepo) GetPatient(req *pb.GetPatientRequest) (*pb.Patient, error) {

	var patient pb.Patient
	var updatedAtPtr *string
	err := r.db.QueryRow(`
		SELECT 
			id, 
			first_name, 
			last_name, 
			birth_date,
			gender, 
			city, 
			phone_number,
			created_at,
			updated_at
		FROM patients
		WHERE id = $1 AND deleted_at IS NULL
	`, req.Id).Scan(
		&patient.Id,
		&patient.FirstName,
		&patient.LastName,
		&patient.BirthDate,
		&patient.Gender,
		&patient.City,
		&patient.PhoneNumber,
		&patient.CreateAt,
		&updatedAtPtr,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if updatedAtPtr != nil {
		patient.UpdateAt = *updatedAtPtr
	}

	return &patient, nil
}

func (r *bookingRepo) GetPatients(req *pb.GetPatientsRequest) (*pb.Patients, error) {

	limit, err := strconv.ParseInt(req.Limit, 10, 64)
	if err != nil {
		return nil, err
	}
	page, err := strconv.ParseInt(req.Page, 10, 64)
	if err != nil {
		return nil, err
	}
	offset := (page - 1) * limit

	rows, err := r.db.Query(`
        SELECT 
            id, 
            first_name, 
            last_name, 
            birth_date,
            gender, 
            city, 
            phone_number,
			created_at,
			updated_at
        FROM patients
		WHERE deleted_at IS NULL
        ORDER BY id
        LIMIT $1
        OFFSET $2
    `, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []*pb.Patient

	for rows.Next() {
		var patient pb.Patient
		var updatedAtPtr *string
		err := rows.Scan(
			&patient.Id,
			&patient.FirstName,
			&patient.LastName,
			&patient.BirthDate,
			&patient.Gender,
			&patient.City,
			&patient.PhoneNumber,
			&patient.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}

		if updatedAtPtr != nil {
			patient.UpdateAt = *updatedAtPtr
		}

		patients = append(patients, &patient)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := &pb.Patients{
		Patient: patients,
	}

	return response, nil
}

func (r *bookingRepo) UpdatePatient(req *pb.UpdatePatientRequest) (*pb.Patient, error) {
	_, err := r.db.Exec(`
		UPDATE patients
		SET 
			first_name = $1, 
			last_name = $2, 
			birth_date = $3, 
			gender = $4, 
			city = $5, 
			phone_number = $6,
			updated_at = $7
		WHERE id = $8 AND deleted_at IS NULL
	`, req.Patient.FirstName,
		req.Patient.LastName,
		req.Patient.BirthDate,
		req.Patient.Gender,
		req.Patient.City,
		req.Patient.PhoneNumber,
		time.Now(),
		req.Id)
	if err != nil {
		return nil, err
	}

	return r.GetPatient(&pb.GetPatientRequest{Id: req.Id})
}

func (r *bookingRepo) DeletePatient(req *pb.GetPatientRequest) (bool, error) {
	result, err := r.db.Exec(`
        UPDATE patients
        SET deleted_at = $1
        WHERE id = $2 AND deleted_at IS NULL
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
