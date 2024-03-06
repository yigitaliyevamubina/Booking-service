package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
	"errors"
	"fmt"
)

func (s *BookingService) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().CreatePatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to create patient in the database:")
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to create patient: returned patient is nil")
	}
	return patient, nil
}

func (s *BookingService) GetPatient(ctx context.Context, req *pb.GetPatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().GetPatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to fetch patient from the database:")
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to get patient: patient not found")
	}
	return patient, nil
}

func (s *BookingService) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().UpdatePatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update patient in the database:")
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to update patient: patient not found")
	}
	return patient, nil
}

func (s *BookingService) DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (del *pb.IsDeleted, err error) {
	_, err = s.storage.Booking().DeletePatient(ctx, req)
	if err != nil {
		del.IsDeleted = false
		return del, fmt.Errorf("failed to delete patient: %w", err)
	}
	del.IsDeleted = true

	return del, nil
}
