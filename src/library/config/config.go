package config

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
)

var Cfg = &config{
	kv: map[string]string{
		"server_id":   "0x84a5d1600002c001",
		"server_name": "xixihaha",
	},
}

func itemNotFound(item string) error {
	return errors.New(fmt.Sprintf("config: %s not found\n", item))
}

type config struct {
	sync.Mutex
	kv map[string]string
}

func (cfg *config) GetString(key string) (string, error) {
	cfg.Lock()
	val, ok := cfg.kv[key]
	cfg.Unlock()
	if !ok {
		return val, itemNotFound(key)
	}
	return val, nil
}

func (cfg *config) GetUint64(key string) (uint64, error) {
	val, err := cfg.GetString(key)
	if err != nil {
		return 0, err
	}
	ui64, err := strconv.ParseUint(val, 10, 64)
	return ui64, err
}

func (cfg *config) GetFloat64(key string) (float64, error) {
	val, err := cfg.GetString(key)
	if err != nil {
		return 0, err
	}
	ui64, err := strconv.ParseFloat(val, 64)
	return ui64, err
}
