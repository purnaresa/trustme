package keygen

import (
	"archive/zip"
	"bytes"
	"io"
	"log"
	"os"

	"github.com/purnaresa/bulwark/rsa-encryption"
)

func GenerateKeyPair() (output []byte, err error) {
	_, _, privateKey, publicKey, err := rsa.GenerateKeyPairInPEM()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	privateKeyFile, err := w.Create("private-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	_, err = privateKeyFile.Write(privateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKeyFile, err := w.Create("public-key.pem")
	if err != nil {
		log.Fatal(err)
	}
	_, err = publicKeyFile.Write(publicKey)
	if err != nil {
		log.Fatal(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Create("key-pair.zip")
	if err != nil {
		log.Fatal(err)
	}
	// buf.
	if _, err = io.Copy(file, buf); err != nil {
		log.Fatal(err)
	}
	file.Close()

	output = buf.Bytes()
	return

}
