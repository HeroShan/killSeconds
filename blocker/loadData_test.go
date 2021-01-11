package blocker

import(
	"testing"
	"fmt"
	"os"
	"encoding/json"
	"log"
)



func TestNewBlockers(t *testing.T){
	b := NewBlockers(30)
	fmt.Printf("%#v\n",b)
}
type Ed struct{
	Ddd	[]Email `json:"data"`
}

type Email struct{
	Id		uint 	`gorm:"-;primary_key;AUTO_INCREMENT"`
	Email	string
	Name 	string 
}
func TestDataLoading(t *testing.T){
	b := NewBlockers(300)
	path := "./email.json"
	data := loadJson(path)
	//log.Println(string(data))
	var ldstr LoadDataInfo
	var emails Ed

	var ldfunc LoadData
	err := json.Unmarshal(data,&emails);if err != nil{
		panic(err)
	}
	//log.Printf("%#v\n",emails)
	ldfunc = func(datajson []byte) (ldd LoadDataStr){
		for k,v := range datajson{
			ldstr.Index = k
			ldstr.Value = v
			ldd.data = append(ldd.data,ldstr)
		}
		return ldd
	}
	
	b.DataLoading(ldfunc,data)
}

func loadJson(path string) []byte {
	
	file,err := os.Open(path); if err != nil{
		log.Printf("open file is fail\n")
	}
	finfo,_ := file.Stat()
	size := finfo.Size()
	jsonData := make([]byte,int(size))
	file.Read(jsonData)
	file.Close()
	return jsonData
}
