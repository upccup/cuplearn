package main

import (
	"archive/tar"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"time"

	containertty "github.com/Dataman-Cloud/crane/src/plugins/tty"
	log "github.com/Sirupsen/logrus"
	docker "github.com/fsouza/go-dockerclient"
	"github.com/gin-gonic/gin"
)

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()
	r.LoadHTMLGlob("tty/*.html")
	r.Static("/static", "tty")
	r.GET("/terminal", TerminalWeb)
	r.GET("/v1/registryauth/docker.tar.gz", GetDockerCredentialsFile)
	r.GET("/v1/containers/:ip", ListContainers)
	r.GET("/v1/containers/:ip/:name", InspectContainer)
	r.GET("/v1/containers/:ip/:name/terminal", ContainerTerminal)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8085")
}

func TerminalWeb(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"message":   "success",
		"timestamp": time.Now().Unix(),
	})
}

//func DockerProxy() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		dockerIP := ctx.Request.Header.Get("docker-ip")
//		if dockerIP == "" {
//			return
//		}
//
//		dockerPort := ctx.Request.Header.Get("")
//	}
//}

func ListContainers(c *gin.Context) {
	ip := c.Param("ip")

	client, err := docker.NewVersionedClient(ip+":2357", "1.21")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"data": err.Error()})
		return
	}

	containers, err := client.ListContainers(docker.ListContainersOptions{All: true})
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": containers})
	return
}

func InspectContainer(c *gin.Context) {
	ip := c.Param("ip")
	cName := c.Param("name")

	client, err := docker.NewVersionedClient(ip+":2357", "1.21")
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"data": err.Error()})
		return
	}

	cInfo, err := client.InspectContainer(cName)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"data": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": cInfo})
	return
}

func ContainerTerminal(ctx *gin.Context) {
	ip := ctx.Param("ip")
	cName := ctx.Param("name")

	req := ctx.Request
	conn, err := containertty.Upgrader.Upgrade(ctx.Writer, req, nil)
	if err != nil {
		log.Error("Upgrade websocket connect got error: ", err)
		return
	}

	_, stream, err := conn.ReadMessage()
	if err != nil {
		log.Error("Get websocket init message got error: ", err)
		return
	}
	log.Info("Init message: ", string(stream))

	cmd := exec.Command("docker", "-H", "tcp://"+ip+":2357", "exec", "-it", cName, "sh")
	client, err := containertty.New(cmd, conn, req, containertty.DefaultOptions)
	if err != nil {
		log.Error("Create tty client got error: ", err)
		return
	}

	client.HandleClient()
	return
}

func GetDockerCredentialsFile(ctx *gin.Context) {
	ctx.Request.Header.Set("Content-Type", "application/x-compressed-tar")
	ctx.Request.Header.Set("Content-Disposition",
		fmt.Sprintf("attachment; filename=\"%s\"", "docker.tar.gz"))

	if err := WriteDockerCredentialsFile(ctx.Writer); err != nil {
		log.Errorf("write docker credentialsfile failed. Error: %s", err.Error())
		return
	}

	return
}

type RegistryAuthInfo struct {
	Auths map[string]RegistryUserInfo `json:"auths"`
}

type RegistryUserInfo struct {
	Auth  string `json:"auth"`
	Email string `json:"email"`
}

var DefaultDockerCredentialsInfo RegistryAuthInfo

func init() {
	DefaultDockerCredentialsInfo = RegistryAuthInfo{
		Auths: map[string]RegistryUserInfo{
			"dockerhub.jd.com": RegistryUserInfo{
				Auth:  "eWFveXVuOnlhb0AwODA3MjIyMEE=",
				Email: "yaoyun@jd.com",
			},
		},
	}
}

func WriteDockerCredentialsFile(w io.Writer) error {
	gw := gzip.NewWriter(w)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	dirHeader := &tar.Header{
		// do not forget the '/' after dir
		Name:    ".docker/",
		Mode:    int64(os.ModePerm),
		Uid:     os.Getuid(),
		Gid:     os.Getgid(),
		ModTime: time.Now(),
	}

	if err := tw.WriteHeader(dirHeader); err != nil {
		return err
	}

	content, err := json.Marshal(DefaultDockerCredentialsInfo)
	if err != nil {
		return err
	}

	fileHeader := &tar.Header{
		Name:    ".docker/config.json",
		Mode:    int64(os.ModePerm),
		Uid:     os.Getuid(),
		Gid:     os.Getgid(),
		Size:    int64(len(content)),
		ModTime: time.Now(),
	}

	if err := tw.WriteHeader(fileHeader); err != nil {
		return err
	}

	_, err = tw.Write(content)
	if err != nil {
		return err
	}

	return nil
}
