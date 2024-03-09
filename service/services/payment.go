package services

import (
	"context"
	"errors"

	pb "Booking-service/genproto/booking-service"
)

func (s *BookingService) MakePayment(ctx context.Context, req *pb.PatientPayment) (*pb.PatientPayment, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().MakePayment(req)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) GetPayment(ctx context.Context, req *pb.GetPaymentReq) (*pb.PatientPayment, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().GetPayment(req)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) GetPaymentsByPatienId(ctx context.Context, req *pb.GetPaymentReq) (*pb.GetPaymentsResp, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	payments, err := s.storage.Booking().GetPaymentsByPatienId(req)
	if err != nil {
		return nil, err
	}

	return payments, nil
}

func (s *BookingService) UpdatePayment(ctx context.Context, req *pb.UpdatePaymentRequest) (*pb.PatientPayment, error) {
	if req == nil || req.Payment == nil {
		return nil, errors.New("invalid request")
	}

	payment, err := s.storage.Booking().UpdatePayment(req)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (s *BookingService) DeletePayment(ctx context.Context, req *pb.GetPaymentReq) (*pb.Status, error) {
	if req == nil {
		return nil, errors.New("invalid request")
	}

	status, err := s.storage.Booking().DeletePayment(req)
	if err != nil {
		return nil, err
	}

	return &pb.Status{Status: status}, nil
}
