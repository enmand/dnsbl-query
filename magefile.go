// +build mage

package main

import (
	"context"
	"fmt"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"gopkg.in/src-d/go-git.v4"
)

type Go mg.Namespace
type Docker mg.Namespace

const (
	linter = "github.com/golangci/golangci-lint/cmd/golangci-lint"

	imageRepo = "enmand/dnsbl-query"
)

var (
	grun   = sh.RunCmd("go", "run")
	gtest  = sh.RunCmd("go", "test")
	ggen   = sh.RunCmd("go", "generate")
	gbuild = sh.RunCmd("go", "build")

	dbuild = sh.RunCmd("docker", "build")
	dtag   = sh.RunCmd("docker", "tag")
)

func (Go) Lint(ctx context.Context) error {
	if err := grun(linter, "cache", "clean"); err != nil {
		return err
	}

	return grun(linter, "run", "./...")
}

func (Go) Test(ctx context.Context) error {
	return gtest("-v", "./...")
}

func (Go) Generate(ctx context.Context) error {
	return ggen("./...")
}

func (Go) Build(ctx context.Context) error {
	return gbuild("-o", "bin/dnsbl-query", ".")
}

func (Docker) Build() error {
	return dbuild("-t", imageRepo, ".")
}

func (Docker) Tag() error {
	mg.Deps(Docker.Build)

	r, err := git.PlainOpen(".")
	if err != nil {
		return err
	}

	ref, err := r.Head()
	if err != nil {
		return err
	}

	hash := ref.Hash().String()
	return dtag(imageRepo+":latest", fmt.Sprintf("%s:%s", imageRepo, hash))
}
