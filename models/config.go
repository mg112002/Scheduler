package models

var Config Configuration

type Configuration struct {
	Mysql       Mysql             `json:"mysql"`
	CacheSizeGb int               `json:"cacheSizeGb"`
	Tables      map[string]string `json:"tables"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
