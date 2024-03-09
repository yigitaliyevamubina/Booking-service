package services

import (
	"context"
	pb "Booking-service/genproto/booking-service"
)

func (s *BookingService) CreateDoctorAvailability(ctx context.Context, req *pb.CreateDoctorAvailabilitys) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().CreateDoctorAvailability(req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityById) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().GetDoctorAvailability(req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) GetDoctorAvailabilityByDoctorId(ctx context.Context, req *pb.GetDoctorAvailabilityById) (*pb.DoctorAvailabilitys, error) {
	doctorAvailability, err := s.storage.Booking().GetDoctorAvailabilityByDoctorId(req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityById) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().UpdateDoctorAvailability(req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) DeleteDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityById) (del *pb.Status, err error) {
	t, err := s.storage.Booking().DeleteDoctorAvailability(req)
	if err != nil {
		return nil, err
	}
	
	return &pb.Status{Status: t}, nil
}
