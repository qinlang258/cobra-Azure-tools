package exportexcel

import (
	"cobra-Azure-tools/database"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ExportVmCpu() {
	var nodes []database.NodeInfo
	db := database.ContentMysql()

	db.Model(&database.NodeInfo{}).
		Find(&nodes)

	file := excelize.NewFile()
	sheetName := "Sheet1"

	// 设置表头
	file.SetCellValue(sheetName, "A1", "ID")
	file.SetCellValue(sheetName, "B1", "服务器IP")
	file.SetCellValue(sheetName, "C1", "服务器时间")
	file.SetCellValue(sheetName, "D1", "服务器状态")
	file.SetCellValue(sheetName, "E1", "运行时间")
	file.SetCellValue(sheetName, "F1", "有多少个用户登录")
	file.SetCellValue(sheetName, "G1", "过去1分钟平均负载")
	file.SetCellValue(sheetName, "H1", "过去5分钟平均负载")
	file.SetCellValue(sheetName, "I1", "过去15分钟平均负载")

	for i, node := range nodes {
		row := i + 2
		file.SetCellValue(sheetName, "A"+strconv.Itoa(row), node.ID)
		file.SetCellValue(sheetName, "B"+strconv.Itoa(row), node.IP)
		file.SetCellValue(sheetName, "C"+strconv.Itoa(row), node.Datetime)
		file.SetCellValue(sheetName, "D"+strconv.Itoa(row), node.Status)
		file.SetCellValue(sheetName, "E"+strconv.Itoa(row), node.StartedTime)
		file.SetCellValue(sheetName, "F"+strconv.Itoa(row), node.UserNumber)
		file.SetCellValue(sheetName, "G"+strconv.Itoa(row), node.LoadAverage1m)
		file.SetCellValue(sheetName, "H"+strconv.Itoa(row), node.LoadAverage5m)
		file.SetCellValue(sheetName, "I"+strconv.Itoa(row), node.LoadAverage15m)

	}
	file.SaveAs("Node的CPU运行情况.xlsx")

}
