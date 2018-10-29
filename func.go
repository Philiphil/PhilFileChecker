package main
import (
	"os"
	"runtime"
	"crypto/tls"
	"fmt"
    "crypto/sha256"
    "io"
	"encoding/hex"
//	"regexp"
//	"strings"
	"github.com/dutchcoders/goftp"
)

func orderDelete(O Order){
	for _,e := range O.Params{
		runtime.Gosched()
		if ok, _ := isDirectory(e); ok{
			os.RemoveAll(e)
		}else{
			os.Remove(e)
		}
	}
}

func orderFtpSync(O Order){
	/*
		prevoir lecture params pour prot/host/port/login/pass/key location
		prevoir de pas recreer /local/dir  dans /dist/dir qui ferait /dist/dir/dir

	
	host := O.Params[0]
	port := O.Params[1]
	user := O.Params[2]
	pass := O.Params[3]
	folder := O.Params[4]
*/


	var ftp *goftp.FTP

    var err error
    // For debug messages: goftp.ConnectDbg("ftp.server.com:21")
    if ftp, err = goftp.Connect("ftp.server.com:21"); err != nil {
        panic(err)
    }

    defer ftp.Close()

    config := tls.Config{
        InsecureSkipVerify: true,
        ClientAuth:         tls.RequestClientCert,
    }

    if err = ftp.AuthTLS(&config); err != nil {
        panic(err)
	}

    // Username / password authentication
    if err = ftp.Login("username", "password"); err != nil {
        panic(err)
	}
	if err = ftp.Cwd("/"); err != nil {
        panic(err)
	}
	 // Upload a file
	 var file *os.File
	 if file, err = os.Open("/tmp/test.txt"); err != nil {
		 panic(err)
	 }
 
	 if err := ftp.Stor("/test.txt", file); err != nil {
		 panic(err)
	 }
 
	 // Download each file into local memory, and calculate it's sha256 hash
	 err = ftp.Walk("/", func(path string, info os.FileMode, err error) error {
		 _, err = ftp.Retr(path, func(r io.Reader) error {
			 var hasher = sha256.New()
			 if _, err = io.Copy(hasher, r); err != nil {
				 return err
			 }
 
			 hash := fmt.Sprintf("%s %x", path, hex.EncodeToString(hasher.Sum(nil)))
			 fmt.Println(hash)
 
			 return err
		 })
 
		 return nil
	 })
	
}

func orderMinify(O Order){
	for _,e := range O.Params{
		runtime.Gosched()
		if boolean, r := isDirectory( e ); (boolean && r == nil){
			//folder
			files := explore(e)
			for file,_ := range files{
				if b_f, r_f := isDirectory( file ); (!b_f && r_f == nil){
					minify(file)
				}
			}
        }else if (r == nil){
			minify(e)
        }else{
			fmt.Println(e, "does not exists")
		}
	}
}

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
