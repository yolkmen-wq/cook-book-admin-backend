package respositories

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var RedisClient *redis.Client

func init() {
	// 建立数据库连接
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:wuqi9457@tcp(47.121.201.137:3306)/cook_book?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                      // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                     // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                     // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                     // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                    // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("初始化数据库成功！")

	// 建立Redis连接
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "47.121.201.137:6011",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	fmt.Println("连接Redis中...", RedisClient)
	if RedisClient != nil {
		fmt.Println("连接Redis成功！")
	}
}
