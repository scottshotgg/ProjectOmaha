function startDemo() {
	$.post("/demo/write", function(data) {
		console.log(data);
	})
}