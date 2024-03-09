package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
	"errors"
)

func (s *BookingService) UploadFile(ctx context.Context, request *pb.CreateUploadedFile) (*pb.UploadedFile, error) {
	if request == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().UploadFile(request)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) GetFileByPatientID(ctx context.Context, request *pb.GetFileRequest) (*pb.UploadedFile, error) {
	if request == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().GetFileByPatientID(request)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) GetFilesByPatientID(ctx context.Context, request *pb.GetFileRequest) (*pb.UploadedFiles, error) {
	if request == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().GetFilesByPatientID(request)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) UpdateFile(ctx context.Context, request *pb.UpdateFileRequest) (*pb.UploadedFile, error) {
	if request == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().UpdateFile(request)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) DeleteFile(ctx context.Context, request *pb.GetFileRequest) (*pb.Status, error) {
	if request == nil {
		return nil, errors.New("invalid request")
	}

	status, err := s.storage.Booking().DeleteFile(request)
	if err != nil {
		return nil, err
	}

	return &pb.Status{Status: status}, nil
}
