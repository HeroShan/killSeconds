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