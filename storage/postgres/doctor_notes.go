package postgres

import (
	"database/sql"
	"errors"
	"time"

	pb "Booking-service/genproto/booking-service"
)

func (r *bookingRepo) CreateDoctorNote(req *pb.CreateDoctorNoteReq) (*pb.DoctorNote, error) {
	query := `
        INSERT INTO doctor_notes(
			id,
			appointment_id, 
			doctor_id, 
			patient_id, 
			note_type, 
			note_text, 
			created_at
		)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING 
			id, 
			appointment_id, 
			doctor_id, 
			patient_id, 
			note_type, 
			note_text, 
			created_at
    `
	row := r.db.QueryRow(
		query,
		req.Id,
		req.AppointmentId,
		req.DoctorId,
		req.PatientId,
		req.NoteType,
		req.NoteText,
		time.Now(),
	)
	var note pb.DoctorNote
	err := row.Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.NoteType,
		&note.NoteText,
		&note.CreateAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *bookingRepo) GetDoctorNote(req *pb.GetDoctorNoteReq) (*pb.DoctorNote, error) {
	var updatedAtPtr *string
	query := `
        SELECT 
			id, 
			appointment_id, 
			doctor_id, 
			patient_id, 
			note_type, 
			note_text, 
			created_at,
			updated_at
        FROM doctor_notes
        WHERE doctor_id = $1 AND deleted_at IS NULL
    `
	row := r.db.QueryRow(query, req.Id)

	var note pb.DoctorNote
	err := row.Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.NoteType,
		&note.NoteText,
		&note.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		return nil, err
	}	

	if updatedAtPtr != nil {
		note.UpdateAt = *updatedAtPtr
	}

	return &note, nil
}

func (r *bookingRepo) GetDoctorNotesByPatienId(req *pb.GetDoctorNoteReq) (*pb.DoctorNotes, error) {
	query := `
        SELECT 
			id, 
			appointment_id, 
			doctor_id, 
			patient_id, 
			note_type, 
			note_text, 
			created_at,
			updated_at
        FROM doctor_notes
        WHERE doctor_id = $1 AND deleted_at IS NULL
    `
	rows, err := r.db.Query(query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []*pb.DoctorNote
	for rows.Next() {
		var note pb.DoctorNote
		var updatedAtPtr *string
		err := rows.Scan(
			&note.Id,
			&note.AppointmentId,
			&note.DoctorId,
			&note.PatientId,
			&note.NoteType,
			&note.NoteText,
			&note.CreateAt,
			&updatedAtPtr,
		)
		if err != nil {
			return nil, err
		}

		if updatedAtPtr != nil {
			note.UpdateAt = *updatedAtPtr
		}
		notes = append(notes, &note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &pb.DoctorNotes{DoctorNote: notes}, nil
}

func (r *bookingRepo) UpdateDoctorNote(req *pb.UpdateDoctorNoteReq) (*pb.DoctorNote, error) {

	query := `
        UPDATE doctor_notes
        	SET 
				note_type = $1,
				note_text = $2, 
				updated_at = $3
        WHERE doctor_id = $4
        RETURNING 
			id, 
			appointment_id, 
			doctor_id, 
			patient_id, 
			note_type, 
			note_text, 
			created_at, 
			updated_at
    `
	// Execute the SQL statement
	row := r.db.QueryRow(
		query,
		req.DoctorNote.NoteType,
		req.DoctorNote.NoteText,
		time.Now(),
		req.Id,
	)

	var note pb.DoctorNote
	err := row.Scan(
		&note.Id,
		&note.AppointmentId,
		&note.DoctorId,
		&note.PatientId,
		&note.NoteType,
		&note.NoteText,
		&note.CreateAt,
		&note.UpdateAt,
	)
	if err != nil {
		return nil, err
	}

	return &note, nil
}

func (r *bookingRepo) DeleteDoctorNote(req *pb.GetDoctorNoteReq) (bool, error) {

	result, err := r.db.Exec(`
		UPDATE doctor_notes
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
