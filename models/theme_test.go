package models

import "testing"

func TestThemeByName(t *testing.T) {
	tests := []struct {
		name string
		want Theme
	}{
		{"dark", themes["dark"]},
		{"light", themes["light"]},
		{"cncf", themes["cncf"]},
		{"", themes[DefaultTheme]},
		{"nonsense", themes[DefaultTheme]},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ThemeByName(tt.name); got != tt.want {
				t.Errorf("ThemeByName(%q) = %+v, want %+v", tt.name, got, tt.want)
			}
		})
	}
}
