package fileio

import (
	"github.com/opn-ooo/challenges/challenge-go/cipher"
	"io"
	"os"
)

type DecodedReader struct {
	io.Reader
	file *os.File
}

func OpenAndDecodeRot128File(path string) (*DecodedReader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	reader, err := cipher.NewRot128Reader(file)
	if err != nil {
		_ = file.Close()

		return nil, err
	}

	return &DecodedReader{Reader: reader, file: file}, nil
}

func (d *DecodedReader) Close() error {
	return d.file.Close()
}
