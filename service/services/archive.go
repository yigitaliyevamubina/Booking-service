package services

import (
	pb "Booking-service/genproto/booking-service"
	"context"
)

// GetArchive retrieves an archive entry
func (s *BookingService) GetArchive(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archive, error) {
	archive, err := s.storage.Booking().GetArchive(ctx, req)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

func (s *BookingService) GetArchiveByPatientID(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archives, error) {

	archive, err := s.storage.Booking().GetArchiveByPatientID(ctx, req)
	if err != nil {
		return nil, err
	}

	return archive, nil

}

// UpdateArchive updates an existing archive entry
func (s *BookingService) UpdateArchive(ctx context.Context, req *pb.UpdateArchiveRequest) (*pb.Archive, error) {
	archive, err := s.storage.Booking().UpdateArchive(ctx, req)
	if err != nil {
		return nil, err
	}

	return archive, nil
}

// DeleteArchive deletes an archive entry
func (s *BookingService) DeleteArchive(ctx context.Context, req *pb.DeleteArchiveRequest) (del *pb.Status, err error) {
	if _, err := s.storage.Booking().DeleteArchive(ctx, req); err != nil {
		return nil, err
	}
	del.Status = true

	return del, nil
}
