package main
import (
	"os"
	//"os/exec"
	"fmt"
	"flag"
	"io/ioutil"
	//"bufio"
	//"strings"
	//"net/http"
	//"strconv"
	//"time"
	//"runtime"
	"encoding/json"
	//"html"
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
)
func init() {
	_GIT_IGNORE = *flag.Bool("git-ignore", true, ".git folder are ignored by default")
	//wordPtr := flag.String("word", "foo", "a string")
	//var svar string
	//flag.StringVar(&svar, "svar", "bar", "a string var")
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
	for ;;{
	}

}

func cycle(T Trigger){
	fmt.Println(T)
	old_state := explore(T.File)
	for ;; {
		new_state := explore(T.File)
		diff := getDiff(old_state, new_state)
		old_state=new_state
		if len(diff) > 0{
			executeTrigger(T)
		}
	}
}

func executeTrigger(T Trigger){
	fmt.Println("Triggered", T)
	for _, order := range T.Orders {
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