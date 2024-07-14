package tuya

import (
	"github.com/varjangn/tuya-middleware/config"
	"github.com/varjangn/tuya-middleware/pkg/logger"
)

type TuyaClient struct {
	logger *logger.AppLogger
	cfg    *config.Tuya
}

func NewTuyaClient(logger *logger.AppLogger, cfg *config.Config) *TuyaClient {
	return &TuyaClient{logger: logger, cfg: &cfg.Tuya}
}
