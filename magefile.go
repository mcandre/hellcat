// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/magefile/mage/mg"
	"github.com/mcandre/hellcat"
	"github.com/mcandre/mage-extras"
)

// artifactsPath describes where artifacts are produced.
var artifactsPath = "bin"

// Default references the default build task.
var Default = Test

// IntegrationTest executes the integration test suite.
func IntegrationTest() error {
	mg.Deps(Install)

	cmdHelp := exec.Command("hh", "-h")
	cmdHelp.Stdout = os.Stdout
	cmdHelp.Stderr = os.Stderr

	if err := cmdHelp.Run(); err != nil {
		return err
	}

	cmdVersion := exec.Command("hh", "-v")
	cmdVersion.Stdout = os.Stdout
	cmdVersion.Stderr = os.Stderr

	if err := cmdVersion.Run(); err != nil {
		return err
	}

	cmdExamples := exec.Command("hh")
	cmdExamples.Dir = "examples"
	cmdExamples.Stdout = os.Stdout
	cmdExamples.Stderr = os.Stderr

	if err := cmdExamples.Run(); err != nil {
		return err
	}

	cmdTest := exec.Command("hh", "test.txt")
	cmdTest.Dir = "examples"
	cmdTest.Stdout = os.Stdout
	cmdTest.Stderr = os.Stderr

	if err := cmdTest.Run(); err != nil {
		return err
	}

	cmdTestExamine := exec.Command("hh", "-x", "test.txt")
	cmdTestExamine.Dir = "examples"
	cmdTestExamine.Stdout = os.Stdout
	cmdTestExamine.Stderr = os.Stderr

	if err := cmdTestExamine.Run(); err != nil {
		return err
	}

	cmdExamplesExamine := exec.Command("hh", "-x")
	cmdExamplesExamine.Dir = "examples"
	cmdExamplesExamine.Stdout = os.Stdout
	cmdExamplesExamine.Stderr = os.Stderr

	if err := cmdExamplesExamine.Run(); err != nil {
		return err
	}

	cmdTestRecursive := exec.Command("hh", "-r", "test.txt")
	cmdTestRecursive.Dir = "examples"
	cmdTestRecursive.Stdout = os.Stdout
	cmdTestRecursive.Stderr = os.Stderr

	if err := cmdTestRecursive.Run(); err != nil {
		return err
	}

	cmdExamplesRecursive := exec.Command("hh", "-r")
	cmdExamplesRecursive.Dir = "examples"
	cmdExamplesRecursive.Stdout = os.Stdout
	cmdExamplesRecursive.Stderr = os.Stderr
	return cmdExamplesRecursive.Run()
}

// Text runs integration tests.
func Test() error { mg.Deps(IntegrationTest); return nil }

// CoverHTML denotes the HTML formatted coverage filename.
var CoverHTML = "cover.html"

// CoverProfile denotes the raw coverage data filename.
var CoverProfile = "cover.out"

// CoverageHTML generates HTML formatted coverage data.
func CoverageHTML() error { mg.Deps(CoverageProfile); return mageextras.CoverageHTML(CoverHTML, CoverProfile) }

// CoverageProfile generates raw coverage data.
func CoverageProfile() error { return mageextras.CoverageProfile(CoverProfile) }

// GoVet runs go vet with shadow checks enabled.
func GoVet() error { return mageextras.GoVetShadow() }

// GoLint runs golint.
func GoLint() error { return mageextras.GoLint() }

// Gofmt runs gofmt.
func GoFmt() error { return mageextras.GoFmt("-s", "-w") }

// GoImports runs goimports.
func GoImports() error { return mageextras.GoImports("-w") }

// Errcheck runs errcheck.
func Errcheck() error { return mageextras.Errcheck("-blank") }

// Nakedret runs nakedret.
func Nakedret() error { return mageextras.Nakedret("-l", "0") }

// Lint runs the lint suite.
func Lint() error {
	mg.Deps(GoVet)
	mg.Deps(GoLint)
	mg.Deps(GoFmt)
	mg.Deps(GoImports)
	mg.Deps(Errcheck)
	mg.Deps(Nakedret)
	return nil
}

// portBasename labels the artifact basename.
var portBasename = fmt.Sprintf("hellcat-%s", hellcat.Version)

// repoNamespace identifies the Go namespace for this project.
var repoNamespace = "github.com/mcandre/hellcat"

// Goxcart cross-compiles Go binaries with additional targets enabled.
func Goxcart() error {
	return mageextras.Goxcart(
		artifactsPath,
		"-repo",
		repoNamespace,
		"-banner",
		portBasename,
	)
}

// Port builds and compresses artifacts.
func Port() error { mg.Deps(Goxcart); return mageextras.Archive(portBasename, artifactsPath) }

// Install builds and installs Go applications.
func Install() error { return mageextras.Install() }

// Uninstall deletes installed Go applications.
func Uninstall() error { return mageextras.Uninstall("hh") }

// CleanCoverage deletes coverage data.
func CleanCoverage() error {
	if err := os.RemoveAll(CoverHTML); err != nil {
		return err
	}

	return os.RemoveAll(CoverProfile)
}

// Clean deletes artifacts.
func Clean() error { mg.Deps(CleanCoverage); return os.RemoveAll(artifactsPath) }
