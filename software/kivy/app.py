from kivy.app import App
from kivy.uix.button import Button
from kivy.uix.widget import Widget
from kivy.properties import ObjectProperty
import json, urllib2

class MyWidget(Widget):
    volumeBtn = ObjectProperty(None)
    volumeSlider = ObjectProperty(None)

    def setVolume(self, instance):
		reqData = {
			'updatedAttributes': ['volume'],
			'attributeValues': {
				'volume': int(self.volumeSlider.value)
			},
			'speaker': 107
			}
		data = json.dumps(reqData)
		opener = urllib2.build_opener(urllib2.HTTPHandler)
		request = urllib2.Request("http://localhost:8080/system/speaker/", data=data)
		request.add_header('Content-Type', 'application/json')
		request.get_method = lambda: 'PUT'
		url = opener.open(request)

class OmahaApp(App):

    def build(self):
        widget = MyWidget()
        widget.volumeBtn.bind(on_press=widget.setVolume)
        return widget

if __name__ == '__main__':
    OmahaApp().run()