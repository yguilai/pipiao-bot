package conf

// BotConf qq bot configuration
type BotConf struct {
	AppId     uint64
	Token     string
	SecretKey string `json:",optional"`
}

type NsqConf struct {
	Addr   string
	Lookup string
	Topic  string
}
