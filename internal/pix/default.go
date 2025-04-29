package pix

var defaultOpts = Options{
	Name: "Gabriel Luiz",
	Key:  "7a067b11-bce7-406f-af8a-2dcf82c429d6",
	City: "Caruaru",
}

func Default() (copyPaste, error) {
	return Generate(defaultOpts)
}
