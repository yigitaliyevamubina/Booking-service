package services

import (
	pb "Booking-service/genproto/booking-service"
	l "Booking-service/pkg/logger"
	"Booking-service/storage"

	"github.com/jmoiron/sqlx"
)

type BookingService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedPatientServiceServer
	pb.UnimplementedDoctorAvailabilityServiceServer
	pb.UnimplementedBookedAppointmentServiceServer
	pb.UnimplementedArchiveServiceServer
	pb.UnimplementedPatientPaymentServiceServer
	pb.UnimplementedUploadedFileServiceServer
	pb.UnimplementedDoctorNoteServiceServer
}

func NewBookingService(db *sqlx.DB, log l.Logger) *BookingService {
	return &BookingService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}
