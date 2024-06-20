package service

import (
	"context"
	"framework-kratos/app/test/service/v1/internal/biz"
	"framework-kratos/app/test/service/v1/model"
	"github.com/go-kratos/kratos/v2/log"

	pb "framework-kratos/api/test/service/v1"
)

type TestService struct {
	pb.UnimplementedTestServer
	log      *log.Helper
	testCase *biz.TestCase
}

func NewTestService(tc *biz.TestCase, logger log.Logger) *TestService {
	return &TestService{
		log:      log.NewHelper(log.With(logger, "module", "service/test")),
		testCase: tc,
	}
}

func (s *TestService) CreateTest(ctx context.Context, req *pb.CreateTestRequest) (*pb.CreateTestReply, error) {
	err := s.testCase.CreateTest(ctx, &model.Test{})
	if err != nil {
		return nil, err
	}
	return &pb.CreateTestReply{}, nil
}
func (s *TestService) UpdateTest(ctx context.Context, req *pb.UpdateTestRequest) (*pb.UpdateTestReply, error) {
	err := s.testCase.UpdateTest(ctx, &model.Test{})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateTestReply{}, nil
}
func (s *TestService) DeleteTest(ctx context.Context, req *pb.DeleteTestRequest) (*pb.DeleteTestReply, error) {
	err := s.testCase.DeleteTest(ctx, 1)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTestReply{}, nil
}
func (s *TestService) GetTest(ctx context.Context, req *pb.GetTestRequest) (*pb.GetTestReply, error) {
	_, err := s.testCase.GetTest(ctx, 1)
	if err != nil {
		return nil, err
	}

	return &pb.GetTestReply{}, nil
}
func (s *TestService) ListTest(ctx context.Context, req *pb.ListTestRequest) (*pb.ListTestReply, error) {
	return &pb.ListTestReply{}, nil
}
