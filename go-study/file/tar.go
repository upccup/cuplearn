package main

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func writeFileToTar(writer *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		return err
	}

	header := &tar.Header{
		Name:    sanitizedName(filename),
		Mode:    int64(stat.Mode()),
		Uid:     os.Getuid(),
		Gid:     os.Getgid(),
		Size:    stat.Size(),
		ModTime: stat.ModTime(),
	}
	if err = writer.WriteHeader(header); err != nil {
		return err
	}

	_, err = io.Copy(writer, file)
	return err
}

func sanitizedName(fileName string) string {
	if len(fileName) > 1 && fileName[1] == ':' && runtime.GOOS == "windows" {
		fileName = fileName[2:]
	}

	fileName = filepath.ToSlash(fileName)
	fileName = strings.TrimLeft(fileName, "/.")

	return strings.Replace(fileName, "../", "", -1)
}

func main() {
	var tarFileName string = "test.tar"
	var files []string = []string{"file1.txt", "file2.txt"}

	file, err := os.Create(tarFileName)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	var fileWriter io.WriteCloser = file

	if strings.HasSuffix(tarFileName, ".gz") {
		fileWriter = gzip.NewWriter(file)
		defer fileWriter.Close()
	}

	writer := tar.NewWriter(fileWriter)
	defer writer.Close()

	for _, name := range files {
		if err := writeFileToTar(writer, name); err != nil {
			log.Panicln(err)
		}
	}

	return
}
