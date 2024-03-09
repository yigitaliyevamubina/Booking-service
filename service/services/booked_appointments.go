package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
)

func (s *BookingService) CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointments) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().CreateBookedAppointment(req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) GetBookedAppointment(ctx context.Context, req *pb.GetRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().GetBookedAppointment(req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) GetBookedAppointmentsByPatientID(ctx context.Context, patientID *pb.GetRequest) (*pb.GetBookedAppointments, error) {

	BookedAppointment, err := s.storage.Booking().GetBookedAppointmentsByPatientID(patientID)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) GetBookedAppointmentsByDoctorID(ctx context.Context, doctorID *pb.GetRequest) (*pb.GetBookedAppointments, error) {
	BookedAppointment, err := s.storage.Booking().GetBookedAppointmentsByDoctorID(doctorID)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().UpdateBookedAppointment(req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) UpdatePatientStatusByToken(ctx context.Context, req *pb.UpdRequest) (*pb.GetBookedAppointments, error) {
	BookedAppointment, err := s.storage.Booking().UpdatePatientStatusByToken(req)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil
}

func (s *BookingService) DeleteBookedAppointment(ctx context.Context, req *pb.GetRequest) (del *pb.Status, err error) {
    t, err := s.storage.Booking().DeleteBookedAppointment(req)
    if err != nil {
        return nil, err
    }
	
	return &pb.Status{Status: t}, nil
}
