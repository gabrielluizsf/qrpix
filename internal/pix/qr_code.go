package pix

import (
	"github.com/i9si-sistemas/pix"
)

func New(copyPaste copyPaste) ([]byte, error) {
	return pix.QRCode(pix.QRCodeOptions{
		Content: copyPaste.String(),
	})
}