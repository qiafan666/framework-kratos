package service

import (
	"context"
	"framework-kratos/api/test/service/v1/gen"
	"framework-kratos/app/test/service/v1/internal/biz"
	"framework-kratos/app/test/service/v1/model"
	"github.com/go-kratos/kratos/v2/log"
)

type TestService struct {
	gen.UnimplementedTestServer
	log      *log.Helper
	testCase *biz.TestCase
}

func NewTestService(tc *biz.TestCase, logger log.Logger) *TestService {
	return &TestService{
		log:      log.NewHelper(log.With(logger, "module", "service/test")),
		testCase: tc,
	}
}

func (s *TestService) CreateTest(ctx context.Context, req *gen.CreateTestRequest) (*gen.CreateTestReply, error) {
	err := s.testCase.CreateTest(ctx, &model.Test{})
	if err != nil {
		return nil, err
	}
	return &gen.CreateTestReply{}, nil
}
func (s *TestService) UpdateTest(ctx context.Context, req *gen.UpdateTestRequest) (*gen.UpdateTestReply, error) {
	err := s.testCase.UpdateTest(ctx, &model.Test{})
	if err != nil {
		return nil, err
	}
	return &gen.UpdateTestReply{}, nil
}
func (s *TestService) DeleteTest(ctx context.Context, req *gen.DeleteTestRequest) (*gen.DeleteTestReply, error) {
	err := s.testCase.DeleteTest(ctx, 1)
	if err != nil {
		return nil, err
	}
	return &gen.DeleteTestReply{}, nil
}
func (s *TestService) GetTest(ctx context.Context, req *gen.GetTestRequest) (*gen.GetTestReply, error) {
	_, err := s.testCase.GetTest(ctx, 1)
	if err != nil {
		return nil, err
	}

	return &gen.GetTestReply{}, nil
}
func (s *TestService) ListTest(ctx context.Context, req *gen.ListTestRequest) (*gen.ListTestReply, error) {
	return &gen.ListTestReply{}, nil
}
