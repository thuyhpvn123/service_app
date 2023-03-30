package config

type IConfig interface {
	GetVersion() string
	GetNodeType() string
	GetPrivateKey() []byte
}
