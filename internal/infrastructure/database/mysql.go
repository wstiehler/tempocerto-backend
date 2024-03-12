package config

import (
	"fmt"

	"github.com/wstiehler/tempocerto-backend/internal/environment"
	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger"
	"github.com/wstiehler/tempocerto-backend/internal/infrastructure/logger/logwrapper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var zapLogger *zap.Logger

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func buildDBConfig() *DBConfig {
	env := environment.GetInstance()

	dbConfig := DBConfig{
		Host:     env.MYSQL_HOST,
		Port:     int(env.MYSQL_PORT),
		User:     env.MYSQL_USER,
		Password: env.MYSQL_PASSWORD,
		DBName:   env.MYSQL_DBNAME,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)

	return dbURL
}

func InitLogger() {
	var dispose func()
	zapLogger, dispose = logger.New()
	defer dispose()
}

func InitDB() error {
	InitLogger()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zapLogger})

	dbConfig := buildDBConfig()
	logger.Info("Host Database", zap.String("Host", dbConfig.Host))

	if err := ConnectDb(); err != nil {
		logger.Error("Failed to initialize database", zap.Error(err))
		return err
	}

	return nil
}

func ConnectDb() error {
	zapLogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zapLogger})

	config, err := gorm.Open(mysql.Open(DbURL(buildDBConfig())), &gorm.Config{})
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return err
	}

	logger.Info("Successfully connected to database")
	DB = config

	return nil
}

func CloseConnection(mySqlConfig *gorm.DB) {
	zapLogger, dispose := logger.New()
	defer dispose()

	logger := logwrapper.New(&logwrapper.Zap{Logger: *zapLogger})

	logger.Info("Closing database connection")
	sqlDB, err := mySqlConfig.DB()
	if err != nil {
		logger.Error("Failed to get database connection", zap.Error(err))
		return
	}
	err = sqlDB.Close()
	if err != nil {
		logger.Error("Failed to close database connection", zap.Error(err))
	}
}
