package restfulexcel

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//Route is restful excel api
func Route(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		get(res, req)
	case "POST":
		post(res, req)
	case "PUT":

	case "DELETE":

	default:
	}
}

func get(res http.ResponseWriter, req *http.Request) {

	path := strings.ToLower(req.URL.Path)
	fileName := path[1:]

	var sheets []Sheet
	if xlFile, err := xlsx.OpenFile(fileName); err == nil {
		sheets = make([]Sheet, len(xlFile.Sheets))
		for s, sheet := range xlFile.Sheets {
			sheets[s].Name = sheet.Name
			sheets[s].Table = make([][]string, sheet.MaxRow)
			sheets[s].ColType = make([]string, sheet.MaxCol)
			for r, row := range sheet.Rows {
				sheets[s].Table[r] = make([]string, sheet.MaxCol)
				for c, cell := range row.Cells {
					sheets[s].Table[r][c] = cell.String()
				}
			}
			latestRow := sheet.Row(sheet.MaxRow - 1)
			for c, cell := range latestRow.Cells {
				switch cell.Type() {
				case xlsx.CellTypeBool:
					sheets[s].ColType[c] = "select"
				case xlsx.CellTypeDate: // 此类型不会出现, 真正的date类型为 xlsx.CellTypeNumeric
					sheets[s].ColType[c] = "date"
				case xlsx.CellTypeNumeric:
					if _, err := strconv.Atoi(cell.String()); err != nil {
						sheets[s].ColType[c] = "date"
					} else {
						sheets[s].ColType[c] = "number"
					}
				case xlsx.CellTypeString:
					sheets[s].ColType[c] = "text"
				case xlsx.CellTypeStringFormula:
					sheets[s].ColType[c] = "text"
				case xlsx.CellTypeError:
					sheets[s].ColType[c] = "text"
				case xlsx.CellTypeInline:
					sheets[s].ColType[c] = "text"
				}
			}

		}
	} else {
		fmt.Println(err.Error())
	}

	totalTemplate.Execute(res, TemplateData{
		Title:  fileName,
		Sheets: sheets,
	})
}

func post(res http.ResponseWriter, req *http.Request) {
	if req.ParseForm() != nil {
		http.Error(res, "", http.StatusNotAcceptable)
		return
	}
	path := strings.ToLower(req.URL.Path)
	fileName := path[1:]
	sheetName := req.PostFormValue("sheet")
	if xlFile, err := xlsx.OpenFile(fileName); err == nil {
		sheet := xlFile.Sheet[sheetName]
		latestRow := sheet.Row(sheet.MaxRow - 1)
		newRow := sheet.AddRow()
		for i := 0; i < sheet.MaxCol; i++ {
			if formValue := req.PostFormValue(strconv.Itoa(i)); formValue != "" {
				switch latestRow.Cells[i].Type() {
				case xlsx.CellTypeBool:
					if b, err := strconv.ParseBool(formValue); err == nil {
						newRow.AddCell().SetBool(b)
					} else {
						http.Error(res, "", http.StatusBadRequest)
						return
					}
				case xlsx.CellTypeDate:

				case xlsx.CellTypeNumeric:
					if i, err := strconv.Atoi(formValue); err == nil {
						newRow.AddCell().SetInt(i)
					} else {
						if t, err := time.Parse("2006-01-02", formValue); err == nil {
							newRow.AddCell().SetDate(t)
						} else {
							fmt.Println(err.Error())
							http.Error(res, "", http.StatusBadRequest)
							return
						}
					}
				case xlsx.CellTypeString:
					newRow.AddCell().SetString(formValue)
				case xlsx.CellTypeStringFormula:

				case xlsx.CellTypeError:

				case xlsx.CellTypeInline:

				}
			}

		}
		xlFile.Save(fileName)
	} else {
		http.Error(res, "", http.StatusNotFound)
	}

}
