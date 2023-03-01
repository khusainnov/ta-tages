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

	e := endpoint.NewEndpoint(client)

	fmt.Printf("Choice witch method to request:\n1 – UploadImages\n2 – ListImages\n> ")

	var cmd string
	for cmd != "0" {
		cmd, err = readInput()
		if err != nil {
			log.Error("cannot read input", zap.Error(err))
		}
		cmd = strings.TrimSuffix(cmd, "\n")

		switch cmd {
		case "1":
			if err = e.Upload(context.Background(), log); err != nil {
				log.Fatal("cannot upload images", zap.Error(err))
			}
		case "2":
			if err = e.List(context.Background(), log); err != nil {
				log.Fatal("cannot get list of images", zap.Error(err))
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
