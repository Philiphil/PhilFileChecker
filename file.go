package main
import (
	"os"
	//"os/exec"
	//"fmt"
	"io/ioutil"
	//"bufio"
	//"strings"
	//"net/http"
	//"strconv"
	"time"
	//"runtime"
	//"encoding/json"
	//"html"
)

func isDirectory(path string) (bool, error) {
    if  fileInfo, err := os.Stat(path); err == nil{
    	return fileInfo.IsDir(),nil
    }else{
    	 return false, err
    }
}
   
func directorySeparator(path string) (newpath string) {
	newpath =  path
	if   path[len(path)-1:] != string (os.PathSeparator){
		newpath += string (os.PathSeparator) 
	}
	return 
}

func explore(monitoredfilelocation string)(contents map[string]string){
	contents = map[string]string{}
 	files, _ := ioutil.ReadDir(monitoredfilelocation)
   	for _, f := range files {
	   	if f.Name() == ".git" && _GIT_IGNORE {
			continue
	   	}
   		str := monitoredfilelocation + f.Name()
        if boolean, _ := isDirectory( str ); boolean{
			str = directorySeparator(str)
			contents[str] =  f.ModTime().Format(time.RFC3339) 
			for k, v := range explore(str) {
				contents[k] = v
			}
        }else{
			contents[str] =  f.ModTime().Format(time.RFC3339) 
        }
	}
	return
}

func getDiff(oldstate  map[string]string, newstate  map[string]string) (diff map[string]string){
	diff = map[string]string{}
	for key, val := range newstate {
		if o_val, ok :=oldstate[key];ok {
			delete(oldstate, key)
			if o_val != val{
				diff[key] = val
			}
		}else{
			diff[key] = val
		}
	}
	for key, val := range oldstate {
		diff[key] = val	
	}
	return
}