package svg

import (
	"strings"
	"testing"

	"github.com/Skarlso/devstats/models"
)

func TestGenerateSVG(t *testing.T) {
	out := GenerateSVG(models.CardData{
		Score:      1234567,
		PRs:        42,
		Issues:     7,
		TitleColor: "#0086FF",
		TextColor:  "#555555",
		Radius:     10,
	})

	wants := []string{
		"<svg",
		"CNCF DevStats",
		"#0086FF",
		"#555555",
		"1,234,567",
	}
	for _, want := range wants {
		if !strings.Contains(out, want) {
			t.Errorf("output missing %q\n%s", want, out)
		}
	}
}

func TestGenerateSVGFormatsLargeNumbers(t *testing.T) {
	out := GenerateSVG(models.CardData{Score: 1000000})
	if !strings.Contains(out, "1,000,000") {
		t.Errorf("score not comma-formatted\n%s", out)
	}
}
