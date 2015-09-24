package handlers

import (
	"net/http"
	"omaha/util"
)

var CssHandler http.Handler = http.StripPrefix("/css/", http.FileServer(http.Dir(util.GetOmahaPath()+"/css")))
