# Challenge 1 | Write your own wc tool

First challenge of the challenge series by John Cricket <https://codingchallenges.fyi/challenges/challenge-wc>

## Description

This WC tool is written using golang under ccwc. This tool is used to count the number of bytes, lines, words and multibytes in a given file and from the standard input.

## Usage

1. Download the file test.txt, provided by the challenge, as defined in the step zero, using the following command:
    ```curl -sSL -o tests/test.txt 'https://www.dropbox.com/scl/fi/d4zs6aoq6hr3oew2b6a9v/test.txt?rlkey=20c9d257pxd5emjjzd1gcbn03&e=3&dl=1'```

2. Use `go run main.go` to list the available flags
    The following options are supported:

    - -c: prints the number of bytes in the given file or stdin
    - -w: prints the number of words in the given file or stdin
    - -l: prints the number of lines in the given file or stdin
    - -m: prints the number of multibytes in the given file or stdin

3. Use `go run main.go [option] <file path>` to run the program

4. The tool can also be used in stdin mode as follows:
   - `printf "content" | go run main.go [option]`
   - `cat filename | go run main.go [option]`

5. Run tests
    To run the tests go to the root ccwc repository and run the following command:

    `go test ./..`

    All unit test under /cmd and integration tests under /test will be executed.
