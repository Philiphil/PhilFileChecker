package main
import (
	"reflect"
)


func mergeFiles(file string, files []string){
	var s string
	for _, f := range files {
		s += getFileContent(f)  + "\n"
	}
	writeToFile(s, file)
}


func HasElem(s interface{}, elem interface{}) bool {
    arrV := reflect.ValueOf(s)

    if arrV.Kind() == reflect.Slice {
        for i := 0; i < arrV.Len(); i++ {

            if arrV.Index(i).Interface() == elem {
                return true
            }
        }
    }

    return false
}