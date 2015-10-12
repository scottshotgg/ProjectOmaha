package handlers

import (
	"net/http"
	"omaha/util"
)

var CssHandler http.Handler = http.StripPrefix("/css/", http.FileServer(http.Dir(util.GetOmahaPath()+"/templates/css")))
var BowerHandler http.Handler = http.StripPrefix("/bower_components/", http.FileServer(http.Dir(util.GetOmahaPath()+"/templates/bower_components")))
var ComponentsHandler http.Handler = http.StripPrefix("/components/", http.FileServer(http.Dir(util.GetOmahaPath()+"/templates/components")))
