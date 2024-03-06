package service

import (
	"context"

	pb "Booking-service/genproto/booking-service"
)

type BookingStorageI interface {
	// PatientService Methods
	CreatePatient(ctx context.Context, req *pb.CreatePatientRequest) (*pb.Patient, error)
	GetPatient(ctx context.Context, req *pb.GetPatientRequest) (*pb.Patient, error)
	UpdatePatient(ctx context.Context, req *pb.UpdatePatientRequest) (*pb.Patient, error)
	DeletePatient(ctx context.Context, req *pb.DeletePatientRequest) (*pb.IsDeleted, error)
	//DoctorAvailabilityService Methods
	CreateDoctorAvailability(ctx context.Context, req *pb.CreateDoctorAvailabilityRequest) (resp *pb.DoctorAvailability,err error)
	GetDoctorAvailability(ctx context.Context, req *pb.GetDoctorAvailabilityRequest1) (resp *pb.DoctorAvailability, err error)
	UpdateDoctorAvailability(ctx context.Context, req *pb.UpdateDoctorAvailabilityRequest) (*pb.DoctorAvailability, error)
	DeleteDoctorAvailability(ctx context.Context, req *pb.DeleteDoctorAvailabilityRequest) (resp *pb.IsDeleted, err error)
	// BookedAppointmentService Methods
	CreateBookedAppointment(ctx context.Context, req *pb.CreateBookedAppointmentRequest) (*pb.BookedAppointment, error)
	GetBookedAppointment(ctx context.Context, req *pb.GetBookedAppointmentRequest) (*pb.BookedAppointment, error)
	UpdateBookedAppointment(ctx context.Context, req *pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error)
	DeleteBookedAppointment(ctx context.Context, req *pb.DeleteBookedAppointmentRequest) (*pb.IsDeleted, error)
	// ArchiveService Methods
	CreateArchive(ctx context.Context, req *pb.InsertArchive) (*pb.Archive, error)
	GetArchive(ctx context.Context, req *pb.GetArchiveRequest) (*pb.Archive, error)
	UpdateArchive(ctx context.Context, req *pb.UpdateArchiveRequest) (*pb.Archive, error)
	DeleteArchive(ctx context.Context, req *pb.DeleteArchiveRequest) (*pb.IsDeleted, error)
	// PatientPaymentService Methods
	// MakePayment(ctx context.Context, req *pb.MakePaymentRequest) (*pb.PatientPayment, error)
	// GetPayment(ctx context.Context, req *pb.GetPaymentRequest) (*pb.PatientPayment, error)
	// UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.PatientPayment, error)
	// DeletePayment(ctx context.Context, req *pb.DeletePaymentRequest) (*pb.PatientPayment, error)

	// AuthenticationService Methods
	// Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error)
	// Logout(ctx context.Context, req *pb.LogoutRequest) (*pb.LogoutResponse, error)

	// BillingService Methods
	// GenerateInvoice(ctx context.Context, req *pb.GenerateInvoiceRequest) (*pb.GenerateInvoiceResponse, error)
	// ProcessPayment(ctx context.Context, req *pb.ProcessPaymentRequest) (*pb.ProcessPaymentResponse, error)

	// DoctorNoteService Methods
	// CreateDoctorNote(ctx context.Context, req *pb.CreateDoctorNoteRequest) (*pb.DoctorNote, error)
	// GetDoctorNote(ctx context.Context, req *pb.GetDoctorNoteRequest) (*pb.DoctorNote, error)
	// UpdateDoctorNote(ctx context.Context, req *pb.UpdateDoctorNoteRequest) (*pb.DoctorNote, error)
	// DeleteDoctorNote(ctx context.Context, req *pb.DeleteDoctorNoteRequest) (*pb.DoctorNote, error)

	// FeedbackService Methods
	//SubmitFeedback(ctx context.Context, req *pb.SubmitFeedbackRequest) (*pb.SubmitFeedbackResponse, error)

	// NotificationService Methods
	//SendNotification(ctx context.Context, req *pb.SendNotificationRequest) (*pb.SendNotificationResponse, error)


	// ReportingService Methods
	//GenerateReport(ctx context.Context, req *pb.GenerateReportRequest) (*pb.GenerateReportResponse, error)

	// SearchService Methods
	// SearchDoctors(ctx context.Context, req *pb.SearchDoctorsRequest) (*pb.SearchDoctorsResponse, error)
	// SearchPatients(ctx context.Context, req *pb.SearchPatientsRequest) (*pb.SearchPatientsResponse, error)
	// SearchAppointments(ctx context.Context, req *pb.SearchAppointmentsRequest) (*pb.SearchAppointmentsResponse, error)

	// UploadedFileService Methods
	// UploadFile(ctx context.Context, req *pb.UploadFileRequest) (*pb.UploadedFile, error)
	// GetFile(ctx context.Context, req *pb.GetFileRequest) (*pb.UploadedFile, error)
	// UpdateFile(ctx context.Context, req *pb.UpdateFileRequest) (*pb.UploadedFile, error)
	// DeleteFile(ctx context.Context, req *pb.DeleteFileRequest) (*pb.UploadedFile, error)
}
