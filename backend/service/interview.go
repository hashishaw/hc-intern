package service

import (
	"context"
	pb "github.com/hashishaw/hc-intern/backend/backend/gen"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InterviewService struct {
	// Embed the unimplemented server
	// See https://github.com/grpc/grpc-go/issues/3794
	pb.UnimplementedInterviewServiceServer
}

func (s *InterviewService) Get (ctx context.Context, req *pb.GetRequest) (*pb.Interview, error) {
	if req.InterviewId == "foobar" {
		return &pb.Interview{
			Exists: true,
			CandidateName: "Chelsea Shaw",
		}, nil
	}
	return nil, status.Error(codes.NotFound, "no interview found for id")
}

func (s *InterviewService) Schedule (ctx context.Context, req *pb.ScheduleRequest) (*pb.ScheduleResponse, error) {
	if req.CandidateId != 0 {
		return &pb.ScheduleResponse{
			InterviewId: "foobar",
		}, nil
	}
	return nil, status.Error(codes.InvalidArgument, "candidate id not valid")
}