package db

import (
	"testing"

	"github.com/hi20160616/scoresheet"
	"github.com/hi20160616/scoresheet/configs"
)

func TestNewExcel(t *testing.T) {
	cfg := configs.NewConfig("scoresheet")
	log, err := scoresheet.NewLog(cfg)
	if err != nil {
		t.Error(err)
		return
	}
	opts := scoresheet.NewOpts(cfg, log)
	fpath := "test.xlsx"
	f := NewExcel(fpath, opts)
	if f.Err != nil {
		t.Error(f.Err)
		return
	}
	xls := f.File
	i := xls.NewSheet("Sheet1")
	xls.SetActiveSheet(i)
	xls.SetCellValue("Sheet1", "A2", "test")
	xls.SetCellValue("Sheet1", "B2", 200)
	if err := xls.Save(); err != nil {
		t.Error(err)
	}
}
