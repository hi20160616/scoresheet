package configs

import (
	"fmt"
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg := NewConfig("scoresheet")
	if cfg.Err != nil {
		t.Error(cfg.Err)
	}
	fmt.Println(cfg)
}

func TestRootPath4Test(t *testing.T) {
	cfg := &Config{ProjectName: "scoresheet"}
	cfg.RootPath, cfg.Err = os.Getwd()
	if cfg.Err != nil {
		t.Error(cfg.Err)
	}
	cfg = rootPath4Test(cfg)
	if cfg.Err != nil {
		t.Error(cfg.Err)
	}
	fmt.Println(cfg.RootPath)
}
