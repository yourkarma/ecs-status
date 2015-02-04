package main

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/ecs"
	"github.com/yourkarma/ecs-status/config"
	"github.com/yourkarma/ecs-status/server"
)

var (
	credentials aws.CredentialsProvider
)

type Credentials struct {
	AccessKeyID     string
	SecretAccessKey string
}

func main() {
	app := cli.NewApp()
	app.Name = "ecs-status"
	app.Usage = "Status page for an ECS cluster. Accepts the ECS cluster name as argument."
	app.Version = "0.1.0"
	app.Action = func(context *cli.Context) {
		log.Info("main/main: starting ecs-status...")

		start()
	}

	app.Run(os.Args)
}

func start() {
	loadConfig()
	client := newECSClient()

	server.Initialize(client)
	server.Start()
}

func loadConfig() {
	err := config.Load("./config.toml")
	if err != nil {
		log.Infof("main/loadConfig: %s. Falling back to environment variables.", err)
	}
}

func newECSClient() *ecs.ECS {
	credentials = aws.Creds(
		config.Get().Credentials.AccessKeyID,
		config.Get().Credentials.SecretAccessKey,
		"",
	)

	return ecs.New(credentials, "us-east-1", nil)
}
