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
	http.Handle("/system/controllerUpdate/", GenericHandler{GET: ControllerUpdateHandler})
	http.Handle("/system/speaker/", GenericHandler{PUT: SpeakerPutHandler})
	http.Handle("/system/backtofront/", GenericHandler{POST: SpeakerGetHandler})
	http.Handle("/system/addPreset/", GenericHandler{PUT: AddPresetHandler})
	http.Handle("/system/addTarget/", GenericHandler{PUT: AddTargetHandler})
	http.Handle("/system/pagingRequest/", GenericHandler{POST: PagingRequestHandler})
	http.Handle("/system/createZone/", GenericHandler{PUT: CreateZoneHandler})
	http.Handle("/system/createPagingZone/", GenericHandler{PUT: CreatePagingZoneHandler})

	// I'll consolidate these later, easier to forge my own path for now
	http.Handle("/system/zone/", GenericHandler{POST: ZonePostHandler})


	// file handlers
	http.Handle("/css/", CssHandler)
	http.Handle("/bower_components/", BowerHandler)
	http.Handle("/components/", ComponentsHandler)
}
