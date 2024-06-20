package data

import (
	"context"
	common "framework-kratos/api/common/v1"
	"framework-kratos/app/test/service/v1/model"
	"framework-kratos/pkg/function"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type TestRepo struct {
	data Dao
	log  *log.Helper
}

func (r TestRepo) CreateTest(ctx context.Context, test *model.Test) error {
	return r.data.WithContext(ctx).Create(test)
}

func (r TestRepo) UpdateTest(ctx context.Context, test *model.Test) error {
	_, err := r.data.WithContext(ctx).Update(model.Test{}, map[string]interface{}{
		model.TestColumns.ID: test.ID,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r TestRepo) DeleteTest(ctx context.Context, id int64) error {
	_, err := r.data.WithContext(ctx).Update(model.Test{
		IsDeleted: 1,
	}, map[string]interface{}{
		model.TestColumns.ID: id,
	}, nil)
	if err != nil {
		return err
	}
	return nil
}

func (r TestRepo) FirstTest(ctx context.Context, id int64) (model.Test, error) {
	var test model.Test
	err := r.data.WithContext(ctx).First([]string{"*"}, map[string]interface{}{
		model.TestColumns.ID: id,
	}, nil, &test)
	if err != nil {
		return model.Test{}, err
	}
	return test, nil
}

func (r TestRepo) ListTest(ctx context.Context, page *common.Page, test *model.Test) (model.Tests, int64, error) {
	var tests model.Tests
	err := r.data.WithContext(ctx).Find([]string{"*"}, map[string]interface{}{
		model.TestColumns.IsDeleted: 0,
	}, func(db *gorm.DB) *gorm.DB {
		db = db.Scopes(function.Paginate(page.Current, page.Size))
		return db
	}, &tests)
	if err != nil {
		return nil, 0, err
	}

	total, err := r.data.WithContext(ctx).Count(model.Test{}, map[string]interface{}{
		model.TestColumns.IsDeleted: 0,
	}, nil)
	if err != nil {
		return nil, 0, err
	}

	return tests, total, nil
}

var daoInstance Dao

func NewTestRepo(data *Data, logger log.Logger) *TestRepo {
	once.Do(func() {
		daoInstance = Instance(data.db)
	})

	return &TestRepo{
		data: daoInstance,
		log:  log.NewHelper(log.With(logger, "module", "data/test")),
	}
}
