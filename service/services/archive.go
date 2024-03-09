package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
)

func (s *BookingService) CreateArchive(ctx context.Context, req *pb.CreateArchiveReq) (*pb.Archive, error) {
	archive, err := s.storage.Booking().CreateArchive(req)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (s *BookingService) GetArchive(ctx context.Context, req *pb.GetArchiveReq) (*pb.Archive, error) {
	archive, err := s.storage.Booking().GetArchive(req)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (s *BookingService) GetArchivesByPatientID(ctx context.Context, req *pb.GetArchiveReq) (*pb.Archives, error) {

	archive, err := s.storage.Booking().GetArchivesByPatientID(req)
	if err != nil {
		return nil, err
	}

	return archive, nil

}

func (s *BookingService) UpdateArchive(ctx context.Context, req *pb.UpdateArchiveRequest) (*pb.Archive, error) {
	archive, err := s.storage.Booking().UpdateArchive(req)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (s *BookingService) DeleteArchive(ctx context.Context, req *pb.GetArchiveReq) (del *pb.Status, err error) {
	t, err := s.storage.Booking().DeleteArchive(req)
	if err != nil {
		return nil, err
	}
	return &pb.Status{Status: t}, nil
}
