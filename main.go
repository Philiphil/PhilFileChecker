package main
import (
	//"os"
	//"os/exec"
	"fmt"
	//"io/ioutil"
	//"bufio"
	//"strings"
	//"net/http"
	//"strconv"
	//"time"
	//"runtime"
	//"encoding/json"
	//"html"
)

type OrderType int;

const(
	OT_delete OrderType = 1
	OT_git_push OrderType = 2
	OT_ftp_share OrderType = 3
	OT_execute	OrderType=4
)

type Order struct{
	Type OrderType
	Param []string
}

type Trigger struct{
	File string
	Orders []Order
}

var(
	_GIT_IGNORE=true;
)
func init() {
}

func main(){
	fmt.Println(1)
	/*
	r:=explore("/home/tsorriaux/src/script/shortcut_bash_script/")
	fmt.Println(r)*/

}

func cycle(T Trigger){
	old_state := explore(T.File)
	for ;; {
		new_state := explore(T.File)
		diff := getDiff(old_state, new_state)
		if len(diff) > 0{

		}
	}
}

func executeTrigger(T Trigger){
	for _, order := range T.Orders {
		switch order.Type {
			case OT_delete:
			
		}
	}
}