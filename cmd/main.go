package main

import (
	"Booking-service/config"
	pb "Booking-service/genproto/booking-service"
	"Booking-service/pkg/db"
	"Booking-service/pkg/logger"
	"Booking-service/service/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "USER_SERVICE")
	defer logger.Cleanup(log)

	log.Info("main: sqlConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))

	connDB, err, _ := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	userService := services.NewBookingService(connDB, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterPatientServiceServer(s, userService)
	pb.RegisterDoctorAvailabilityServiceServer(s, userService)
	pb.RegisterBookedAppointmentServiceServer(s, userService)
	pb.RegisterArchiveServiceServer(s, userService)
	pb.RegisterDoctorNoteServiceServer(s, userService)
	pb.RegisterPatientPaymentServiceServer(s,userService)
	pb.RegisterUploadedFileServiceServer(s,userService)
	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))

	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
