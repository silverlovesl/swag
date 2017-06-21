package gen

import (
	"encoding/json"
	"github.com/easonlin404/gin-swagger/parser"
	"log"
	"os"
	"text/template"
	"time"
)

type Gen struct {
}

func New() *Gen {
	return &Gen{}
}

func (g *Gen) Build(searchDir, mainApiFile string) {
	log.Println("build")
	p := parser.New()
	p.ParseApi(searchDir, mainApiFile)
	swagger := p.GetSwagger()

	b, _ := json.MarshalIndent(swagger, "", "    ")

	err := os.MkdirAll("docs", os.ModePerm)
	die(err)
	docs, err := os.Create("docs/docs.go")
	die(err)
	defer docs.Close()

	packageTemplate.Execute(docs, struct {
		Timestamp time.Time
		Doc       string
	}{
		Timestamp: time.Now(),
		Doc:       "`" + string(b) + "`",
	})

	//TODO: print file path
	log.Printf("create docs.go at  %+v", docs)

}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var packageTemplate = template.Must(template.New("").Parse(`// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by gin-swagger at
// {{ .Timestamp }}

package docs

import (
	"github.com/easonlin404/gin-swagger/swagger"
)

var doc = {{.Doc}}

type s struct{}

func (s *s) ReadDoc() string {
	return doc
}
func init() {
	swagger.Register(swagger.Name, &s{})
}
`))
