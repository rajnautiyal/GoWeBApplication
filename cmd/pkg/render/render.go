package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parseTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parseTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("this is some error", err)
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	_, inMap := tc[t]
	log.Println("this is my home template", t)
	if !inMap {
		log.Println("creating a template and storing in the logs")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}

	} else {
		log.Println("using the cach function")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println("this is some error", err)
	}

}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil
}
