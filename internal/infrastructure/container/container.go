package container

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	helper "rakamin-final/internal/helper"
	"rakamin-final/internal/infrastructure/mysql"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

const currentfilepath = "internal/infrastructure/container/container.go"

type (
	Container struct {
		MySqlDB *gorm.DB
		Apps    *Apps
	}
	Apps struct {
		Name     string `mapstructure:"name"`
		Host     string `mapstructure:"host"`
		Version  string `mapstructure:"version"`
		Address  string `mapstructure:"address"`
		HttpPort string `mapstructure:"httpport"`

		JwtSecret string `mapstructure:"jwt_secret"`
	}
)

var v *viper.Viper

func loadEnv() {
	projectDirName := "rakamin-final"
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	v.SetConfigFile(string(rootPath) + `/.env`)
}

func AppsInit(v *viper.Viper) (apps Apps) {
	err := v.Unmarshal(&apps)
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprint("Error when unmarshal configuration file : ", err.Error()))
	}
	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed when unmarshal configuration file")
	return
}
func InitContainer() (cont *Container) {
	apps := AppsInit(v)
	mysqldb := mysql.InitDB(v)

	return &Container{
		MySqlDB: mysqldb,
		Apps:    &apps,
	}
}

func init() {
	v = viper.New()

	v.AutomaticEnv()
	loadEnv()

	path, err := os.Executable()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("os.Executable panic : %s", err.Error()))
	}

	dir := filepath.Dir(path)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed read config : %s", err.Error()))
	}

	err = v.ReadInConfig()
	if err != nil {
		helper.Logger(currentfilepath, helper.LoggerLevelPanic, fmt.Sprintf("failed init config : %s", err.Error()))
	}

	helper.Logger(currentfilepath, helper.LoggerLevelInfo, "Succeed read configuration file")
}
