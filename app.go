/*

app.go --> MAIN

Generic application with rotating log

[2019-03-15] Massimo Strambini

*/

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/maxstrambini/goutils"
)

var appName = "generic"
var appDescription = "Generic application with rotating log"
var appVersion = "1.0.1"

var start = time.Now()

func init() {
	log.Printf("init app.go (main)")
}

func main() {

	//setting up the log:
	logname := fmt.Sprintf("%s.log", appName)
	log.Printf("SETTING LOG '%s' ...\n", logname)
	/*
		f, err := os.OpenFile("fwm.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			log.Printf("error opening file: %v", err)
		}
		defer f.Close()
		////to log to file only:
		//log.SetOutput(f)
		//to log to stdout AND file:
		mw := io.MultiWriter(os.Stdout, f)
	*/
	//using max rotating log
	mw := goutils.NewMaxRotateWriter(logname, 5*1024*1024, true, 100) //filename string, maxBytes int, rotateFilesByNumber bool, maxRotatedFilesByNumber int
	log.SetOutput(mw)
	//log set up completed.

	log.Printf("===============================================")
	log.Printf("%s\n", appDescription)
	log.Printf("%s - Version %s\n", appName, appVersion)
	log.Printf("===============================================")

	log.Println("Reading config ...")
	ReadConfig() //read into var conf

	log.Println("Done with config")

}
