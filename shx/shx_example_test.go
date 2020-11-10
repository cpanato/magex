package shx_test

import (
	"fmt"
	"log"
	"os"

	"github.com/carolynvs/magex/shx"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

func ExampleInDir() {
	// Run `go get -u github.com/gobuffalo/packr/v2/packr2` in /tmp
	pkg := "github.com/gobuffalo/packr/v2/packr2"
	err := shx.InDir(os.TempDir(), func() error {
		return sh.Run("go", "get", "-u", pkg)
	})
	if err != nil {
		log.Fatal("could not install packr2")
	}
}

func ExampleRunS() {
	// Determine whether or not a container exists, without logging stdout
	// or stderr
	containerName := "registry"
	err := shx.RunS("docker", "inspect", containerName)
	if err != nil {
		log.Println("container exists")
	} else {
		log.Println("container does not exist")
	}
}

func ExampleOutputE() {
	// Get the output of `docker --version`, only logging stderr when the command fails.
	versionOutput, err := shx.OutputE("docker", "--version")
	if err != nil {
		log.Println("could not get the docker version")
	}

	fmt.Println(versionOutput)
}

func ExampleCollapseArgs() {
	// Only pass -v to go test when the target was called with -v
	// mage -v test -> go test -v ./...
	// mage test -> go test ./...
	v := ""
	if mg.Verbose() {
		v = "-v"
	}
	sh.RunV("go", "test", v, "./...")
}
