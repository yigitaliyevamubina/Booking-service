package services

import (
	"context"
	pb "Booking-service/genproto/booking-service"
)

// DoctorAvailabilityService Methods
func (s *BookingService) CreateDoctorAvailability(ctx context.Context, req *pb.CreateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().CreateDoctorAvailability(ctx, req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityRequest1) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().GetDoctorAvailability(ctx, req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().UpdateDoctorAvailability(ctx, req)
	if err != nil {
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) DeleteDoctorAvailability(ctx context.Context, req *pb.DeleteDoctorAvailabilityRequest) (del *pb.IsDeleted, err error) {
	_, err = s.storage.Booking().DeleteDoctorAvailability(ctx, req)
	if err != nil {
		del.IsDeleted = false
		return nil, err
	}
	del.IsDeleted = true
	
	return del, nil
}
