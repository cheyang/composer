package main

import (
	"context"
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/docker/libcompose/docker"
	"github.com/docker/libcompose/docker/ctx"
	"github.com/docker/libcompose/project"
)

const (
	PROJECT_NAME = "PROJECT_NAME"
)

func main() {
	projectName := os.Getenv(PROJECT_NAME)

	if projectName == "" {
		logrus.Panicln("Project Name is not set")
	}

	project, err := docker.NewProject(&ctx.Context{
		Context: project.Context{
			ProjectName: projectName,
		},
	}, nil)

	info, err := project.Ps(context.Background())

	if err != nil {
		logrus.Panicf("err is %s", err.Error())
	}

	logrus.Infof("%v", info)

}
