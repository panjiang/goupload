package global

import (
	"encoding/json"
	"log"
	"os"
)

var (
	Conf *Config
)

type Config struct {
	Release       int
	BindPort      int
	UploadDir     string
	AllowFileType []string
	MaxFileSize   int
	HTTPServerURL string
	FTPServer     int
	FTPServerAddr string
	FTPUploadPath string
	FTPUsername   string
	FTPPassword   string
}

func Parser(configFile string) {
	file, _ := os.Open(configFile)
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&Conf)
	if err != nil {
		log.Panicln("parser config error:", err)
	}
	log.Println("parser config success")
}
