package services

import (
	"context"
	pb "Booking-service/genproto/booking-service"
)

func (s *BookingService) CreateDoctorNote(ctx context.Context, request *pb.CreateDoctorNoteReq) (*pb.DoctorNote, error) {

	doctorNote, err := s.storage.Booking().CreateDoctorNote(request)
    if err != nil {
        return nil, err
    }
    return doctorNote, nil
}

func (s *BookingService) GetDoctorNote(ctx context.Context, request *pb.GetDoctorNoteReq) (*pb.DoctorNote, error) {
	doctorNote, err := s.storage.Booking().GetDoctorNote(request)
    if err != nil {
        return nil, err
    }
    return doctorNote, nil
}

func (s *BookingService) GetDoctorNotesByPatienId(ctx context.Context, request *pb.GetDoctorNoteReq) (*pb.DoctorNotes, error) {
    doctorNotes, err := s.storage.Booking().GetDoctorNotesByPatienId(request)
    if err != nil {
        return nil, err
    }
    return doctorNotes, nil
}

func (s *BookingService) UpdateDoctorNote(ctx context.Context, request *pb.UpdateDoctorNoteReq) (*pb.DoctorNote, error) {
    doctorNote, err := s.storage.Booking().UpdateDoctorNote(request)
    if err != nil {
        return nil, err
    }
    return doctorNote, nil
}

func (s *BookingService) DeleteDoctorNote(ctx context.Context, request *pb.GetDoctorNoteReq) (*pb.Status, error) {
    t, err := s.storage.Booking().DeleteDoctorNote(request)
	if err != nil {
		return nil, err
	}
	
	return &pb.Status{Status: t}, nil
}
