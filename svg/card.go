package svg

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/dustin/go-humanize"
	"github.com/Skarlso/devstats/models"
)

//go:embed card.svg.tmpl
var cardTemplate string

var tmpl = template.Must(template.New("svg").Funcs(template.FuncMap{
	"formatNumber": func(n int) string {
		return humanize.Comma(int64(n))
	},
}).Parse(cardTemplate))

func GenerateSVG(data models.CardData) string {
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "<svg><text>Render error</text></svg>"
	}

	return buf.String()
}
