package config

type (
	Config struct {
		Bot BotConfig
		Nsq NsqConfig
	}

	BotConfig struct {
		AppId     uint64
		Token     string
		SecretKey string
	}

	NsqConfig struct {
		Addr  string
		Topic string
	}
)
