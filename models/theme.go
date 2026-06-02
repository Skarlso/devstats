package models

type Theme struct {
	BgTop    string
	BgBottom string
	Border   string
	Divider  string
	Title    string
	Label    string
	Value    string
	Radius   int
}

const DefaultTheme = "dark"

var themes = map[string]Theme{
	"dark": {
		BgTop:    "#1b2430",
		BgBottom: "#0d1117",
		Border:   "#30363d",
		Divider:  "#21262d",
		Title:    "#58a6ff",
		Label:    "#8b949e",
		Value:    "#e6edf3",
		Radius:   12,
	},
	"light": {
		BgTop:    "#ffffff",
		BgBottom: "#f6f8fa",
		Border:   "#d0d7de",
		Divider:  "#d8dee4",
		Title:    "#0969da",
		Label:    "#57606a",
		Value:    "#1f2328",
		Radius:   12,
	},
	"cncf": {
		BgTop:    "#241b35",
		BgBottom: "#15101f",
		Border:   "#3b2d54",
		Divider:  "#2c2140",
		Title:    "#bd93f9",
		Label:    "#9785b0",
		Value:    "#f4f0ff",
		Radius:   12,
	},
}

func ThemeByName(name string) Theme {
	if t, ok := themes[name]; ok {
		return t
	}
	return themes[DefaultTheme]
}
