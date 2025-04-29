package pix

import (
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestDefaultPix(t *testing.T) {
	cp, err := Default()
	assert.NoError(t, err)
	result := cp.String()
	expected := "00020126620014BR.GOV.BCB.PIX01367a067b11-bce7-406f-af8a-2dcf82c429d6020052040000530398654040.005802BR5912Gabriel Luiz6007Caruaru62410503***50300017BR.GOV.BCB.BRCODE01051.0.0630434EF"
	assert.Equal(t, result, expected)
}
