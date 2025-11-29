package tests

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"testing"
)

// The file can be downloaded from https://codingchallenges.fyi/challenges/challenge-wc
var filePath = "./test.txt"
var binaryName = "ccwc-test"

var contents = []string{
	"hello",
	"line1\nline2\nline3\n",
	"hello world, this is Go",
	"héllo",
	"Go is awesome\nLet's code\n",
}

var tests = []struct {
	name     string
	args     []string
	expected []string
}{
	{
		name:     "Count bytes",
		args:     []string{"-c"},
		expected: []string{"5", "18", "23", "6", "25"},
	},
	{
		name:     "Count lines",
		args:     []string{"-l"},
		expected: []string{"0", "3", "0", "0", "2"},
	},
	{
		name:     "Count words",
		args:     []string{"-w"},
		expected: []string{"1", "3", "5", "1", "5"},
	},
	{
		name:     "Count multibytes",
		args:     []string{"-m"},
		expected: []string{"5", "18", "23", "5", "25"},
	},
	{
		name:     "Count default (-l, -w, -c)",
		args:     []string{""},
		expected: []string{"0 1 5", "3 3 18", "0 5 23", "0 1 6", "2 5 25"},
	},
	{
		name:     "Count wrong",
		args:     []string{"-a"},
		expected: []string{},
	},
}

var fileTests = []struct {
	name     string
	args     []string
	expected []string
}{
	{
		name:     "Count bytes",
		args:     []string{"-c"},
		expected: []string{"342190"},
	},
	{
		name:     "Count lines",
		args:     []string{"-l"},
		expected: []string{"7145"},
	},
	{
		name:     "Count words",
		args:     []string{"-w"},
		expected: []string{"58164"},
	},
	{
		name:     "Count multibytes",
		args:     []string{"-m"},
		expected: []string{"339292"},
	},
	{
		name:     "Count default (-l, -w, -c)",
		args:     []string{""},
		expected: []string{"7145 58164 342190"},
	},
	{
		name:     "Count wrong",
		args:     []string{"-a"},
		expected: []string{},
	},
}

func buildBinary() error {
	// Add .exe extension if on Windows
	if runtime.GOOS == "windows" && binaryName[len(binaryName)-4:] != ".exe" {
		binaryName += ".exe"
	}

	buildCmd := exec.Command("go", "build", "-o", binaryName, "../main.go")
	err := buildCmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to compile main.go: %v", err)
	}
	return nil
}

func TestCCWCContents(t *testing.T) {
	err := buildBinary()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("./" + binaryName)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for index, content := range contents {
				cmd := exec.Command("./"+binaryName, test.args...)
				cmd.Stdin = strings.NewReader(content)
				var out bytes.Buffer
				cmd.Stdout = &out

				err := cmd.Run()
				if err != nil {
					expectedErr := "exit status 1"
					if !strings.Contains(strings.TrimSpace(err.Error()), expectedErr) {
						t.Errorf("%s: expected: %s got: %s, for command %s\n", test.name, expectedErr, err, test.args)
					}
				} else if strings.TrimSpace(out.String()) != test.expected[index] {
					t.Errorf("%s:\n\texpected %s\n\tgot %s", test.name, test.expected[index], out.String())
				}
			}
		})
	}
}

func TestCCWCFile(t *testing.T) {
	err := buildBinary()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("./" + binaryName)

	for _, test := range fileTests {
		t.Run(test.name, func(t *testing.T) {
			for index, arg := range test.args {
				args := []string{filePath}
				if len(arg) > 0 {
					args = append([]string{arg}, args...)
				}

				cmd := exec.Command("./"+binaryName, args...)
				var out bytes.Buffer
				cmd.Stdout = &out

				err := cmd.Run()
				if err != nil {
					expectedErr := "exit status 1"
					if !strings.Contains(strings.TrimSpace(err.Error()), expectedErr) {
						t.Errorf("%s: expected: %s got: %s, for command %s\n", test.name, expectedErr, err, test.args)
					}
				} else if output := strings.TrimSpace(strings.Replace(out.String(), filePath, "", -1)); output != test.expected[index] {
					t.Errorf("%s:\n\texpected %s\n\tgot %s", test.name, test.expected[index], output)
				}
			}
		})
	}
}

func TestCCWCStdin(t *testing.T) {
	err := buildBinary()
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove("./" + binaryName)

	for _, test := range fileTests {
		t.Run(test.name, func(t *testing.T) {
			for index, arg := range test.args {
				cmd := exec.Command("./"+binaryName, arg)

				inputFile, fileErr := os.Open(filePath)
				if fileErr != nil {
					log.Fatalf("Error opening input file: %v", fileErr)
					return
				}
				defer inputFile.Close()

				cmd.Stdin = inputFile
				var out bytes.Buffer
				cmd.Stdout = &out

				err := cmd.Run()
				if err != nil {
					expectedErr := "exit status 1"
					if !strings.Contains(strings.TrimSpace(err.Error()), expectedErr) {
						t.Errorf("%s: expected: %s got: %s, for command %s\n", test.name, expectedErr, err, test.args)
					}
				} else if output := strings.TrimSpace(strings.Replace(out.String(), filePath, "", -1)); output != test.expected[index] {
					t.Errorf("%s:\n\texpected %s\n\tgot %s", test.name, test.expected[index], output)
				}
			}
		})
	}
}
