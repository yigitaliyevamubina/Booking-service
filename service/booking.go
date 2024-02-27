package service

import (
	pb "Booking-service/genproto/booking-service"
	l "Booking-service/pkg/logger"
	"Booking-service/storage"
	"context"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BookingService struct {
	storage storage.IStorage
	logger  l.Logger
	pb.UnimplementedPatientServiceServer
	pb.UnimplementedDoctorAvailabilityServiceServer
	pb.UnimplementedBookedAppointmentServiceServer
	// pb.UnimplementedArchiveServiceServer
	// pb.UnimplementedAuthenticationServiceServer
	// pb.UnimplementedBillingServiceServer
	// pb.UnimplementedBookedAppointmentServiceServer
	// pb.UnimplementedDoctorNoteServiceServer
	// pb.UnimplementedFeedbackServiceServer
	// pb.UnimplementedNotificationServiceServer
	// pb.UnimplementedPatientPaymentServiceServer
	// pb.UnimplementedReportingServiceServer
	// pb.UnimplementedSearchServiceServer
	// pb.UnimplementedUploadedFileServiceServer
}

func NewBookingService(db *sqlx.DB, log l.Logger) *BookingService {
	return &BookingService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

// PatientService Methods

// PatientService Methods

func (s *BookingService) CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().CreatePatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to create patient in the database:")
		return nil, fmt.Errorf("failed to create patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to create patient: returned patient is nil")
	}
	return patient, nil
}

func (s *BookingService) GetPatient(ctx context.Context, req *pb.GetPatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().GetPatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to fetch patient from the database:")
		return nil, fmt.Errorf("failed to get patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to get patient: patient not found")
	}
	return patient, nil
}

func (s *BookingService) UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().UpdatePatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to update patient in the database:")
		return nil, fmt.Errorf("failed to update patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to update patient: patient not found")
	}
	return patient, nil
}

func (s *BookingService) DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.Patient, error) {
	patient, err := s.storage.Booking().DeletePatient(ctx, req)
	if err != nil {
		s.logger.Error("Failed to delete patient from the database:")
		return nil, fmt.Errorf("failed to delete patient: %w", err)
	}
	if patient == nil {
		return nil, errors.New("failed to delete patient: patient not found")
	}
	return patient, nil
}

// DoctorAvailabilityService Methods

func (s *BookingService) CreateDoctorAvailability(ctx context.Context, req *pb.CreateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().CreateDoctorAvailability(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityRequest1) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().GetDoctorAvailability(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().UpdateDoctorAvailability(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return doctorAvailability, nil
}

func (s *BookingService) DeleteDoctorAvailability(ctx context.Context, req *pb.DeleteDoctorAvailabilityRequest) (*pb.DoctorAvailability, error) {
	doctorAvailability, err := s.storage.Booking().DeleteDoctorAvailability(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return doctorAvailability, nil
}

// // BookedAppointmentService Methods
func (s *BookingService) CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().CreateBookedAppointment(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return BookedAppointment, nil
}

func (s *BookingService) GetBookedAppointment(ctx context.Context, req *pb.GetBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().GetBookedAppointment(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return BookedAppointment, nil
}

func (s *BookingService) UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().UpdateBookedAppointment(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return BookedAppointment, nil
}

func (s *BookingService) DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (*pb.BookedAppointment, error) {
	BookedAppointment, err := s.storage.Booking().DeleteBookedAppointment(ctx, req)
	if err != nil {
		// Handle error
		return nil, err
	}
	return BookedAppointment, nil
}

// // ArchiveService Methods

// func (s *BookingService) CreateArchive(ctx context.Context, req *pb.CreateArchiveRequest) (*pb.Archive, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) GetArchive(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archive, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) UpdateArchive(ctx context.Context, req *pb.UpdateArchiveRequest) (*pb.Archive, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) DeleteArchive(ctx context.Context, req *pb.DeleteArchiveRequest) (*pb.Archive, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // AuthenticationService Methods

// func (s *BookingService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // BillingService Methods

// func (s *BookingService) GenerateInvoice(ctx context.Context, req *pb.GenerateInvoiceRequest) (*pb.GenerateInvoiceResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // DoctorNoteService Methods

// func (s *BookingService) CreateDoctorNote(ctx context.Context, req *pb.CreateDoctorNoteRequest) (*pb.DoctorNote, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) GetDoctorNote(ctx context.Context, req *pb.GetDoctorNoteRequest) (*pb.DoctorNote, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) UpdateDoctorNote(ctx context.Context, req *pb.UpdateDoctorNoteRequest) (*pb.DoctorNote, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) DeleteDoctorNote(ctx context.Context, req *pb.DeleteDoctorNoteRequest) (*pb.DoctorNote, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // FeedbackService Methods

// func (s *BookingService) SubmitFeedback(ctx context.Context, req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // NotificationService Methods

// func (s *BookingService) SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // PatientPaymentService Methods

// func (s *BookingService) MakePayment(ctx context.Context, req *pb.MakePaymentRequest) (*pb.PatientPayment, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PatientPayment, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.PatientPayment, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.PatientPayment, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // ReportingService Methods

// func (s *BookingService) GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // SearchService Methods

// func (s *BookingService) SearchDoctors(ctx context.Context, req *pb.SearchDoctorsRequest) (*pb.SearchDoctorsResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) SearchPatients(ctx context.Context, req *pb.SearchPatientsRequest) (*pb.SearchPatientsResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) SearchAppointments(ctx context.Context, req *pb.SearchAppointmentsRequest) (*pb.SearchAppointmentsResponse, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// // UploadedFileService Methods

// func (s *BookingService) UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadedFile, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.UploadedFile, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UploadedFile, error) {
// 	// Implement your logic here
// 	return nil, nil
// }

// func (s *BookingService) DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.UploadedFile, error) {
// 	// Implement your logic here
// 	return nil, nil
// }
