package models

//template data holds the data to be send from handler to template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSTToken  string
	Flash     string
	Warning   string
	Error     string
}
