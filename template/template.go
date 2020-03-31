package template

import (
    "bytes"
    "text/template"
)

func Parse(goTemplate string, data interface{}) (string, error ){
    t := template.Must(template.New("").Parse(goTemplate))
    var buf bytes.Buffer
    err := t.Execute(&buf, data)
    return buf.String(), err
}