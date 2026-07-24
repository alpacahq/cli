package useragent

import (
	"regexp"
	"testing"
)

func TestBuildFormat(t *testing.T) {
	got := Build("1.2.3")
	re := regexp.MustCompile(`^APCA-CLI/1\.2\.3 \S+/\S+$`)
	if !re.MatchString(got) {
		t.Errorf("Build(%q) = %q, want format matching %s", "1.2.3", got, re.String())
	}
}

func TestBuildEmptyVersion(t *testing.T) {
	got := Build("")
	re := regexp.MustCompile(`^APCA-CLI/ \S+/\S+$`)
	if !re.MatchString(got) {
		t.Errorf("Build(\"\") = %q, want format matching %s", got, re.String())
	}
}
