package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildSixtySteps(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestBuildSixtySteps")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "foobuildsixtysteps", ".")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("foobuildsixtysteps")

	logDone("build - build an image with sixty build steps")
}

func TestAddSingleFileToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to root")
}

func TestAddSingleFileToExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to existing dir")
}

func TestAddSingleFileToNonExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToNonExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to non-existing dir")
}

func TestAddDirContentToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "DirContentToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add directory contents to root")
}

func TestAddDirContentToExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "DirContentToExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add directory contents to existing dir")
}

func TestAddWholeDirToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "WholeDirToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add whole directory to root")
}

func TestContextTar(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestContextTar")
	buildBuilderCmd := exec.Command(dockerBinary, "build", "-t", "contexttarbuilder", ".")
	buildBuilderCmd.Dir = buildDirectory

	out, exitCode, err := runCommandWithOutput(buildBuilderCmd)
	if err != nil || exitCode != 0 {
		t.Fatalf("builder build failed to complete: %v %v", out, err)
	}

	builderCmd := exec.Command(dockerBinary, "run", "contexttarbuilder")
	stdout, err := builderCmd.StdoutPipe()
	if err != nil {
		t.Fatalf("failed to get builder stdout, %v", err)
	}
	buildRunnerCmd := exec.Command(dockerBinary, "build", "-t", "contexttarrunner", "-")
	buildRunnerCmd.Stdin = stdout

	err = builderCmd.Start()
	if err != nil {
		t.Fatalf("failed to start builder, %v", err)
	}

	out, exitCode, err = runCommandWithOutput(buildRunnerCmd)
	if err != nil || exitCode != 0 {
		t.Fatalf("runner build failed to complete: %v %v", out, err)
	}

	err = builderCmd.Wait()
	if err != nil {
		t.Fatalf("builder run failed, %v", err)
	}

	out, exitCode, err = cmd(t, "run", "contexttarrunner")
	if out != "bar\n" {
		t.Fatalf("runner produced invalid output: %q, expected %q", out, "bar")
	}

	deleteImages("contexttarbuilder")
	deleteImages("contexttarrunner")
	logDone("build - build an image with a context tar")
}

func TestNoContext(t *testing.T) {
	buildCmd := exec.Command(dockerBinary, "build", "-t", "nocontext", "-")
	buildCmd.Stdin = strings.NewReader("FROM busybox\nCMD echo ok\n")

	out, exitCode, err := runCommandWithOutput(buildCmd)
	if err != nil || exitCode != 0 {
		t.Fatalf("build failed to complete: %v %v", out, err)
	}

	out, exitCode, err = cmd(t, "run", "nocontext")
	if out != "ok\n" {
		t.Fatalf("run produced invalid output: %q, expected %q", out, "ok")
	}

	deleteImages("nocontext")
	logDone("build - build an image with no context")
}

// TODO: TestCaching

// TODO: TestADDCacheInvalidation
