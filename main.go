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
	Params []string
}

type Trigger struct{
	Files []string
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
		return
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

	old_states :=make([]map[string]string, len(T.Files))
	for i,e := range T.Files{
		old_states[i] = explore(e)
	}

	for {
		for i,e := range T.Files{
			runtime.Gosched()
			new_state := explore(e)
			diff := getDiff(old_states[i], new_state)
			old_states[i]=new_state
			if len(diff) > 0{
				executeTrigger(T)
			}
			time.Sleep(500 * time.Millisecond)
			runtime.Gosched()
		}
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
	for _,e := range O.Params{
		if ok, _ := isDirectory(e); ok{
			os.RemoveAll(e)
		}else{
			os.Remove(e)
		}
	}
}