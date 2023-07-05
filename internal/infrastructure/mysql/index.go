package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/spf13/viper"
)

type MysqlConf struct {
	Username           string `mapstructure:"mysql_username"`
	Password           string `mapstructure:"mysql_password"`
	DbName             string `mapstructure:"mysql_Dbname"`
	Host               string `mapstructure:"mysql_host"`
	Port               int    `mapstructure:"mysql_port"`
	Schema             string `mapstructure:"mysql_schema"`
	LogMode            bool   `mapstructure:"mysql_logMode"`
	MaxLifetime        int    `mapstructure:"mysql_maxLifetime"`
	MinIdleConnections int    `mapstructure:"mysql_minIdleConnections"`
	MaxOpenConnections int    `mapstructure:"mysql_maxOpenConnections"`
}

func InitDB(v *viper.Viper) *gorm.DB {
	mysqlConf := MysqlConf{}
	err := v.Unmarshal(&mysqlConf)
	fmt.Print(mysqlConf)
	if err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlConf.Username,
		mysqlConf.Password,
		mysqlConf.Host,
		mysqlConf.Port,
		mysqlConf.DbName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db

}
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}

	dbSQL.Close()
}
