package fileio_test

import (
	"go-tamboon/internal/infrastructure/fileio"
	"io"
	"log"
	"testing"
)

func TestOpenAndDecodeRot128File(t *testing.T) {
	t.Run("Should simple open rot128 file", func(t *testing.T) {
		reader, err := fileio.OpenAndDecodeRot128File("../../../data/fng.1000.csv.rot128")
		if err != nil {
			t.Fatal(err)
		}

		defer reader.Close()

		data, err := io.ReadAll(reader.Reader)
		if err != nil {
			t.Fatal(err)
		}

		log.Print(string(data))
	})
}
