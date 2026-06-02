package svg

import (
	"strings"
	"testing"

	"github.com/Skarlso/devstats/models"
)

func TestGenerateSVG(t *testing.T) {
	dark := models.ThemeByName("dark")
	out := GenerateSVG(models.CardData{
		Score:  1234567,
		PRs:    42,
		Issues: 7,
		Theme:  dark,
	})

	wants := []string{
		"<svg",
		"CNCF DevStats",
		dark.Title,
		dark.Value,
		dark.BgBottom,
		"1,234,567",
	}
	for _, want := range wants {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n%s", want, out)
		}
	}
}

func TestGenerateSVGUsesTheme(t *testing.T) {
	light := models.ThemeByName("light")
	out := GenerateSVG(models.CardData{Score: 1, Theme: light})

	if !strings.Contains(out, light.BgTop) {
		t.Errorf("light background %q not applied\n%s", light.BgTop, out)
	}
	dark := models.ThemeByName("dark")
	if strings.Contains(out, dark.BgTop) {
		t.Errorf("dark background %q leaked into light card", dark.BgTop)
	}
}

func TestGenerateSVGFormatsLargeNumbers(t *testing.T) {
	out := GenerateSVG(models.CardData{Score: 1000000, Theme: models.ThemeByName("dark")})
	if !strings.Contains(out, "1,000,000") {
		t.Errorf("score not comma-formatted\n%s", out)
	}
}
