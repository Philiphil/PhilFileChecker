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


func mergeFiles(file string, files []string){
	var s string
	for _, f := range files {
		s += getFileContent(f)  + "\n"
	}
	writeToFile(s, file)
}