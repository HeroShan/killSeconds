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

	var ldfunc LoadData
	var emails Ed
	//log.Printf("%#v\n",emails)
	ldfunc = func(wqq []byte) (ldd LoadDataStr){
		
		err := json.Unmarshal(wqq,&emails);if err != nil{
			panic(err)
		}
		for k,v := range emails.Ddd{
			ldstr.Index = k
			ldstr.Value = v
			ldd.data = append(ldd.data,ldstr)
		}
		return ldd
	}
	
	b.DataLoading(ldfunc,data)
	log.Printf("%#v",len(b.Bmap))
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
