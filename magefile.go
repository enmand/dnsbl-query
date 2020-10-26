// +build mage

package main

import (
	"context"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Go mg.Namespace

const (
	linter = "github.com/golangci/golangci-lint/cmd/golangci-lint"
)

var (
	grun  = sh.RunCmd("go", "run")
	gtest = sh.RunCmd("go", "test")
	ggen  = sh.RunCmd("go", "generate")
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
