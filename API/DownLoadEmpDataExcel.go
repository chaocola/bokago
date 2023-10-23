package API

import (
	"bytes"
	"fmt"
	"github.com/oddbug/bokago/DefaultConfig"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
	"sync"
)

var LetterArr = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var HeaderIndex = []interface{}{"工号", "姓名", "总现金", "卡金", "总疗程", "美发疗程", "美容疗程", "总消疗", "产品", "消卡/点数", "工资"}

var SelectMap = map[string]int{
	// 表1
	"工号": 0, "姓名": 1, "职称": 3, "总现金": 4, "产品": 16, "卡金": 24, "美容课程": 30, "业绩": 54, "工资": 59,
	// 表2
	"美容消疗": 125, "美发课程": 5, "美发消疗": 23,
}

var ExclusionNumber = []string{"777", "808", "810", "500", "300", "555"}

func DownExcel(url string) map[string][]string {

	var err error
	var row, row2 []string

	res := DefaultConfig.BOKA.GET(url, nil)
	// 直接读取下载的byte转换成io数据流
	f, _ := excelize.OpenReader(bytes.NewReader(res))
	//cell, _ := f.GetCellValue("专业周计表(汇总)", "A3")
	// 读取 专业周计表(汇总)
	rows, err := f.Rows(f.GetSheetName(0))
	// 读取 专业周计表(类别详情)
	rows2, err := f.Rows(f.GetSheetName(1))
	defer func() {
		_ = rows.Close()
		_ = rows2.Close()
		_ = f.Close()
	}()

	if err != nil {
		fmt.Println(err)
	}
	// 处理表1
	userData := map[string][]string{}
	for rows.Next() {

		row, err = rows.Columns()
		if err != nil {
			fmt.Println(err)
		}
		if len(row) > 0 && len(row[0]) == 3 {
			userData[row[SelectMap["工号"]]] = []string{
				row[SelectMap["工号"]],
				row[SelectMap["姓名"]],
				row[SelectMap["职称"]],
				row[SelectMap["总现金"]],
				row[SelectMap["产品"]],
				row[SelectMap["美容课程"]],
				row[SelectMap["业绩"]],
				row[SelectMap["工资"]],
				row[SelectMap["卡金"]],
			}
		}

	}
	// 处理表2
	for rows2.Next() {
		row2, err = rows2.Columns()
		if err != nil {
			fmt.Println(err)
		}
		if len(row2) > 0 && len(row2[0]) == 3 {

			userData[row2[SelectMap["工号"]]] = append(
				userData[row2[SelectMap["工号"]]],
				row2[SelectMap["美容消疗"]],
				row2[SelectMap["美发课程"]],
				row2[SelectMap["美发消疗"]],
			)

		}
	}

	return userData
}

func GetTwoExcel(stime string, etime string) *excelize.File {
	url1 := fmt.Sprintf("https://s3.boka.vc/api/report/export/weeklyProfessionExcel?v=1&compid=002&compName=孔雀宫-迎春路&fromdate=%s&todate=%s&fromdepart=&todepart=&fromstaff=&tostaff=&isquit=false&fromprj=&toprj=&fromprjgroup=&toprjgroup=&amttype=3&jobtype=1&jobs=&jobsName=全部&recalculate=false&incEmpty=0&pnum=2&pamt=2&srvName1=老客&srvName2=轮班&srvName14=内创&srvName15=外创&countZong=11&countPrd=4&countCard=6&countZa=10&countTicheng=7&aZongLs=[true,true,true,true,true,true,true,true,true,true,true]&aPrdLs=[true,true,true,true]&aCardLs=[true,true,true,true]&aZaLs=[true,true,true,true,true,true,true,true]&aPrj=true&aPrjLs=[true,true,true,true,true,true,true,true,true,true]&aPrjGp=true&aPrjGpLs=[true,true,true,true,true,true,true]&aTicLs=[true,true,true,true,true,true,true]", stime, etime)
	url2 := fmt.Sprintf(`https://s3.boka.vc/api/report/export/weeklyProfessionExcel?v=1&compid=002&compName=孔雀宫-迎春路&fromdate=%s&todate=%s&fromdepart=&todepart=&fromstaff=&tostaff=&isquit=false&fromprj=&toprj=&fromprjgroup=&toprjgroup=&amttype=0&jobtype=1&jobs=&jobsName=全部&recalculate=false&incEmpty=0&pnum=2&pamt=2&srvName1=老客&srvName2=轮班&srvName14=内创&srvName15=外创&countZong=11&countPrd=4&countCard=6&countZa=10&countTicheng=7&aZongLs=[true,true,true,true,true,true,true,true,true,true,true]&aPrdLs=[true,true,true,true]&aCardLs=[true,true,true,true]&aZaLs=[true,true,true,true,true,true,true,true]&aPrj=true&aPrjLs=[true,true,true,true,true,true,true,true,true,true]&aPrjGp=true&aPrjGpLs=[true,true,true,true,true,true,true]&aTicLs=[true,true,true,true,true,true,true]`, stime, etime)

	var wg sync.WaitGroup
	wg.Add(2)
	var userData map[string][]string
	var empData map[string][]string
	go func() {
		userData = DownExcel(url1)
		wg.Done()
	}()

	go func() {
		empData = DownExcel(url2)
		wg.Done()
	}()

	wg.Wait()

	for k, v := range userData {
		if string(k[0]) == "3" {
			for m, n := range empData {
				if m == k {
					v[10] = n[10]
					v[11] = n[11]
				}
			}
		}
		userData[k] = v
	}

	data := make([][]string, 0)
	for _, v := range userData {
		data = append(data, v)
	}
	sort.Slice(data, func(i, j int) bool {
		return data[i][0] < data[j][0]
	})
	file := WriteOption{
		File:      excelize.NewFile(),
		SheetName: "工资表",
	}
	// 总表
	_ = file.File.SetSheetName(file.File.GetSheetName(0), file.SheetName)
	_ = file.File.SetRowHeight(file.SheetName, 1, 23)

	startLetter := LetterArr[0]
	endLetter := LetterArr[len(data[0])-1]

	_ = file.File.SetCellStyle(file.SheetName, startLetter+"1", endLetter+"1", HeaderStyle(file.File))
	file.WriteCells(1, HeaderIndex)

	InArray := func(arr []string, str string) bool {
		for _, v := range arr {
			if v == str {
				return true
			}
		}
		return false
	}

	num := 1

	for _, v := range data {
		num++
		if InArray(ExclusionNumber, v[0]) {
			num--
			continue
		}
		_ = file.File.SetRowHeight(file.SheetName, num, 18)
		_ = file.File.SetCellStyle(file.SheetName, fmt.Sprintf("%v%d", startLetter, num), fmt.Sprintf("%v%d", endLetter, num), BodyStyle(file.File, num))
		file.WriteCells(num, []interface{}{v[0], v[1], STF(v[3]), StrInt(v[8], v[5], "-"), StrInt(v[10], v[5], "+"), STF(v[10]), STF(v[5]), StrInt(v[9], v[11], "+"), STF(v[4]), FloatInt(v[6]), FloatInt(v[7])})
		_ = file.File.AutoFilter(file.SheetName, "A1:K26", nil)
	}

	return file.File
}

func STF(s string) int64 {
	s1, _ := strconv.ParseFloat(s, 64)
	return int64(s1)
}

func StrInt(str1 string, str2 string, c string) int64 {
	s1, err := strconv.ParseFloat(str1, 64)
	s2, err := strconv.ParseFloat(str2, 64)
	if err != nil {
		fmt.Println(err)
	}
	if c == "+" {
		return int64(s1 + s2)
	}

	return int64(s1 - s2)

}
func FloatInt(str string) int64 {
	s1, _ := strconv.ParseFloat(str, 64)
	return int64(s1)
}

type WriteOption struct {
	File      *excelize.File
	SheetName string
}

func (f WriteOption) WriteCells(num int, value []interface{}) {

	for i := 0; i < len(value); i++ {
		word := LetterArr[i]
		err := f.File.SetCellValue(f.SheetName, fmt.Sprintf(`%s%d`, word, num), value[i])
		if err != nil {
			fmt.Println(err)
		}
	}

}

func HeaderStyle(file *excelize.File) int {
	headStyle, _ := file.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
			Size: 12,
		},
		Fill: excelize.Fill{Type: "gradient", Color: []string{"0066ff", "0066ff"}, Shading: 2},
		Border: []excelize.Border{
			{
				Type:  "right",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "left",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "top",
				Color: "#000000",
				Style: 1,
			},
			{
				Type:  "bottom",
				Color: "#000000",
				Style: 1,
			},
		},
		Alignment: &excelize.Alignment{
			Horizontal: "center",
			Vertical:   "center",
		},
	})
	return headStyle
}

func BodyStyle(file *excelize.File, num int) int {
	var style int

	if num%2 == 0 {
		style, _ = file.NewStyle(&excelize.Style{

			Fill: excelize.Fill{
				Type:    "gradient",
				Shading: 0,
				Color:   []string{"#f8f8f8"},
			},

			Font: &excelize.Font{
				Size: 12,
			},
			Border: []excelize.Border{
				{
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "top",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				},
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
		})
	} else {
		style, _ = file.NewStyle(&excelize.Style{

			Font: &excelize.Font{
				Size: 12,
			},
			Border: []excelize.Border{
				{
					Type:  "right",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "left",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "top",
					Color: "#000000",
					Style: 1,
				},
				{
					Type:  "bottom",
					Color: "#000000",
					Style: 1,
				},
			},
			Alignment: &excelize.Alignment{
				Horizontal: "right",
				Vertical:   "center",
			},
		})
	}

	return style
}
