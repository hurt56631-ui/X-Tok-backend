package config

import (
	"fmt"
	"log"
	"os"
	"strconv" // 导入strconv包用于字符串转数字
	"strings" // 导入strings包用于处理字符串列表

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ServerConfig 服务器配置
type ServerConfig struct {
	AppDebug         bool `mapstructure:"AppDebug"`
	HttpServer       HttpServerConfig `mapstructure:"HttpServer"`
	Token            TokenConfig      `mapstructure:"Token"`
	LoginPolicy      LoginPolicyConfig `mapstructure:"LoginPolicy"`
	Redis            RedisConfig      `mapstructure:"Redis"`
	Logs             LogsConfig       `mapstructure:"Logs"`
	Websocket        WebsocketConfig  `mapstructure:"Websocket"`
	SnowFlake        SnowFlakeConfig  `mapstructure:"SnowFlake"`
	FileUploadSetting FileUploadSettingConfig `mapstructure:"FileUploadSetting"`
	RabbitMq         RabbitMqConfig   `mapstructure:"RabbitMq"`
	Casbin           CasbinConfig     `mapstructure:"Casbin"`
	Captcha          CaptchaConfig    `mapstructure:"Captcha"`
}

type HttpServerConfig struct {
	Web            WebConfig `mapstructure:"Web"`
	AllowCrossDomain bool `mapstructure:"AllowCrossDomain"`
	TrustProxies   TrustProxiesConfig `mapstructure:"TrustProxies"`
}

type WebConfig struct {
	Port string `mapstructure:"Port"`
}

type TrustProxiesConfig struct {
	IsOpen          int      `mapstructure:"IsOpen"`
	ProxyServerList []string `mapstructure:"ProxyServerList"`
}

type TokenConfig struct {
	JwtTokenSignKey       string `mapstructure:"JwtTokenSignKey"`
	JwtDefaultUid         int    `mapstructure:"JwtDefaultUid"`
	JwtTokenOnlineUsers   int    `mapstructure:"JwtTokenOnlineUsers"`
	JwtTokenCreatedExpireAt int `mapstructure:"JwtTokenCreatedExpireAt"`
	JwtTokenRefreshAllowSec int `mapstructure:"JwtTokenRefreshAllowSec"`
	JwtTokenRefreshExpireAt int `mapstructure:"JwtTokenRefreshExpireAt"`
	BindContextKeyName    string `mapstructure:"BindContextKeyName"`
	IsCacheToRedis        int    `mapstructure:"IsCacheToRedis"`
}

type LoginPolicyConfig struct {
	IsOpenPolicy      int `mapstructure:"IsOpenPolicy"`
	MaxLoginFailTimes int `mapstructure:"MaxLoginFailTimes"`
	LoginFailCountDown int `mapstructure:"LoginFailCountDown"`
}

type RedisConfig struct {
	Host               string `mapstructure:"Host"`
	Port               int    `mapstructure:"Port"`
	Auth               string `mapstructure:"Auth"`
	MaxIdle            int    `mapstructure:"MaxIdle"`
	MaxActive          int    `mapstructure:"MaxActive"`
	IdleTimeout        int    `mapstructure:"IdleTimeout"`
	IndexDb            int    `mapstructure:"IndexDb"`
	ConnFailRetryTimes int    `mapstructure:"ConnFailRetryTimes"`
	ReConnectInterval  int    `mapstructure:"ReConnectInterval"`
}

type LogsConfig struct {
	GinLogName        string `mapstructure:"GinLogName"`
	GoSkeletonLogName string `mapstructure:"GoSkeletonLogName"`
	TextFormat        string `mapstructure:"TextFormat"`
	TimePrecision     string `mapstructure:"TimePrecision"`
	MaxSize           int    `mapstructure:"MaxSize"`
	MaxBackups        int    `mapstructure:"MaxBackups"`
	MaxAge            int    `mapstructure:"MaxAge"`
	Compress          bool   `mapstructure:"Compress"`
}

type WebsocketConfig struct {
	Start              int `mapstructure:"Start"`
	WriteReadBufferSize int `mapstructure:"WriteReadBufferSize"`
	MaxMessageSize     int `mapstructure:"MaxMessageSize"`
	PingPeriod         int `mapstructure:"PingPeriod"`
	HeartbeatFailMaxTimes int `mapstructure:"HeartbeatFailMaxTimes"`
	ReadDeadline       int `mapstructure:"ReadDeadline"`
	WriteDeadline      int `mapstructure:"WriteDeadline"`
}

type SnowFlakeConfig struct {
	SnowFlakeMachineId int `mapstructure:"SnowFlakeMachineId"`
}

type FileUploadSettingConfig struct {
	Size                      int      `mapstructure:"Size"`
	UploadFileField           string   `mapstructure:"UploadFileField"`
	SourceUrlPrefix           string   `mapstructure:"SourceUrlPrefix"`
	UploadRootPath            string   `mapstructure:"UploadRootPath"`
	AvatarSmallUploadFileSavePath string `mapstructure:"AvatarSmallUploadFileSavePath"`
	AvatarLargeUploadFileSavePath string `mapstructure:"AvatarLargeUploadFileSavePath"`
	CoverUploadFileSavePath   string `mapstructure:"CoverUploadFileSavePath"`
	WhiteCoverUploadFileSavePath string `mapstructure:"WhiteCoverUploadFileSavePath"`
	VideoUploadFileSavePath   string `mapstructure:"VideoUploadFileSavePath"`
	VideoCoverUploadFileSavePath string `mapstructure:"VideoCoverUploadFileSavePath"`
	AllowMimeType             []string `mapstructure:"AllowMimeType"`
}

type RabbitMqConfig struct {
	HelloWorld     RabbitMqQueueConfig `mapstructure:"HelloWorld"`
	WorkQueue      RabbitMqQueueConfig `mapstructure:"WorkQueue"`
	PublishSubscribe RabbitMqExchangeConfig `mapstructure:"PublishSubscribe"`
	Routing        RabbitMqExchangeConfig `mapstructure:"Routing"`
	Topics         RabbitMqExchangeConfig `mapstructure:"Topics"`
}

type RabbitMqQueueConfig struct {
	Addr                string `mapstructure:"Addr"`
	QueueName           string `mapstructure:"QueueName"`
	Durable             bool   `mapstructure:"Durable"`
	ConsumerChanNumber  int    `mapstructure:"ConsumerChanNumber"`
	OffLineReconnectIntervalSec int `mapstructure:"OffLineReconnectIntervalSec"`
	RetryCount          int    `mapstructure:"RetryCount"`
}

type RabbitMqExchangeConfig struct {
	Addr                string `mapstructure:"Addr"`
	ExchangeType        string `mapstructure:"ExchangeType"`
	ExchangeName        string `mapstructure:"ExchangeName"`
	DelayedExchangeName string `mapstructure:"DelayedExchangeName"`
	Durable             bool   `mapstructure:"Durable"`
	QueueName           string `mapstructure:"QueueName"`
	ConsumerChanNumber  int    `mapstructure:"ConsumerChanNumber"`
	OffLineReconnectIntervalSec int `mapstructure:"OffLineReconnectIntervalSec"`
	RetryCount          int    `mapstructure:"RetryCount"`
}

type CasbinConfig struct {
	IsInit            int    `mapstructure:"IsInit"`
	AutoLoadPolicySeconds int `mapstructure:"AutoLoadPolicySeconds"`
	TablePrefix       string `mapstructure:"TablePrefix"`
	TableName         string `mapstructure:"TableName"`
	ModelConfig       string `mapstructure:"ModelConfig"`
}

type CaptchaConfig struct {
	CaptchaId    string `mapstructure:"captchaId"`
	CaptchaValue string `mapstructure:"captchaValue"`
	Length       int    `mapstructure:"length"`
}


// GormV2Config 数据库配置，注意这里包含了不同的数据库类型
type GormV2Config struct {
	UseDbType         string `mapstructure:"UseDbType"`
	Mysql             DbConfig `mapstructure:"Mysql"`
	SqlServer         DbConfig `mapstructure:"SqlServer"`
	PostgreSql        DbConfig `mapstructure:"PostgreSql"`
}

type DbConfig struct {
	IsInitGlobalGormMysql    int `mapstructure:"IsInitGlobalGormMysql"` // 这里的字段名需要和yaml中的完全匹配
	IsInitGlobalGormSqlserver int `mapstructure:"IsInitGlobalGormSqlserver"` // for sqlserver
	IsInitGlobalGormPostgreSql int `mapstructure:"IsInitGlobalGormPostgreSql"` // for postgresql
	SlowThreshold        int    `mapstructure:"SlowThreshold"`
	Write                DbConnConfig `mapstructure:"Write"`
	IsOpenReadDb         int    `mapstructure:"IsOpenReadDb"`
	Read                 DbConnConfig `mapstructure:"Read"`
}

type DbConnConfig struct {
	Host             string `mapstructure:"Host"`
	DataBase         string `mapstructure:"DataBase"`
	Port             int    `mapstructure:"Port"`
	Prefix           string `mapstructure:"Prefix"`
	User             string `mapstructure:"User"`
	Pass             string `mapstructure:"Pass"`
	Charset          string `mapstructure:"Charset"`
	SetMaxIdleConns  int    `mapstructure:"SetMaxIdleConns"`
	SetMaxOpenConns  int    `mapstructure:"SetMaxOpenConns"`
	SetConnMaxLifetime int `mapstructure:"SetConnMaxLifetime"`
}


// GlobalServerConfig 全局服务器配置实例
var GlobalServerConfig = &ServerConfig{}

// GlobalGormV2Config 全局Gorm V2配置实例
var GlobalGormV2Config = &GormV2Config{}

// Init 初始化配置
func Init() {
	// 读取 config.yaml
	viper.SetConfigFile("config.yaml")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Warning: config.yaml not found. Relying on environment variables. Error: %v", err)
		} else {
			log.Fatalf("Fatal error reading config.yaml: %v", err)
		}
	} else {
		if err := viper.Unmarshal(GlobalServerConfig); err != nil {
			log.Fatalf("Fatal error unmarshalling config.yaml: %v", err)
		}
	}


	// 读取 gorm_v2.yaml
	// 注意: Gormv2 是顶层键，需要先读取 Gormv2Config，再从中获取 Mysql/SqlServer/PostgreSql 配置
	gormViper := viper.New()
	gormViper.SetConfigFile("gorm_v2.yaml")
	gormViper.SetConfigType("yaml")
	if err := gormViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("Warning: gorm_v2.yaml not found. Relying on environment variables. Error: %v", err)
		} else {
			log.Fatalf("Fatal error reading gorm_v2.yaml: %v", err)
		}
	} else {
		// 先解包到 GormV2Config 结构体
		if err := gormViper.Unmarshal(GlobalGormV2Config); err != nil {
			log.Fatalf("Fatal error unmarshalling gorm_v2.yaml: %v", err)
		}
	}


	// --- 环境变量覆盖逻辑 ---
	// AppDebug
	if debug := os.Getenv("APP_DEBUG"); debug != "" {
		GlobalServerConfig.AppDebug = strings.ToLower(debug) == "true" || debug == "1"
	}

	// HttpServer
	if port := os.Getenv("HTTP_SERVER_WEB_PORT"); port != "" {
		GlobalServerConfig.HttpServer.Web.Port = port
	}
	if allowCross := os.Getenv("HTTP_SERVER_ALLOW_CROSS_DOMAIN"); allowCross != "" {
		GlobalServerConfig.HttpServer.AllowCrossDomain = strings.ToLower(allowCross) == "true" || allowCross == "1"
	}
	if trustProxiesOpen := os.Getenv("HTTP_SERVER_TRUST_PROXIES_IS_OPEN"); trustProxiesOpen != "" {
		if val, err := strconv.Atoi(trustProxiesOpen); err == nil {
			GlobalServerConfig.HttpServer.TrustProxies.IsOpen = val
		}
	}
	if proxyListStr := os.Getenv("HTTP_SERVER_PROXY_SERVER_LIST"); proxyListStr != "" {
		GlobalServerConfig.HttpServer.TrustProxies.ProxyServerList = strings.Split(proxyListStr, ",")
	}


	// Token
	if key := os.Getenv("TOKEN_JWT_TOKEN_SIGN_KEY"); key != "" {
		GlobalServerConfig.Token.JwtTokenSignKey = key
	}
	if uid := os.Getenv("TOKEN_JWT_DEFAULT_UID"); uid != "" {
		if val, err := strconv.Atoi(uid); err == nil {
			GlobalServerConfig.Token.JwtDefaultUid = val
		}
	}
	if onlineUsers := os.Getenv("TOKEN_JWT_TOKEN_ONLINE_USERS"); onlineUsers != "" {
		if val, err := strconv.Atoi(onlineUsers); err == nil {
			GlobalServerConfig.Token.JwtTokenOnlineUsers = val
		}
	}
	if expireAt := os.Getenv("TOKEN_JWT_TOKEN_CREATED_EXPIRE_AT"); expireAt != "" {
		if val, err := strconv.Atoi(expireAt); err == nil {
			GlobalServerConfig.Token.JwtTokenCreatedExpireAt = val
		}
	}
	if refreshAllowSec := os.Getenv("TOKEN_JWT_TOKEN_REFRESH_ALLOW_SEC"); refreshAllowSec != "" {
		if val, err := strconv.Atoi(refreshAllowSec); err == nil {
			GlobalServerConfig.Token.JwtTokenRefreshAllowSec = val
		}
	}
	if refreshExpireAt := os.Getenv("TOKEN_JWT_TOKEN_REFRESH_EXPIRE_AT"); refreshExpireAt != "" {
		if val, err := strconv.Atoi(refreshExpireAt); err == nil {
			GlobalServerConfig.Token.JwtTokenRefreshExpireAt = val
		}
	}
	if bindKey := os.Getenv("TOKEN_BIND_CONTEXT_KEY_NAME"); bindKey != "" {
		GlobalServerConfig.Token.BindContextKeyName = bindKey
	}
	if cacheToRedis := os.Getenv("TOKEN_IS_CACHE_TO_REDIS"); cacheToRedis != "" {
		if val, err := strconv.Atoi(cacheToRedis); err == nil {
			GlobalServerConfig.Token.IsCacheToRedis = val
		}
	}


	// LoginPolicy
	if isOpenPolicy := os.Getenv("LOGIN_POLICY_IS_OPEN_POLICY"); isOpenPolicy != "" {
		if val, err := strconv.Atoi(isOpenPolicy); err == nil {
			GlobalServerConfig.LoginPolicy.IsOpenPolicy = val
		}
	}
	if maxFailTimes := os.Getenv("LOGIN_POLICY_MAX_LOGIN_FAIL_TIMES"); maxFailTimes != "" {
		if val, err := strconv.Atoi(maxFailTimes); err == nil {
			GlobalServerConfig.LoginPolicy.MaxLoginFailTimes = val
		}
	}
	if countDown := os.Getenv("LOGIN_POLICY_LOGIN_FAIL_COUNT_DOWN"); countDown != "" {
		if val, err := strconv.Atoi(countDown); err == nil {
			GlobalServerConfig.LoginPolicy.LoginFailCountDown = val
		}
	}

	// Redis
	if host := os.Getenv("REDIS_HOST"); host != "" {
		GlobalServerConfig.Redis.Host = host
	}
	if port := os.Getenv("REDIS_PORT"); port != "" {
		if val, err := strconv.Atoi(port); err == nil {
			GlobalServerConfig.Redis.Port = val
		}
	}
	if auth := os.Getenv("REDIS_AUTH"); auth != "" {
		GlobalServerConfig.Redis.Auth = auth
	}
	if maxIdle := os.Getenv("REDIS_MAX_IDLE"); maxIdle != "" {
		if val, err := strconv.Atoi(maxIdle); err == nil {
			GlobalServerConfig.Redis.MaxIdle = val
		}
	}
	if maxActive := os.Getenv("REDIS_MAX_ACTIVE"); maxActive != "" {
		if val, err := strconv.Atoi(maxActive); err == nil {
			GlobalServerConfig.Redis.MaxActive = val
		}
	}
	if idleTimeout := os.Getenv("REDIS_IDLE_TIMEOUT"); idleTimeout != "" {
		if val, err := strconv.Atoi(idleTimeout); err == nil {
			GlobalServerConfig.Redis.IdleTimeout = val
		}
	}
	if indexDb := os.Getenv("REDIS_INDEX_DB"); indexDb != "" {
		if val, err := strconv.Atoi(indexDb); err == nil {
			GlobalServerConfig.Redis.IndexDb = val
		}
	}
	if retryTimes := os.Getenv("REDIS_CONN_FAIL_RETRY_TIMES"); retryTimes != "" {
		if val, err := strconv.Atoi(retryTimes); err == nil {
			GlobalServerConfig.Redis.ConnFailRetryTimes = val
		}
	}
	if reconnectInterval := os.Getenv("REDIS_RECONNECT_INTERVAL"); reconnectInterval != "" {
		if val, err := strconv.Atoi(reconnectInterval); err == nil {
			GlobalServerConfig.Redis.ReConnectInterval = val
		}
	}

	// Logs (在Vercel上，这些通常不直接生效，因为日志会输出到控制台)
	if ginLogName := os.Getenv("LOGS_GIN_LOG_NAME"); ginLogName != "" {
		GlobalServerConfig.Logs.GinLogName = ginLogName
	}
	if goLogName := os.Getenv("LOGS_GO_SKELETON_LOG_NAME"); goLogName != "" {
		GlobalServerConfig.Logs.GoSkeletonLogName = goLogName
	}
	if textFormat := os.Getenv("LOGS_TEXT_FORMAT"); textFormat != "" {
		GlobalServerConfig.Logs.TextFormat = textFormat
	}
	if timePrecision := os.Getenv("LOGS_TIME_PRECISION"); timePrecision != "" {
		GlobalServerConfig.Logs.TimePrecision = timePrecision
	}
	if maxSize := os.Getenv("LOGS_MAX_SIZE"); maxSize != "" {
		if val, err := strconv.Atoi(maxSize); err == nil {
			GlobalServerConfig.Logs.MaxSize = val
		}
	}
	if maxBackups := os.Getenv("LOGS_MAX_BACKUPS"); maxBackups != "" {
		if val, err := strconv.Atoi(maxBackups); err == nil {
			GlobalServerConfig.Logs.MaxBackups = val
		}
	}
	if maxAge := os.Getenv("LOGS_MAX_AGE"); maxAge != "" {
		if val, err := strconv.Atoi(maxAge); err == nil {
			GlobalServerConfig.Logs.MaxAge = val
		}
	}
	if compress := os.Getenv("LOGS_COMPRESS"); compress != "" {
		GlobalServerConfig.Logs.Compress = strings.ToLower(compress) == "true" || compress == "1"
	}

	// Websocket
	if start := os.Getenv("WEBSOCKET_START"); start != "" {
		if val, err := strconv.Atoi(start); err == nil {
			GlobalServerConfig.Websocket.Start = val
		}
	}
	if bufferSize := os.Getenv("WEBSOCKET_WRITE_READ_BUFFER_SIZE"); bufferSize != "" {
		if val, err := strconv.Atoi(bufferSize); err == nil {
			GlobalServerConfig.Websocket.WriteReadBufferSize = val
		}
	}
	if maxMsgSize := os.Getenv("WEBSOCKET_MAX_MESSAGE_SIZE"); maxMsgSize != "" {
		if val, err := strconv.Atoi(maxMsgSize); err == nil {
			GlobalServerConfig.Websocket.MaxMessageSize = val
		}
	}
	if pingPeriod := os.Getenv("WEBSOCKET_PING_PERIOD"); pingPeriod != "" {
		if val, err := strconv.Atoi(pingPeriod); err == nil {
			GlobalServerConfig.Websocket.PingPeriod = val
		}
	}
	if failMaxTimes := os.Getenv("WEBSOCKET_HEARTBEAT_FAIL_MAX_TIMES"); failMaxTimes != "" {
		if val, err := strconv.Atoi(failMaxTimes); err == nil {
			GlobalServerConfig.Websocket.HeartbeatFailMaxTimes = val
		}
	}
	if readDeadline := os.Getenv("WEBSOCKET_READ_DEADLINE"); readDeadline != "" {
		if val, err := strconv.Atoi(readDeadline); err == nil {
			GlobalServerConfig.Websocket.ReadDeadline = val
		}
	}
	if writeDeadline := os.Getenv("WEBSOCKET_WRITE_DEADLINE"); writeDeadline != "" {
		if val, err := strconv.Atoi(writeDeadline); err == nil {
			GlobalServerConfig.Websocket.WriteDeadline = val
		}
	}

	// SnowFlake
	if machineId := os.Getenv("SNOWFLAKE_MACHINE_ID"); machineId != "" {
		if val, err := strconv.Atoi(machineId); err == nil {
			GlobalServerConfig.SnowFlake.SnowFlakeMachineId = val
		}
	}

	// FileUploadSetting
	if size := os.Getenv("FILE_UPLOAD_SETTING_SIZE"); size != "" {
		if val, err := strconv.Atoi(size); err == nil {
			GlobalServerConfig.FileUploadSetting.Size = val
		}
	}
	if field := os.Getenv("FILE_UPLOAD_SETTING_UPLOAD_FILE_FIELD"); field != "" {
		GlobalServerConfig.FileUploadSetting.UploadFileField = field
	}
	if urlPrefix := os.Getenv("FILE_UPLOAD_SETTING_SOURCE_URL_PREFIX"); urlPrefix != "" {
		GlobalServerConfig.FileUploadSetting.SourceUrlPrefix = urlPrefix
	}
	if rootPath := os.Getenv("FILE_UPLOAD_SETTING_UPLOAD_ROOT_PATH"); rootPath != "" {
		GlobalServerConfig.FileUploadSetting.UploadRootPath = rootPath
	}
	if avatarSmallPath := os.Getenv("FILE_UPLOAD_SETTING_AVATAR_SMALL_UPLOAD_FILE_SAVE_PATH"); avatarSmallPath != "" {
		GlobalServerConfig.FileUploadSetting.AvatarSmallUploadFileSavePath = avatarSmallPath
	}
	if avatarLargePath := os.Getenv("FILE_UPLOAD_SETTING_AVATAR_LARGE_UPLOAD_FILE_SAVE_PATH"); avatarLargePath != "" {
		GlobalServerConfig.FileUploadSetting.AvatarLargeUploadFileSavePath = avatarLargePath
	}
	if coverPath := os.Getenv("FILE_UPLOAD_SETTING_COVER_UPLOAD_FILE_SAVE_PATH"); coverPath != "" {
		GlobalServerConfig.FileUploadSetting.CoverUploadFileSavePath = coverPath
	}
	if whiteCoverPath := os.Getenv("FILE_UPLOAD_SETTING_WHITE_COVER_UPLOAD_FILE_SAVE_PATH"); whiteCoverPath != "" {
		GlobalServerConfig.FileUploadSetting.WhiteCoverUploadFileSavePath = whiteCoverPath
	}
	if videoPath := os.Getenv("FILE_UPLOAD_SETTING_VIDEO_UPLOAD_FILE_SAVE_PATH"); videoPath != "" {
		GlobalServerConfig.FileUploadSetting.VideoUploadFileSavePath = videoPath
	}
	if videoCoverPath := os.Getenv("FILE_UPLOAD_SETTING_VIDEO_COVER_UPLOAD_FILE_SAVE_PATH"); videoCoverPath != "" {
		GlobalServerConfig.FileUploadSetting.VideoCoverUploadFileSavePath = videoCoverPath
	}
	if allowMimeType := os.Getenv("FILE_UPLOAD_SETTING_ALLOW_MIME_TYPE"); allowMimeType != "" {
		GlobalServerConfig.FileUploadSetting.AllowMimeType = strings.Split(allowMimeType, ",")
	}

	// RabbitMq (如果不需要在Vercel上使用，可以跳过配置)
	// HelloWorld
	if addr := os.Getenv("RABBITMQ_HELLO_WORLD_ADDR"); addr != "" {
		GlobalServerConfig.RabbitMq.HelloWorld.Addr = addr
	}
	if queueName := os.Getenv("RABBITMQ_HELLO_WORLD_QUEUE_NAME"); queueName != "" {
		GlobalServerConfig.RabbitMq.HelloWorld.QueueName = queueName
	}
	// ... 其他RabbitMq配置类似

	// Casbin
	if isInit := os.Getenv("CASBIN_IS_INIT"); isInit != "" {
		if val, err := strconv.Atoi(isInit); err == nil {
			GlobalServerConfig.Casbin.IsInit = val
		}
	}
	if autoLoadSeconds := os.Getenv("CASBIN_AUTO_LOAD_POLICY_SECONDS"); autoLoadSeconds != "" {
		if val, err := strconv.Atoi(autoLoadSeconds); err == nil {
			GlobalServerConfig.Casbin.AutoLoadPolicySeconds = val
		}
	}
	if tablePrefix := os.Getenv("CASBIN_TABLE_PREFIX"); tablePrefix != "" {
		GlobalServerConfig.Casbin.TablePrefix = tablePrefix
	}
	if tableName := os.Getenv("CASBIN_TABLE_NAME"); tableName != "" {
		GlobalServerConfig.Casbin.TableName = tableName
	}
	if modelConfig := os.Getenv("CASBIN_MODEL_CONFIG"); modelConfig != "" {
		GlobalServerConfig.Casbin.ModelConfig = modelConfig
	}

	// Captcha
	if captchaId := os.Getenv("CAPTCHA_CAPTCHA_ID"); captchaId != "" {
		GlobalServerConfig.Captcha.CaptchaId = captchaId
	}
	if captchaValue := os.Getenv("CAPTCHA_CAPTCHA_VALUE"); captchaValue != "" {
		GlobalServerConfig.Captcha.CaptchaValue = captchaValue
	}
	if length := os.Getenv("CAPTCHA_LENGTH"); length != "" {
		if val, err := strconv.Atoi(length); err == nil {
			GlobalServerConfig.Captcha.Length = val
		}
	}


	// --- Gormv2 MySQL数据库配置 (来自 gorm_v2.yaml) ---
	if dbType := os.Getenv("GORM_USE_DB_TYPE"); dbType != "" {
		GlobalGormV2Config.UseDbType = dbType
	}

	// MySQL Write DB
	if isInitGlobalGormMysql := os.Getenv("GORM_MYSQL_IS_INIT_GLOBAL_GORM_MYSQL"); isInitGlobalGormMysql != "" {
		if val, err := strconv.Atoi(isInitGlobalGormMysql); err == nil {
			GlobalGormV2Config.Mysql.IsInitGlobalGormMysql = val
		}
	}
	if dbHost := os.Getenv("MYSQL_DB_WRITE_HOST"); dbHost != "" {
		GlobalGormV2Config.Mysql.Write.Host = dbHost
	}
	if dbName := os.Getenv("MYSQL_DB_WRITE_DATABASE"); dbName != "" {
		GlobalGormV2Config.Mysql.Write.DataBase = dbName
	}
	if dbPort := os.Getenv("MYSQL_DB_WRITE_PORT"); dbPort != "" {
		if p, err := strconv.Atoi(dbPort); err == nil {
			GlobalGormV2Config.Mysql.Write.Port = p
		}
	}
	if dbUser := os.Getenv("MYSQL_DB_WRITE_USER"); dbUser != "" {
		GlobalGormV2Config.Mysql.Write.User = dbUser
	}
	if dbPassword := os.Getenv("MYSQL_DB_WRITE_PASS"); dbPassword != "" {
		GlobalGormV2Config.Mysql.Write.Pass = dbPassword
	}
	if charset := os.Getenv("MYSQL_DB_WRITE_CHARSET"); charset != "" {
		GlobalGormV2Config.Mysql.Write.Charset = charset
	}
	if maxIdle := os.Getenv("MYSQL_DB_WRITE_SET_MAX_IDLE_CONNS"); maxIdle != "" {
		if val, err := strconv.Atoi(maxIdle); err == nil {
			GlobalGormV2Config.Mysql.Write.SetMaxIdleConns = val
		}
	}
	if maxOpen := os.Getenv("MYSQL_DB_WRITE_SET_MAX_OPEN_CONNS"); maxOpen != "" {
		if val, err := strconv.Atoi(maxOpen); err == nil {
			GlobalGormV2Config.Mysql.Write.SetMaxOpenConns = val
		}
	}
	if lifetime := os.Getenv("MYSQL_DB_WRITE_SET_CONN_MAX_LIFE_TIME"); lifetime != "" {
		if val, err := strconv.Atoi(lifetime); err == nil {
			GlobalGormV2Config.Mysql.Write.SetConnMaxLifetime = val
		}
	}

	// MySQL Read DB (如果 IsOpenReadDb 为 1 才会生效)
	if isOpenReadDb := os.Getenv("GORM_MYSQL_IS_OPEN_READ_DB"); isOpenReadDb != "" {
		if val, err := strconv.Atoi(isOpenReadDb); err == nil {
			GlobalGormV2Config.Mysql.IsOpenReadDb = val
		}
	}
	if GlobalGormV2Config.Mysql.IsOpenReadDb == 1 { // 只有开启读写分离才从环境变量读取读库配置
		if dbHost := os.Getenv("MYSQL_DB_READ_HOST"); dbHost != "" {
			GlobalGormV2Config.Mysql.Read.Host = dbHost
		}
		if dbName := os.Getenv("MYSQL_DB_READ_DATABASE"); dbName != "" {
			GlobalGormV2Config.Mysql.Read.DataBase = dbName
		}
		if dbPort := os.Getenv("MYSQL_DB_READ_PORT"); dbPort != "" {
			if p, err := strconv.Atoi(dbPort); err == nil {
				GlobalGormV2Config.Mysql.Read.Port = p
			}
		}
		if dbUser := os.Getenv("MYSQL_DB_READ_USER"); dbUser != "" {
			GlobalGormV2Config.Mysql.Read.User = dbUser
		}
		if dbPassword := os.Getenv("MYSQL_DB_READ_PASS"); dbPassword != "" {
			GlobalGormV2Config.Mysql.Read.Pass = dbPassword
		}
		if charset := os.Getenv("MYSQL_DB_READ_CHARSET"); charset != "" {
			GlobalGormV2Config.Mysql.Read.Charset = charset
		}
		if maxIdle := os.Getenv("MYSQL_DB_READ_SET_MAX_IDLE_CONNS"); maxIdle != "" {
			if val, err := strconv.Atoi(maxIdle); err == nil {
				GlobalGormV2Config.Mysql.Read.SetMaxIdleConns = val
			}
		}
		if maxOpen := os.Getenv("MYSQL_DB_READ_SET_MAX_OPEN_CONNS"); maxOpen != "" {
			if val, err := strconv.Atoi(maxOpen); err == nil {
				GlobalGormV2Config.Mysql.Read.SetMaxOpenConns = val
			}
		}
		if lifetime := os.Getenv("MYSQL_DB_READ_SET_CONN_MAX_LIFE_TIME"); lifetime != "" {
			if val, err := strconv.Atoi(lifetime); err == nil {
				GlobalGormV2Config.Mysql.Read.SetConnMaxLifetime = val
			}
		}
	}
	// --- 其他数据库 (SqlServer, PostgreSql) 如果你需要使用，也需要类似地添加环境变量覆盖逻辑 ---


	// 监听配置文件变化 (本地开发时有用，Vercel上无意义)
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		// 重新读取 config.yaml
		if err := viper.Unmarshal(GlobalServerConfig); err != nil {
			log.Fatalf("Fatal error re-unmarshalling config.yaml: %v", err)
		}
		// 重新读取 gorm_v2.yaml
		if err := gormViper.Unmarshal(GlobalGormV2Config); err != nil { // 注意这里是 gormViper
			log.Fatalf("Fatal error re-unmarshalling gorm_v2.yaml: %v", err)
		}
	})
}
