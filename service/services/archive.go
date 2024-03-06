package services

import (
	"context"
	pb "Booking-service/genproto/booking-service"
)

// GetArchive retrieves an archive entry
func (s *BookingService) GetArchive(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archive, error) {
	archive, err := s.storage.Booking().GetArchive(ctx, req)
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
func (s *BookingService) DeleteArchive(ctx context.Context, req *pb.DeleteArchiveRequest) (del *pb.IsDeleted, err error) {
	if _, err := s.storage.Booking().DeleteArchive(ctx, req); err != nil {
		del.IsDeleted = false
		return del, err
	}
	del.IsDeleted = true

	return del, nil
}
