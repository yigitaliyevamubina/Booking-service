package postgres

import (
	"database/sql"
	"errors"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
)

func (r *bookingRepo) CreateArchive(req *pb.CreateArchiveReq) (*pb.Archive, error) {
	resp := &pb.Archive{}
	err := r.db.QueryRow(`
    INSERT INTO archive (
		id,
		department_id, 
		doctor_id, 
		patient_id, 
		patient_token, 
		patient_problem, 
		consultation_type, 
		booked_date, 
		booked_time, 
		appointment_id, 
		status, 
		visits_count, 
		created_at
	)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
	RETURNING 
		id, 
		department_id, 
		doctor_id, 
		patient_id, 
		patient_token, 
		patient_problem, 
		consultation_type, 
		booked_date, 
		booked_time, 
		appointment_id, 
		status, 
		visits_count, 
		created_at`,
		req.Id,
		req.DepartmentId,
		req.DoctorId,
		req.PatientId,
		req.PatientToken,
		req.PatientProblem,
		req.ConsultationType,
		req.BookedDate,
		req.BookedTime,
		req.AppointmentId,
		req.Status,
		req.VisitsCount,
		time.Now()).Scan(
		&resp.Id,
		&resp.DepartmentId,
		&resp.DoctorId,
		&resp.PatientId,
		&resp.PatientToken,
		&resp.PatientProblem,
		&resp.ConsultationType,
		&resp.BookedDate,
		&resp.BookedTime,
		&resp.AppointmentId,
		&resp.Status,
		&resp.VisitsCount,
		&resp.CreateAt)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *bookingRepo) GetArchive(req *pb.GetArchiveReq) (*pb.Archive, error) {
	var updatedAtPtr *string
	query := `
		SELECT 
		id, 
		department_id, 
		doctor_id, 
		patient_id, 
		patient_token, 
		patient_problem, 
		consultation_type, 
		booked_date, 
		booked_time, 
		appointment_id, 
		status, 
		visits_count, 
		created_at,
		updated_at
		FROM archive
		WHERE patient_id = $1 AND deleted_at IS NULL`
	var archive pb.Archive
	err := r.db.QueryRow(query, req.Id).Scan(
		&archive.Id,
		&archive.DepartmentId,
		&archive.DoctorId,
		&archive.PatientId,
		&archive.PatientToken,
		&archive.PatientProblem,
		&archive.ConsultationType,
		&archive.BookedDate,
		&archive.BookedTime,
		&archive.AppointmentId,
		&archive.Status,
		&archive.VisitsCount,
		&archive.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("archive not found")
		}
		return nil, err
	}

	if updatedAtPtr != nil {
		archive.UpdateAt = *updatedAtPtr
	}

	return &archive, nil
}

func (r *bookingRepo) GetArchivesByPatientID(req *pb.GetArchiveReq) (*pb.Archives, error) {
	query := `
        SELECT 
            id, 
            department_id, 
            doctor_id, 
            patient_id, 
            patient_token, 
            patient_problem, 
            consultation_type, 
            booked_date, 
            booked_time, 
            appointment_id, 
            status, 
            visits_count,
			created_at,
			updated_at
        FROM 
            archive
        WHERE 
            patient_id = $1 AND deleted_at IS NULL
    `
	rows, err := r.db.Query(query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var archives []*pb.Archive
	for rows.Next() {
		var updatedAtPtr *string
		var archive pb.Archive
		err := rows.Scan(
			&archive.Id,
			&archive.DepartmentId,
			&archive.DoctorId,
			&archive.PatientId,
			&archive.PatientToken,
			&archive.PatientProblem,
			&archive.ConsultationType,
			&archive.BookedDate,
			&archive.BookedTime,
			&archive.AppointmentId,
			&archive.Status,
			&archive.VisitsCount,
			&archive.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}

		if updatedAtPtr != nil {
			archive.UpdateAt = *updatedAtPtr
		}
		archives = append(archives, &archive)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.Archives{Archives: archives}, nil
}

func (r *bookingRepo) UpdateArchive(req *pb.UpdateArchiveRequest) (*pb.Archive, error) {

	query := `
		UPDATE archive
		SET
		patient_token = $1, 
		patient_problem = $2, 
		consultation_type = $3, 
		booked_date = $4, 
		booked_time = $5, 
		appointment_id = $6, 
		status = $7, 
		visits_count = $8, 
		updated_at = $9
		WHERE patient_id = $10 AND deleted_at IS NULL
		RETURNING 
		id, 
		department_id, 
		doctor_id, 
		patient_id, 
		patient_token, 
		patient_problem, 
		consultation_type, 
		booked_date, 
		booked_time, 
		appointment_id, 
		status, 
		visits_count, 
		created_at, 
		updated_at
		`

	var archive pb.Archive
	err := r.db.QueryRow(query,
		req.Archive.PatientToken,
		req.Archive.PatientProblem,
		req.Archive.ConsultationType,
		req.Archive.BookedDate,
		req.Archive.BookedTime,
		req.Archive.AppointmentId,
		req.Archive.Status,
		req.Archive.VisitsCount,
		time.Now(),
		req.Id).Scan(
		&archive.Id,
		&archive.DepartmentId,
		&archive.DoctorId,
		&archive.PatientId,
		&archive.PatientToken,
		&archive.PatientProblem,
		&archive.ConsultationType,
		&archive.BookedDate,
		&archive.BookedTime,
		&archive.AppointmentId,
		&archive.Status,
		&archive.VisitsCount,
		&archive.CreateAt,
		&archive.UpdateAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("archive not found")
		}
		return nil, err
	}

	return &archive, nil
}

func (r *bookingRepo) DeleteArchive(req *pb.GetArchiveReq) (bool, error) {

	result, err := r.db.Exec(`
	UPDATE archive
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
