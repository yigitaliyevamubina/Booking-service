package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
)

func (s *BookingService) CreatePatient(ctx context.Context, req *pb.Patient) (*pb.Patient, error) {
	patient, err := s.storage.Booking().CreatePatient(req)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, err
	}
	return patient, nil
}

func (s *BookingService) GetPatient(ctx context.Context, req *pb.GetPatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().GetPatient(req)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, err
	}
	return patient, nil
}

func (s *BookingService) GetPatients(ctx context.Context, req *pb.PatientsReq) (*pb.Patients, error) {
	patient, err := s.storage.Booking().GetPatients(req)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, err
	}
	return patient, nil
}
func (s *BookingService) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().UpdatePatient(req)
	if err != nil {
		return nil, err
	}
	if patient == nil {
		return nil, err
	}
	return patient, nil
}

func (s *BookingService) DeletePatient(ctx context.Context, req *pb.GetPatientRequest) (del *pb.Status, err error) {
	t, err := s.storage.Booking().DeletePatient(req)
	if err != nil {
		return nil, err
	}

	return &pb.Status{Status: t}, nil
}
