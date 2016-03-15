package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func handleError(_e error) {
	if _e != nil {
		log.Fatal(_e)
	}
}

func TarGzWrite(_path string, tw *tar.Writer, fi os.FileInfo) {
	fr, err := os.Open(_path)
	handleError(err)
	defer fr.Close()

	h := new(tar.Header)
	h.Name = _path
	h.Size = fi.Size()
	h.Mode = int64(fi.Mode())
	h.ModTime = fi.ModTime()

	err = tw.WriteHeader(h)
	handleError(err)

	_, err = io.Copy(tw, fr)
	handleError(err)
}

func IterDirectory(dirPath string, tw *tar.Writer) {
	dir, err := os.Open(dirPath)
	handleError(err)
	defer dir.Close()
	fis, err := dir.Readdir(0)
	handleError(err)
	for _, fi := range fis {
		curPath := dirPath + "/" + fi.Name()
		if fi.IsDir() {
			//TarGzWrite( curPath, tw, fi )
			IterDirectory(curPath, tw)
		} else {
			fmt.Printf("adding... %s\n", curPath)
			TarGzWrite(curPath, tw, fi)
		}
	}
}

func TarGz(outFilePath string, inPath string) {
	// file write
	fw, err := os.Create(outFilePath)
	handleError(err)
	defer fw.Close()

	// gzip write
	gw := gzip.NewWriter(fw)
	defer gw.Close()

	// tar write
	tw := tar.NewWriter(gw)
	defer tw.Close()

	IterDirectory(inPath, tw)

	fmt.Println("tar.gz ok")
}

func main() {
	targetFilePath := "test.tar.gz"
	inputDirPath := ".test1/"
	TarGz(targetFilePath, strings.TrimRight(inputDirPath, "/"))
	fmt.Println("GOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOD")
}
