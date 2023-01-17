package utils

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func BeeFuncMap() template.FuncMap {
	return template.FuncMap{
		"trim":       strings.TrimSpace,
		"bold":       Bold,
		"headline":   MagentaBold,
		"foldername": RedBold,
		"endline":    EndLine,
		"tmpltostr":  TmplToString,
	}
}

func EndLine() string {
	return "\n"
}

func RedBold(message string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", Bold(message))
}

// MagentaBold returns a magenta Bold string
func MagentaBold(message string) string {
	return fmt.Sprintf("\x1b[35m%s\x1b[0m", Bold(message))
}

func Bold(message string) string {
	return fmt.Sprintf("\x1b[1m%s\x1b[0m", message)
}

func Tmpl(text string, data interface{}) {
	t := template.New("Usage").Funcs(BeeFuncMap())
	template.Must(t.Parse(text))

	err := t.Execute(os.Stderr, data)
	if err != nil {
		//beeLogger.Log.Error(err.Error())
	}
}

func TmplToString(tmpl string, data interface{}) string {
	t := template.New("tmpl").Funcs(BeeFuncMap())
	template.Must(t.Parse(tmpl))

	var doc bytes.Buffer
	err := t.Execute(&doc, data)
	MustCheck(err)

	return doc.String()
}

func MustCheck(err error) {
	if err != nil {
		panic(err)
	}
}
