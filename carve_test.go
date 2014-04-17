package carve_test

import (
	carve "github.com/scottmotte/carve"
	"testing"
)

func TestConvertToPngs(t *testing.T) {
	base, err := carve.Download("http://www.unc.edu/~stanlele/task03.doc")
	if err != nil {
		t.Errorf("Error", err)
	}

	path, err := carve.ConvertToPdf(base)
	if err != nil {
		t.Errorf("Error", err)
	}

	pngs, _ := carve.ConvertToPngs(path)
	if len(pngs) == 0 {
		t.Errorf("Pngs were blank")
	}
}
