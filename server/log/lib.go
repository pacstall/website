package log

import (
	"fmt"
	glog "log"
	"os"

	"github.com/fatih/color"
)

var Debug = glog.New(os.Stdout, fmt.Sprintf("%s: ", color.CyanString("DEBUG")), glog.Ldate|glog.Ltime|glog.Lshortfile)
var Info = glog.New(os.Stdout, fmt.Sprintf("%s: ", color.HiBlueString("INFO")), glog.Ldate|glog.Ltime|glog.Lshortfile)
var Error = glog.New(os.Stdout, fmt.Sprintf("%s: ", color.HiRedString("ERROR")), glog.Ldate|glog.Ltime|glog.Lshortfile)
var Warn = glog.New(os.Stdout, fmt.Sprintf("%s: ", color.YellowString("WARN")), glog.Ldate|glog.Ltime|glog.Lshortfile)

var isProduction = false

func Init(production bool) {
	color.NoColor = production
	isProduction = production
}
