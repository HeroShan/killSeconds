package blocker

import(
	"testing"
	"fmt"
)

func TestLoadJson(t *testing.T){
	path := "./email.json"
	LoadJson(path)

}

func TestNewBlockers(t *testing.T){
	b := NewBlockers(30)
	fmt.Printf("%#v\n",b)
}	

func TestLoadData(t *testing.T){
	b := NewBlockers(300)
	path := "./email.json"
	data := LoadJson(path)
	b.LoadData(data)
}

