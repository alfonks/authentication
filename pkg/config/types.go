package config

type ConfigStr struct {
	Mongo       mongoConfig      `json:"mongo"`
	Redis       redisConfig      `json:"redis"`
	Environment string           `json:"environment"`
	API         apiConfig        `json:"api"`
	Encrpytion  encryptionConfig `json:"encryption"`
	JWT         jwtConfig        `json:"jwt"`
}

type mongoConfig struct {
	AppName string   `json:"app_name"`
	Slave   dBConfig `json:"slave"`
	Master  dBConfig `json:"master"`
}
type dBConfig struct {
	URI           string `json:"uri"`
	Timeout       int64  `json:"time"`
	MaxConnection int64  `json:"max_connection"`
	MaxPoolSize   int64  `json:"max_pool_size"`
	MaxIdleTime   int64  `json:"max_idle_time"`
}

type redisConfig struct {
	URL         string `json:"url"`
	MaxRetries  int64  `json:"max_retries"`
	DialTimeout int64  `json:"dial_timeout"`
	IdleTimeout int64  `json:"idle_timeout"`
}

type apiConfig struct {
	DefaultTimeout int64 `json:"default_timeout"`
}

type encryptionConfig struct {
	Cost       int64  `json:"cost"`
	PublicKey  string `json:"public_key"`
	PrivateKey string `json:"private_key"`
}

type jwtConfig struct {
	SecretKey              string `json:"secret_key"`
	ExpiryTime             int64  `json:"expiry_time"`
	RefreshTokenExpiryTime int64  `json:"refresh_token_expiry_time"`
}
