function startDemo() {
	$.post("/demo/start", function(data) {
		console.log(data);
	})
}

function stopDemo() {

	$.post("/demo/stop", function(data) {
		console.log(data);
	})
}