package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
)

// // BookedAppointmentService Methods
func (s *BookingService) CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().CreateBookedAppointment(ctx, req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) GetBookedAppointment(ctx context.Context, req *pb.GetBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().GetBookedAppointment(ctx, req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().UpdateBookedAppointment(ctx, req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (del *pb.IsDeleted, err error) {
	_, err = s.storage.Booking().DeleteBookedAppointment(ctx, req)
	if err != nil {
		del.IsDeleted = false
		return del, err
	}
	del.IsDeleted = true
	
	return del, nil
}
