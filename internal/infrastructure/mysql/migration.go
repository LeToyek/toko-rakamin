package mysql

import (
	"rakamin-final/internal/daos"
	"rakamin-final/internal/helper"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&daos.LogProduk{},
		&daos.Trx{},
		&daos.DetailTrx{},
		&daos.Category{},
		&daos.Toko{},
		&daos.Produk{},
		&daos.FotoProduk{},
		&daos.User{},
		&daos.Alamat{},
	)

	var count int64
	if mysqlDB.Migrator().HasTable(&daos.User{}) {
		mysqlDB.Model(&daos.User{}).Count(&count)
		if count < 1 {
			helper.Logger("internal/infrastructure/mysql/migration.go", helper.LoggerLevelInfo, "No rows in table users, seeding data...")
		}
		if err != nil {
			helper.Logger("internal/infrastructure/mysql/migration.go", helper.LoggerLevelPanic, err.Error())
		}

		helper.Logger("internal/infrastructure/mysql/migration.go", helper.LoggerLevelInfo, "Migration success")
	}
}
