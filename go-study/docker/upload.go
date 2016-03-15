package main

import (
	"archive/tar"
	"bytes"
	"io"
	"log"
	"os"

	docker "github.com/fsouza/go-dockerclient"
)

func uploadTar() {
	client, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.20")

	// Start create .tar file
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	var tarFile string = "./docker.tar.gz"
	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	fileInfo, err := os.Open(tarFile)
	if err != nil {
		log.Fatalln(err)
	}
	stat, _ := fileInfo.Stat()
	hdr := &tar.Header{
		Name: tarFile,
		Mode: int64(stat.Mode()),
		Size: stat.Size(),
	}
	if err := tw.WriteHeader(hdr); err != nil {
		log.Fatalln(err)
	}
	//              if _, err := tw.Write([]byte(file.Body)); err != nil {
	//                      log.Fatalln(err)
	//              }
	_, err = io.Copy(tw, fileInfo)

	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
	// End .tar file

	// Start upload .tar
	uploadOption := docker.UploadToContainerOptions{
		InputStream:          buf,
		Path:                 "/etc/",
		NoOverwriteDirNonDir: true,
	}

	err = client.UploadToContainer("omega-slave", uploadOption)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Good")

	// End upload .tar
}

func main() {
	client, err := docker.NewVersionedClient("unix:///var/run/docker.sock", "1.20")

	// Start create .tar file
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	// Add some files to the archive.
	var files = []struct {
		Name, Body string
	}{
		{"file1.txt", "This archive contains some text files."},
		{"file2.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
	}

	for _, file := range files {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}
	// End .tar file

	// Start upload .tar
	uploadOption := docker.UploadToContainerOptions{
		InputStream:          buf,
		Path:                 "/",
		NoOverwriteDirNonDir: true,
	}

	err = client.UploadToContainer("omega-slave", uploadOption)
	if err != nil {
		log.Panicln(err)
	}

	log.Println("Good")

	// End upload .tar
}
