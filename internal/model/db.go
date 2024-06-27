package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xwxb/simple-fileshare/internal/config"
	"log"
	"xorm.io/xorm"
)

var dbEngine *xorm.Engine

func InitDB() {
	var err error
	dbEngine, err = NewMySQLEngine(GetMySQLDSNByConfig(&config.GlobalConfig.MySQL))
	if err != nil {
		log.Fatal(err)
	}

	err = dbEngine.Sync(new(User))
}

func DeferDBClose() {
	err := dbEngine.Close()
	if err != nil {
		log.Println(err)
	}
}

func GetMySQLDSNByConfig(cfg *config.MySQLConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
}

func NewMySQLEngine(connectionString string) (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("mysql", connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	err = engine.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	return engine, nil
}
