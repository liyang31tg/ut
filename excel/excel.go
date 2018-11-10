package excel

import (
	"github.com/360EntSecGroup-Skylar/excelize"
	"strconv"
	"strings"
)

type XlsxUtil struct {
	Data [][]interface{}
	Sheet string
	Path string
}

func NewXlsxUtil(Data [][]interface{}) *XlsxUtil {
	return  &XlsxUtil{
		Data:Data,
	}
}
var head = [26]string{"A","B","C","D","E","F","G","H","I","J","K","L","M","N","O","P","Q","R","S","T","U","V","W","X","Y","Z"}
func (this *XlsxUtil)head(index int) string{
	if index > 26 *26 {
		panic("index out of 26 * 26")
	}
	if index <1 {
		panic("index out of 1")
	}
	if index <= 26 {
		return head[index-1]
	}
	if index >26 && index < 26 * 26 {
		ib := index / 26
		iy := index % 26
		var start ,end string
		if ib != 0{
			start = head[ib-1]
		}
		end = head[iy]
		return start + end
	}
	return  ""
}

func (this *XlsxUtil)Save() {
	xlsx := excelize.NewFile()
	if this.Sheet == "" {
		this.Sheet =  "Sheet1"
	}else{
		index :=xlsx.NewSheet(this.Sheet)
		xlsx.SetActiveSheet(index)
	}

	if this.Path == "" {
		this.Path =  "./Book.xlsx"
	}
	if !strings.HasSuffix(this.Path,".xlsx") {
		panic("this path is must be suffix xlsx")
	}
	for i :=0;i<len(this.Data);i++{
		for j:=1;j<=len(this.Data[i]);j++{
			cellData := this.Data[i][j-1]
			xlsx.SetCellValue(this.Sheet,this.head(j)+strconv.Itoa(i+1),cellData)
		}
	}
	err := xlsx.SaveAs(this.Path)
	if err != nil {
		panic(err)
	}
}

