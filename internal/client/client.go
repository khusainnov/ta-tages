package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/khusainnov/tag/internal/client/endpoint"
	"github.com/khusainnov/tag/internal/client/internal"
	tapi "github.com/khusainnov/tag/pkg/tages-api"
	"go.uber.org/zap"
)

func main() {
	log, _ := zap.NewProduction()

	conn, err := internal.NewConn()
	if err != nil {
		log.Fatal("error due get new connection", zap.Error(err))
	}
	defer conn.Close()

	client := tapi.NewImageServiceClient(conn)

	e := endpoint.NewEndpoint(client, log)

	fmt.Printf("Choice witch method to request:\n1 – UploadImages\n2 – ListImages\n3 – DownloadImage\n> ")

	var cmd string
	for cmd != "0" {
		cmd, err = readInput()
		if err != nil {
			log.Error("cannot read input", zap.Error(err))
		}
		cmd = strings.TrimSuffix(cmd, "\n")

		switch cmd {
		case "1":
			for _, image := range internal.Images {
				if err = e.Upload(context.Background(), image); err != nil {
					log.Error("cannot upload images", zap.Error(err))
				}
			}
		case "2":
			if err = e.List(context.Background()); err != nil {
				log.Error("cannot get list of images", zap.Error(err))
			}
		case "3":
			for _, id := range internal.ImageIds {
				if err = e.Download(context.Background(), id); err != nil {
					log.Error("cannot download the image", zap.Error(err))
				}
			}
		default:
			log.Error("command doesn't exists")
		}
	}
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}
