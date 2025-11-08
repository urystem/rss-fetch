package configs

import "math"

type dbConfig struct {
	host     string
	port     uint16
	user     string
	password string
	name     string
	sslMode  string
}
type DBConfig interface {
	GetHostName() string
	GetPort() uint16
	GetUser() string
	GetPassword() string
	GetDBName() string
	GetSSLMode() string
}

func initDBConfig() dbConfig {
	dbConf := dbConfig{}
	dbConf.host = mustGetEnvString("POSTGRES_HOST")
	temPort := mustGetEnvInt("POSTGRES_PORT")
	if temPort < 0 || math.MaxUint16 < temPort {
		panic("invalid port psql")
	}
	dbConf.port = uint16(temPort)
	dbConf.user = mustGetEnvString("POSTGRES_USER")
	dbConf.password = mustGetEnvString("POSTGRES_PASSWORD")
	dbConf.name = mustGetEnvString("POSTGRES_DBNAME")
	dbConf.sslMode = mustGetEnvString("POSTGRES_SSLMODE")
	return dbConf
}

func (dbC *dbConfig) GetHostName() string { return dbC.host }

func (dbC *dbConfig) GetPort() uint16 { return dbC.port }

func (dbC *dbConfig) GetUser() string { return dbC.user }

func (dbC *dbConfig) GetPassword() string { return dbC.password }

func (dbC *dbConfig) GetDBName() string { return dbC.name }

func (dbC *dbConfig) GetSSLMode() string { return dbC.sslMode }
