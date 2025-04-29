package pix

import (
	"github.com/i9si-sistemas/pix"
)


type Options pix.Options

func Generate(
	options Options,
) (copyPaste, error) {
	cp, err := pix.New(pix.Options(options))
	if err != nil {
		return "", err
	}
	return copyPaste(cp), nil
}
