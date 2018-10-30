package main
import (
	/*"os"
	"runtime"
	"crypto/tls"
	"fmt"
    "crypto/sha256"
    "io"
	"encoding/hex"
	"regexp"
	"strings"*/
)


func minify(file string){
/**/
	typef := getFileType(file)
	//str := ""
	//instr := ""
	for _, line := range getFileContentLine(file) {
		//langage agnostic
		for _, char := range line{
			switch string(char){
			case "'": 
			case "\"":
			case "`":
			}
		}
		

		switch typef {
			case "css":
			case "html":
			case "js":
			default:
			
		}
	}/**/
}

func mergeFiles(file string, files []string){
	var s string
	for _, f := range files {
		s += getFileContent(f)  + "\n"
	}
	writeToFile(s, file)
}