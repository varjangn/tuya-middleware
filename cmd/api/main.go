package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/varjangn/tuya-middleware/config"
	"github.com/varjangn/tuya-middleware/pkg/logger"
	"github.com/varjangn/tuya-middleware/pkg/tuya"
	"github.com/varjangn/tuya-middleware/utils"
)

func main() {

	configPath := utils.GetConfigPath(os.Getenv("config"))
	cfgFile, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("LoadConfig: %v", err)
	}
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}

	appLogger := logger.NewAppLogger(cfg)
	appLogger.InitLogger()

	tuyaClient := tuya.NewTuyaClient(appLogger, cfg)

	// run goroutine to auto refresh tuya token
	go tuyaClient.AutoRefreshToken()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "running...\n")
	})
	http.ListenAndServe(cfg.Server.Port, nil)

}
