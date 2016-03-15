package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
)

func main() {
	var tarFile, filePath string
	tarFile = "/data/test.tar"
	filePath = "./"

	f, err := os.Open(tarFile)
	if err != nil {
		log.Panicln(err)
	}
	defer f.Close()

	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Panicln(err)
		}

		if hdr.FileInfo().IsDir() {
			os.Mkdir(filePath+string(os.PathSeparator)+hdr.Name, hdr.FileInfo().Mode())
		} else {
			func() {
				fw, err := os.OpenFile(filePath+string(os.PathSeparator)+hdr.Name, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, hdr.FileInfo().Mode())
				if err != nil {
					log.Panicln(err)
				}
				defer fw.Close()
				_, err = io.Copy(fw, tr)
				if err != nil {
					log.Panicln(err)
				}
			}()
		}
	}
}
