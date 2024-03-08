package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	pb "Booking-service/genproto/booking-service"

	_ "github.com/lib/pq"
	"github.com/spf13/cast"
)


func (r *bookingRepo) CreateArchive(ctx context.Context, req *pb.InsertArchive) (*pb.Archive, error) {
	resp := &pb.Archive{}
	err := r.db.QueryRowContext(ctx, `
    INSERT INTO archive (
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
		created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
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
		req.Insert.DepartmentId,
		req.Insert.DoctorId,
		req.Insert.PatientId,
		req.Insert.PatientToken,
		req.Insert.PatientProblem,
		req.Insert.ConsultationType,
		req.Insert.BookedDate,
		req.Insert.BookedTime,
		req.Insert.AppointmentId,
		cast.ToString(req.Insert.Status),
		req.Insert.VisitsCount, time.Now()).Scan(&resp.Id,
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

func (r *bookingRepo) GetArchive(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archive, error) {

	query := `
		SELECT id, 
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
		FROM archive
		WHERE patient_id = $1
		AND deleted_at IS NULL`
	var archive pb.Archive
	err := r.db.QueryRowContext(ctx, query, req.Id).Scan(
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
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("archive not found")
		}
		return nil, err
	}

	return &archive, nil
}


func (r *bookingRepo) GetArchiveByPatientID(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archives, error) {
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
            visits_count
        FROM 
            archive
        WHERE 
            patient_id = $1
    `
    rows, err := r.db.QueryContext(ctx, query, req.Id)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var archives []*pb.Archive
    for rows.Next() {
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
        )
        if err != nil {
            return nil, err
        }
        archives = append(archives, &archive)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }

    return &pb.Archives{Archives: archives}, nil
}

func (r *bookingRepo) UpdateArchive(ctx context.Context, req *pb.UpdateArchiveRequest) (*pb.Archive, error) {

	query := `
		UPDATE archive
		SET id = $2, 
		department_id = $3,
		doctor_id = $4, 
		patient_id = $5, 
		patient_token = $6, 
		patient_problem = $7, 
		consultation_type = $8, 
		booked_date = $9, 
		booked_time = $10, 
		appointment_id = $11, 
		status = $12, 
		visits_count = $13, 
		updated_at = $14
		WHERE patient_id = $1 AND deleted_at IS NULL
		RETURNING id, 
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
		updated_at`
	var archive pb.Archive
	err := r.db.QueryRowContext(ctx, query,
		req.Id,
		req.Archive.Id,
		req.Archive.DepartmentId,
		req.Archive.DoctorId,
		req.Archive.PatientId,
		req.Archive.PatientToken,
		req.Archive.PatientProblem,
		req.Archive.ConsultationType,
		req.Archive.BookedDate,
		req.Archive.BookedTime,
		req.Archive.AppointmentId,
		req.Archive.Status,
		req.Archive.VisitsCount, time.Now()).Scan(
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

func (r *bookingRepo) DeleteArchive(ctx context.Context, req *pb.DeleteArchiveRequest) (del *pb.Status, err error) {

	query := `UPDATE archive
			SET deleted_at = $1
			WHERE patient_id = $2`
	_, err = r.db.Exec(query, time.Now(), req.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("archive not found")
		}
		return nil, err
	}
	del.Status = true
	
	return del, nil
}
