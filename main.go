// 51cz project main.go
package main

import (
	"51cz/service/dbcomm"
	"flag"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	goconf "github.com/pantsing/goconf"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	http_srv   *http.Server
	configMap  = make(map[string]string, 0)
	dbUrl      string
	ccdbUrl    string
	listenPort int
	idleConns  int
	openConns  int
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile | log.Lmicroseconds)
	log.SetOutput(io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   "51cz.log",
		MaxSize:    500, // megabytes
		MaxBackups: 50,
		MaxAge:     90, //days
	}))
	envConf := flag.String("env", "config.json", "select a environment config file")
	flag.Parse()
	log.Println("config file ==", *envConf)
	c, err := goconf.New(*envConf)
	if err != nil {
		log.Fatalln("读配置文件出错", err)
	}
	//填充配置文件
	c.Get("/config/LISTEN_PORT", &listenPort)
	c.Get("/config/DB_URL", &dbUrl)
	c.Get("/config/OPEN_CONNS", &openConns)
	c.Get("/config/IDLE_CONNS", &idleConns)
	dbcomm.InitDB(dbUrl, openConns, openConns)
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	router.GET("/", index)

	router.POST("/test", test)

	router.Run(":8090")

}
