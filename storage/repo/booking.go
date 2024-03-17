package service

import (
	pb "Booking-service/genproto/booking-service"
)

type BookingStorageI interface {
	// PatientService Methods
	CreatePatient(*pb.Patient) (*pb.Patient, error)
	GetPatient(*pb.GetPatientRequest) (*pb.Patient, error)
	GetPatients(*pb.GetPatientsRequest) (*pb.Patients, error)
	UpdatePatient(*pb.UpdatePatientRequest) (*pb.Patient, error)
	DeletePatient(*pb.GetPatientRequest) (bool, error)

	// DoctorAvailabilityService Methods
	CreateDoctorAvailability(*pb.CreateDoctorAvailabilitys) (*pb.DoctorAvailability, error)
	GetDoctorAvailability(*pb.GetDoctorAvailabilityById) (*pb.DoctorAvailability, error)
	GetDoctorAvailabilityByDoctorId(*pb.GetDoctorAvailabilityById) (*pb.DoctorAvailabilitys, error)
	UpdateDoctorAvailability(*pb.UpdateDoctorAvailabilityById) (*pb.DoctorAvailability, error)
	DeleteDoctorAvailability(*pb.GetDoctorAvailabilityById) (bool, error)

	// BookedAppointmentService Methods
	CreateBookedAppointment(*pb.CreateBookedAppointments) (*pb.BookedAppointment, error)
	GetBookedAppointment(*pb.GetRequest) (*pb.BookedAppointment, error)
	GetBookedAppointmentsByPatientID(*pb.GetRequest) (*pb.GetBookedAppointments, error)
	GetBookedAppointmentsByDoctorID(*pb.GetRequest) (*pb.GetBookedAppointments, error)
	UpdateBookedAppointment(*pb.UpdateBookedAppointmentRequest) (*pb.BookedAppointment, error)
	UpdatePatientStatusByToken(*pb.UpdRequest) (*pb.BookedAppointment, error)
	DeleteBookedAppointment(*pb.GetRequest) (bool, error)

	// PatientPaymentService Methods
	MakePayment(*pb.PatientPayment) (*pb.PatientPayment, error)
	GetPayment(*pb.GetPaymentReq) (*pb.PatientPayment, error)
	GetPaymentsByPatienId(*pb.GetPaymentReq) (*pb.GetPaymentsResp, error)
	UpdatePayment(*pb.UpdatePaymentRequest) (*pb.PatientPayment, error)
	DeletePayment(*pb.GetPaymentReq) (bool, error)

	// ArchiveService Methods
	CreateArchive(*pb.CreateArchiveReq) (*pb.Archive, error)
	GetArchive(*pb.GetArchiveReq) (*pb.Archive, error)
	GetArchivesByPatientID(*pb.GetArchiveReq) (*pb.Archives, error)
	UpdateArchive(*pb.UpdateArchiveRequest) (*pb.Archive, error)
	DeleteArchive(*pb.GetArchiveReq) (bool, error)

	// UploadedFileService Methods
	UploadFile(*pb.CreateUploadedFile) (*pb.UploadedFile, error)
	GetFileByPatientID(*pb.GetFileRequest) (*pb.UploadedFile, error)
	GetFilesByPatientID(*pb.GetFileRequest) (*pb.UploadedFiles, error)
	UpdateFile(*pb.UpdateFileRequest) (*pb.UploadedFile, error)
	DeleteFile(*pb.GetFileRequest) (bool, error)

	// DoctorNoteService Methods
	CreateDoctorNote(*pb.CreateDoctorNoteReq) (*pb.DoctorNote, error)
	GetDoctorNote(*pb.GetDoctorNoteReq) (*pb.DoctorNote, error)
	GetDoctorNotesByPatienId(*pb.GetDoctorNoteReq) (*pb.DoctorNotes, error)
	UpdateDoctorNote(*pb.UpdateDoctorNoteReq) (*pb.DoctorNote, error)
	DeleteDoctorNote(*pb.GetDoctorNoteReq) (bool, error)
}
