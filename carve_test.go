package carve_test

import (
	carve "github.com/scottmotte/carve"
	"testing"
)

func TestConvert(t *testing.T) {
	pngs, err := carve.Convert("http://www.bramstoker.org/pdf/stories/03guest/01guest.pdf")
	if err != nil {
		t.Errorf("Error", err)
	}

	if len(pngs) <= 10 {
		t.Errorf("Pngs were blank or not enough")
	}
}
