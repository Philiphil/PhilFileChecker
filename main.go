package main
import (
	"os"
	"fmt"
	"flag"
	"io/ioutil"
	"time"
	"runtime"
	"encoding/json"
)

type OrderType string;

const(
	OT_delete OrderType = "delete"
	OT_git_push OrderType = "git_push"
	OT_ftp_sync OrderType = "ftp_sync"
	OT_execute	OrderType= "execute"
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
	_SILENT=false;
)
func init() {
	_GIT_IGNORE = *flag.Bool("git-ignore", true, ".git folder are ignored by default")
	_SILENT = *flag.Bool("silent", false, "silent mode")
	flag.Parse()

}

func main(){
	var Triggers []Trigger
	file, err := ioutil.ReadFile("./pfc.json")
	if err != nil {
    	fmt.Println("pfc.json not found")
    }
	json.Unmarshal(file, &Triggers)
	for _,Trigger := range Triggers {
		go cycle(Trigger)
	}
	for {
		select{
		default:		
			time.Sleep(250 * time.Millisecond)
			runtime.Gosched()
		}
	}
}

func cycle(T Trigger){
	if(!_SILENT) {
		fmt.Println(T)
	}
	old_state := explore(T.File)
	for {
		runtime.Gosched()
		new_state := explore(T.File)
		//fmt.Print(new_state)
		diff := getDiff(old_state, new_state)
		old_state=new_state
		if len(diff) > 0{
			executeTrigger(T)
		}
        time.Sleep(500 * time.Millisecond)
		runtime.Gosched()
	}
}

func executeTrigger(T Trigger){
	if(!_SILENT) {
		fmt.Println("Triggered", T)
	}
	for _, order := range T.Orders {
		runtime.Gosched()
		switch order.Type {
			case OT_delete : 
				orderDelete(order)
		}
	}
}

func orderDelete(O Order){
	if ok, _ := isDirectory(O.Param[0]); ok{
		os.RemoveAll(O.Param[0])
	}else{
		os.Remove(O.Param[0])
	}
}