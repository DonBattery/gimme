package main_test

// Integration tests for the Gimme CLI

import (
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/suite"
)

type CliSuite struct {
	suite.Suite
}

func (suite *CliSuite) exitCodeFromError(err error) (int, error) {
	cmdExitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus, ok := exitError.Sys().(syscall.WaitStatus)
			if !ok {
				return 1, errors.New("Failed to cast exit status")
			}
			cmdExitCode = waitStatus.ExitStatus()
		}
		return cmdExitCode, nil
	}
	return 0, nil
}

func (suite *CliSuite) runCmd(cmdName string, args ...string) (int, string) {
	cmd := exec.Command(cmdName, args...)
	outBytes, cmdErr := cmd.CombinedOutput()
	suite.NotContainsf(fmt.Sprint(cmdErr), "executable file not found in $PATH", "%s: command not found", cmdName)

	retCode, castErr := suite.exitCodeFromError(cmdErr)
	suite.NoError(castErr, "The return code of the command should be readable")

	return retCode, strings.TrimSpace(string(outBytes))
}

func (suite *CliSuite) SetupSuite() {
	suite.NoError(exec.Command("go", "install").Run(), "Gimme should be installed")
}

func (suite *CliSuite) TestBasics() {
	tests := []struct {
		Command        string
		Args           []string
		ReturnCode     int
		ExpectedOutput []string
	}{
		{
			Command:    "gimme",
			Args:       []string{},
			ReturnCode: 1,
			ExpectedOutput: []string{
				"Failed to load configs from file: Config File \"gimme_conf\" Not Found in",
			},
		},
	}

	for _, testCase := range tests {
		code, text := suite.runCmd(testCase.Command, testCase.Args...)
		suite.Equal(testCase.ReturnCode, code, "Return code should be match")
		for _, line := range testCase.ExpectedOutput {
			suite.Contains(text, line, "The output of the command should contain expected line")
		}
	}
}

func TestCliSuite(t *testing.T) {
	suite.Run(t, new(CliSuite))
}
