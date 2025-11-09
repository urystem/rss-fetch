package configs

type configs struct {
	db     dbConfig
	worker worker
}
type ConfigInter interface {
	GetDBConfig() DBConfig
	GetWorkerCfg() WorkerInter
}

func Load() ConfigInter {
	return &configs{
		db:     initDBConfig(),
		worker: initWorkerCfg(),
	}
}

func (conf *configs) GetDBConfig() DBConfig { return &conf.db }

func (conf *configs) GetWorkerCfg() WorkerInter { return &conf.worker }
