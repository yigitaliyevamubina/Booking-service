package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "Booking-service/genproto/booking-service"
)

func (r *bookingRepo) UploadFile(request *pb.CreateUploadedFile) (*pb.UploadedFile, error) {
	resp := &pb.UploadedFile{}
	query := `
		INSERT INTO uploaded_files (
			file_id, 
			patient_id, 
			request_id,
			file, 
			created_at
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING 
			file_id, 
			patient_id, 
			request_id, 
			file,
			created_at
	`
	err := r.db.QueryRow(
		query,
		request.FileId,
		request.PatientId,
		request.RequestId,
		request.File,
		time.Now()).Scan(
		&resp.FileId,
		&resp.PatientId,
		&resp.RequestId,
		&resp.File,
		&resp.CreateAt,
	)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *bookingRepo) GetFileByPatientID(request *pb.GetFileRequest) (*pb.UploadedFile, error) {
	var updatedAtPtr *string
	resp := &pb.UploadedFile{} 

	query := `
        SELECT 
            file_id, 
            patient_id, 
            request_id, 
            file,
            created_at,
            updated_at
        FROM uploaded_files
        WHERE patient_id = $1 AND deleted_at IS NULL
    `
	err := r.db.QueryRow(query, request.FileId).Scan(
		&resp.FileId,
		&resp.PatientId,
		&resp.RequestId,
		&resp.File,
		&resp.CreateAt,
		&updatedAtPtr,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if updatedAtPtr != nil {
		resp.UpdateAt = *updatedAtPtr
	}
	return resp, nil
}

func (r *bookingRepo) GetFilesByPatientID(request *pb.GetFileRequest) (*pb.UploadedFiles, error) {
	query := `
		SELECT 
			file_id, 
			patient_id, 
			request_id,
			file,
			created_at,
			updated_at
		FROM uploaded_files
		WHERE patient_id = $1 AND deleted_at IS NULL
	`
	rows, err := r.db.Query(query, request.FileId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []*pb.UploadedFile
	for rows.Next() {
		var file pb.UploadedFile
		var updatedAtPtr *string
		if err := rows.Scan(
			&file.FileId,
			&file.PatientId,
			&file.RequestId,
			&file.File,
			&file.CreateAt,
			&updatedAtPtr,
		); err != nil {
			return nil, err
		}

		if updatedAtPtr != nil {
			file.UpdateAt = *updatedAtPtr
		}

		files = append(files, &file)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	response := &pb.UploadedFiles{
		Uploaded: files,
	}
	return response, nil
}

func (r *bookingRepo) UpdateFile(request *pb.UpdateFileRequest) (resp *pb.UploadedFile, err error) {
	fmt.Println(request.UploadedFile)
	query := `
		UPDATE uploaded_files
		SET 
			file = $1, 
			updated_at = $2
		WHERE patient_id = $3 AND deleted_at IS NULL
	`
	_, err = r.db.Exec(query,
		request.UploadedFile.File,
		time.Now(), 
		request.PatientId,
	)

	 if err != nil {
		return nil, err
	 }

	resp, err = r.GetFileByPatientID(&pb.GetFileRequest{FileId: request.PatientId})

	if err != nil {
		return nil, err
	}
	 
	return resp, nil
}

func (r *bookingRepo) DeleteFile(request *pb.GetFileRequest) (bool, error) {

	result, err := r.db.Exec(`
		UPDATE uploaded_files
		SET deleted_at = $1
		WHERE patient_id = $2 AND deleted_at IS NULL
	`, time.Now(), request.FileId)
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
