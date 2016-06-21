package handlers

import (
	"net/http"
)

/*
	InitializeHandlers is used to initialize the large amount of handlers that the server uses to server all of the different pages in the application.
*/
func InitializeHandlers() {
	http.HandleFunc("/login/", LoginPostHandler)
	http.HandleFunc("/", LoginPageHandler)

	http.Handle("/user/", GenericHandler{POST: AccountCreationHandler})

	http.Handle("/app/", GenericHandler{GET: AppHandler})

	http.Handle("/demo/start/", GenericHandler{PUT: SpeakerPutHandler})
	http.Handle("/demo/stop/", GenericHandler{PUT: SpeakerPutHandler})

	http.Handle("/system/", GenericHandler{GET: SystemStatusHandler})
	http.Handle("/system/controllerUpdate/", GenericHandler{GET: ControllerUpdateHandler})

	// Speaker handlers
	http.Handle("/system/speaker/", GenericHandler{PUT: SpeakerPutHandler})
	http.Handle("/system/threshold/", GenericHandler{PUT: ThresholdHandler})
	http.Handle("/system/backtofront/", GenericHandler{POST: SpeakerGetHandler})
	http.Handle("/system/addPreset/", GenericHandler{PUT: AddPresetHandler})
	http.Handle("/system/changeEQMode/", GenericHandler{PUT: ChangeEQMode})
	http.Handle("/system/addTarget/", GenericHandler{PUT: AddTargetHandler})
	http.Handle("/system/pagingRequest/", GenericHandler{POST: PagingRequestHandler})
	http.Handle("/system/createZone/", GenericHandler{PUT: CreateZoneHandler})
	http.Handle("/system/createPagingZone/", GenericHandler{PUT: CreatePagingZoneHandler})
	http.Handle("/system/scheduleTime/", GenericHandler{PUT: ScheduleTime})

	// Zone handlers
	http.Handle("/system/zone/updateVolumesZone/", GenericHandler{PUT: UpdateVolumesZone})
	http.Handle("/system/zone/threshold/", GenericHandler{PUT: ThresholdHandlerZone})
	http.Handle("/system/zone/updateAveragingZone/", GenericHandler{PUT: UpdateAveragingZone})
	http.Handle("/system/zone/updateTargetZone/", GenericHandler{PUT: UpdateSpeakerTargetZone})
	http.Handle("/system/zone/updateEqualizerZone/", GenericHandler{PUT: UpdateSpeakerEqualizerZone})
	http.Handle("/system/zone/addPreset/", GenericHandler{PUT: AddPresetHandlerZone})
	http.Handle("/system/zone/updatePagingValuesZone/", GenericHandler{PUT: UpdatePagingValuesZone})
	http.Handle("/system/zone/scheduleTimeZone/", GenericHandler{PUT: ScheduleTimeZone})
	http.Handle("/system/zone/addTarget/", GenericHandler{PUT: AddTargetHandlerZone})
	http.Handle("/system/zone/changeEQModeZone/", GenericHandler{PUT: ChangeEQModeZone})

	http.Handle("/system/zone/", GenericHandler{POST: ZonePostHandler})

	http.Handle("/css/", CssHandler)
	http.Handle("/bower_components/", BowerHandler)
	http.Handle("/components/", ComponentsHandler)
}
