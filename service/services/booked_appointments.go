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

func (s *BookingService) GetBookedAppointmentsByPatientID(ctx context.Context, patientID *pb.PatientID) (*pb.GetBookedAppointmentsByPatientIDResponse, error) {

	BookedAppointment, err := s.storage.Booking().GetBookedAppointmentsByPatientID(ctx, patientID)
	if err != nil {
		return nil, err
	}

	return BookedAppointment, nil

}

func (s *BookingService) GetBookedAppointmentsByDoctorID(ctx context.Context, doctorID *pb.GetBookedAppointmentRequest) (*pb.GetBookedAppointmentsByPatientIDResponse, error) {
	BookedAppointment, err := s.storage.Booking().GetBookedAppointmentsByDoctorID(ctx, doctorID)
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

func (s *BookingService) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (del *pb.Status, err error) {
    _, err = s.storage.Booking().DeleteBookedAppointment(ctx, req)
    if err != nil {
        return nil, err
    }
	del.Status = true
    return del, nil
}

