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
	//b := NewBlockers(300)
	path := "./email.json"
	data := loadJson(path)
	log.Println(string(data))
	//var ldstr LoadDataStr
	var emails Ed
	err := json.Unmarshal(data,&emails);if err != nil{
		panic(err)
	}
	log.Printf("%#v\n",emails)
	// for k,v := range datajson{
	// 	ldstr.Index = k
	// 	ldstr.Value = v
	// 	log.Println(ldstr)
	// }
	// b.DataLoading()
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
