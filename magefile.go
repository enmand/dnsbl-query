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
	grun = sh.RunCmd("go", "run")
)

func (Go) Lint(ctx context.Context) error {
	if err := grun(linter, "cache", "clean"); err != nil {
		return err
	}

	return grun(linter, "run", "./...")
}
