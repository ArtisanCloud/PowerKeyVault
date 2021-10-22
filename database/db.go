package database

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/golang-module/carbon"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

var (
	DBConnection *gorm.DB
)

func SetupDatabase() (err error) {

	d := config.DatabaseConn
	c := d.Connection
	timezone := config.AppConfigure.Timezone
	if timezone == "" {
		timezone = carbon.UTC
	}
	//dsn := "host=" + c.Host
	//dsn += " user=" + c.Username
	//dsn += " password=" + c.Password
	//dsn += " dbname=" + c.Database
	//dsn += " port=" + c.Port
	//dsn += " sslmode=" + d.SSLMode
	//dsn += " TimeZone=" + timezone
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset)
	fmt.Println("dsn =>", dsn)

	DBConnection, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         255,
		DefaultDatetimePrecision:  nil,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		DontSupportForShareClause: false,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//DBConnection.Exec("SET search_path TO " + d.SearchPath)

	if err != nil {
		// throw a exception here
		log.Fatal("Database init error: ", err)
		return
	}

	DbSql, _ := DBConnection.DB()
	DbSql.SetConnMaxLifetime(time.Hour)
	DbSql.SetConnMaxIdleTime(time.Hour)
	DbSql.SetMaxIdleConns(20)
	DbSql.SetMaxOpenConns(100)
	fmt.Println("Mysql Connect Success")

	return err

}

//func NewContext()  context.Context {
//	return context.Context{}
//}
func GetDBWithContext(ctx context.Context) *gorm.DB {
	//var newCTX context.Context
	return DBConnection.WithContext(ctx)
}
