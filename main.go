package main

import (
	"encoding/gob"
	"log"
	"math/rand"
	"runtime"
	"time"

	"github.com/caarlos0/env"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/gorp.v1"
)

// Global variables
var dbMap *gorp.DbMap
var cfg Env

func main() {
	// Read configuration environment variables
	cfg = Env{}
	err := env.Parse(&cfg)
	checkErr(err, "ERROR:")

	// Application secret
	if cfg.Secret == "" {
		token, err := generateRandomString(32)
		cfg.Secret = token
		checkErr(err, "Secret:")
	}

	// Register user login session struct
	gob.Register(UserLogin{})

	// Performance and deployment related settings
	runtime.GOMAXPROCS(runtime.NumCPU())
	if cfg.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize random number seed
	rand.Seed(time.Now().Unix())

	// Initialize gin request router
	router := gin.Default()
	store := sessions.NewCookieStore([]byte(cfg.Secret))
	router.Use(sessions.Sessions("login", store))
	router.LoadHTMLGlob("res/templates/*")
	router.Static("/static", "res/static")

	// Configure routes
	routesConfig(router)

	// Connect to database
	dbMap = initDb(cfg.DBHost, cfg.DBSocket, cfg.DBConnType, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.Production)

	// Listen on unix socket or tcp based on commandline flags
	log.Println("cfg unix socket: " + cfg.AppSocket)

	if cfg.ListenType == 1 {
		router.Run(cfg.ListenIP + ":" + cfg.Port)
	} else {
		router.RunUnix(cfg.AppSocket)
	}
}
