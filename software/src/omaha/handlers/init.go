package handlers

import (
	"net/http"
)

func InitializeHandlers() {
	http.HandleFunc("/login/", LoginPostHandler)
	http.HandleFunc("/", LoginPageHandler)

	http.Handle("/user/", GenericHandler{POST: AccountCreationHandler})

	http.Handle("/app/", GenericHandler{GET: AppHandler})

	http.Handle("/demo/start/", GenericHandler{PUT: SpeakerPutHandler})
	http.Handle("/demo/stop/", GenericHandler{PUT: SpeakerPutHandler})

	http.Handle("/system/", GenericHandler{GET: SystemStatusHandler})
	http.Handle("/system/speaker/", GenericHandler{PUT: SpeakerPutHandler})
	http.Handle("/system/backtofront/", GenericHandler{GET: SpeakerGetHandler})

	http.Handle("/system/zone/", GenericHandler{POST: ZonePostHandler})

	// file handlers
	http.Handle("/css/", CssHandler)
	http.Handle("/bower_components/", BowerHandler)
	http.Handle("/components/", ComponentsHandler)
}
