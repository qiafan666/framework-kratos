package biz

import (
	"context"
	common "framework-kratos/api/common/v1/gen"
	"framework-kratos/app/test/service/v1/internal/data"
	"framework-kratos/app/test/service/v1/model"
	"github.com/go-kratos/kratos/v2/log"
)

type TestCase struct {
	log      *log.Helper
	TestRepo *data.TestRepo
}

func NewTestCase(testRepo *data.TestRepo, logger log.Logger) *TestCase {
	return &TestCase{
		log:      log.NewHelper(log.With(logger, "module", "biz/test")),
		TestRepo: testRepo,
	}
}

func (t *TestCase) CreateTest(ctx context.Context, test *model.Test) error {
	return t.TestRepo.CreateTest(ctx, test)
}

func (t *TestCase) UpdateTest(ctx context.Context, test *model.Test) error {
	return t.TestRepo.UpdateTest(ctx, test)
}

func (t *TestCase) DeleteTest(ctx context.Context, id int64) error {
	return t.TestRepo.DeleteTest(ctx, id)
}

func (t *TestCase) GetTest(ctx context.Context, id int64) (model.Test, error) {
	test, err := t.TestRepo.FirstTest(ctx, id)
	if err != nil {
		return model.Test{}, err
	}
	return test, nil
}

func (t *TestCase) ListTest(ctx context.Context, page *common.Page, test *model.Test) (model.Tests, int64, error) {
	listTest, i, err := t.TestRepo.ListTest(ctx, page, test)
	if err != nil {
		return nil, 0, err
	}
	return listTest, i, nil
}
