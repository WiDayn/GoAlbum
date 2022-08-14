package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"os"
)

type config struct {
	Debug     bool
	Port      int
	Proxy     Proxy
	SQLServer SQLServer
}

type Proxy struct {
	Enable bool
	Url    string
}

type SQLServer struct {
	IP       string
	Port     int
	Username string
	Password string
	Db       string
	Charset  string
	Loc      string
}

var Config config

func Read() {
	if !isExist("./config.ini") {
		if _, err := os.Create("./config.ini"); err != nil {
			log.Println("[PANIC]Can not create config.ini")
			log.Fatalf(err.Error())
		}
	}

	if cfg, err := ini.Load("./config.ini"); err != nil {
		log.Println("[PANIC] Can not load config.ini")
		log.Fatalf(err.Error())
	} else {
		Config = config{
			Debug: false,
			Port:  12345,
			SQLServer: SQLServer{
				IP:       "127.0.0.1",
				Port:     3306,
				Username: "root",
				Password: "TqOMaR6zrMQUm3BG",
				Db:       "douban",
				Charset:  "utf8mb4",
				Loc:      "Local",
			},
		}

		if err = ini.MapTo(&Config, "./config.ini"); err != nil {
			log.Println("[PANIC] config.ini error")
			log.Fatalf(err.Error())
		}

		if err := ini.ReflectFrom(cfg, &Config); err != nil {
			log.Println("[PANIC] config.ini error")
			log.Fatalf(err.Error())
		}

		if err := cfg.SaveTo("./config.ini"); err != nil {
			log.Println("[PANIC] Can not sava config.ini")
			log.Fatalf(err.Error())
		}
	}
	log.Println("[SUCCESS] Success load config.ini")
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		fmt.Println(err)
		return false
	}
	return true
}
