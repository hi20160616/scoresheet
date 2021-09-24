package db

import (
	"path/filepath"

	"github.com/hi20160616/gears"
	"github.com/hi20160616/scoresheet"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	Path string
	File *excelize.File
	Err  error
}

type Excels struct {
	es []*Excel
}

func NewExcel(src string, opts *scoresheet.Opts) *Excel {
	e := &Excel{Path: src}
	isExcel := func() bool {
		srcExt := filepath.Ext(src)
		for _, e := range opts.Cfg.DBType {
			if e == srcExt {
				return true
			}
		}
		e.Err = scoresheet.ErrTypeExcel
		return false
	}()

	if isExcel {
		if gears.Exists(src) {
			e.File, e.Err = excelize.OpenFile(e.Path)
		} else {
			f := excelize.NewFile()
			// Create a new sheet.
			index := f.NewSheet("Sheet1")
			// Set value of a cell.
			f.SetCellValue("Sheet1", "A1", "Hello world.")
			f.SetCellValue("Sheet1", "B2", 100)
			// Set active sheet of the workbook.
			f.SetActiveSheet(index)
			// Save spreadsheet by the given path.
			if err := f.SaveAs(src); err != nil {
				e.Err = err
			}
			e.File = f
		}
	}
	return e
}
