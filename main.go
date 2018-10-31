package main
import (
	"fmt"
	"flag"
	"io/ioutil"
	"time"
	"runtime"
	"encoding/json"
)

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
	//minify("/home/tsorriaux/src/script/PDT/test.css")
	//return;

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
		fmt.Println(time.Now().Format("3h4:5"),"Triggered", T)
	}
	for _, order := range T.Orders {
		runtime.Gosched()
		switch order.Type {
			case OT_delete : orderDelete(order)
			case OT_ftp_sync : orderFtpSync(order)
			case OT_minify : orderMinify(order)
		}
	}
}