package boostrap

import (
	service "github.com/ArtisanCloud/PowerKeyVault/app/services"
	"github.com/ArtisanCloud/PowerKeyVault/cache"
	"github.com/ArtisanCloud/PowerKeyVault/config"
	"github.com/ArtisanCloud/PowerKeyVault/database"
	logger "github.com/ArtisanCloud/PowerKeyVault/loggerManager"
	"github.com/ArtisanCloud/PowerKeyVault/resources/lang"
)

func InitProject() {
	// Initialize the global config
	envConfigName := "environment"
	dbConfigName := "database"
	cacheConfigName := "cache"
	logConfigName := "log"
	config.LoadEnvConfig(nil, &envConfigName, nil)
	config.LoadDatabaseConfig(nil, &dbConfigName, nil)
	config.LoadCacheConfig(nil, &cacheConfigName, nil)
	config.LoadVersion()
	config.LoadLogConfig(nil, &logConfigName, nil)

	// load locale
	lang.LoadLanguages()

	// setup ssh key path
	service.SetupSSHKeyPath(config.AppConfigure.SSH)

	// Initialize the cache
	_ = cache.SetupCache()

	// Initialize the database
	_ = database.SetupDatabase()

	// Initialize the logger
	_ = logger.SetupLog()
}