package controller

//go:embed public/index.html
var IndexHtml []byte

/*
//go:embed files/bpmn/diagram_1.bpmn
var WorkflowBpmnFile []byte
*/

/*
func handleIndex(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	writer.Write(IndexHtml)
}

func handleBpmn(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/xml")
	writer.Write(WorkflowBpmnFile)
}
*/

/*
	http.HandleFunc("/public/index.html", handleIndex)
	http.HandleFunc("/files/bpmn/diagram_1.bpmn", handleBpmn)
	http.ListenAndServe(":8080", nil)
*/
