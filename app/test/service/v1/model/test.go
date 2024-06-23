package model

import (
	v1 "framework-kratos/api/test/service/v1/gen"
	"github.com/jinzhu/copier"
	"time"
)

/******sql******
CREATE TABLE `test` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  `is_deleted` tinyint DEFAULT '0' COMMENT '是否删除 0-未删除 1-已删除',
  `name` varchar(255) DEFAULT NULL COMMENT '名称',
  `age` int DEFAULT NULL COMMENT '年龄',
  `desc` text COMMENT '描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
******sql******/
// Test [...]
type Test struct {
	ID          int64     `gorm:"primaryKey;column:id" json:"-"` // 主键ID
	CreatedTime time.Time `gorm:"column:created_time" json:"created_time"`
	UpdatedTime time.Time `gorm:"column:updated_time" json:"updated_time"`
	IsDeleted   int8      `gorm:"column:is_deleted" json:"is_deleted"` // 是否删除 0-未删除 1-已删除
	Name        string    `gorm:"column:name" json:"name"`             // 名称
	Age         int       `gorm:"column:age" json:"age"`               // 年龄
	Desc        string    `gorm:"column:desc" json:"desc"`             // 描述
}

// TableName get sql table name.获取数据库表名
func (m *Test) TableName() string {
	return "test"
}

// TestColumns get sql column name.获取数据库列名
var TestColumns = struct {
	ID          string
	CreatedTime string
	UpdatedTime string
	IsDeleted   string
	Name        string
	Age         string
	Desc        string
}{
	ID:          "id",
	CreatedTime: "created_time",
	UpdatedTime: "updated_time",
	IsDeleted:   "is_deleted",
	Name:        "name",
	Age:         "age",
	Desc:        "desc",
}

// ToTest 转成proto
func (a *Test) ToTest() *v1.Test {
	var result v1.Test
	_ = copier.CopyWithOption(&result, a, copier.Option{
		DeepCopy:    true,
		IgnoreEmpty: true,
	})
	return &result
}

// Tests .
type Tests []*Test

// ToTests .
func (a Tests) ToTests() []*v1.Test {
	d := make([]*v1.Test, 0, len(a))
	for i := range a {
		d = append(d, a[i].ToTest())
	}
	return d
}
