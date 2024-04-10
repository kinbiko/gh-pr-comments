package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
	ghprcomments "github.com/kinbiko/gh-pr-comments"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func readFromStdinPipe() ([]byte, error) {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		return nil, fmt.Errorf("nothing to read from stdin pipe")
	}

	in := []byte{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		in = append(in, scanner.Bytes()...)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("unable to read from stdin: %w", err)
	}
	return in, nil
}

type comment = *ghprcomments.PRComment

func groupByThread(comments []comment) [][]comment {
	m := map[string][]comment{}
	for _, c := range comments {
		tid := c.ThreadID()
		v, ok := m[tid]
		if !ok {
			m[tid] = []comment{c}
		} else {
			m[tid] = append(v, c)
		}
	}

	cc := make([][]comment, 0, len(m))
	for _, v := range m {
		cc = append(cc, v)
	}

	sort.Slice(cc, func(i, j int) bool {
		return cc[i][0].CreatedAt.Before(cc[j][0].CreatedAt)
	})

	return cc
}

func run(args []string) error {
	in, err := readFromStdinPipe()
	if err != nil {
		return err
	}

	comments := []comment{}
	if err := json.Unmarshal(in, &comments); err != nil {
		return fmt.Errorf("unable to umarshall data into PRComment: %w", err)
	}

	threads := groupByThread(comments)

	for _, t := range threads {
		fmt.Printf("\n%s\n", color.CyanString(t[0].ThreadID()))
		for _, c := range t {
			fmt.Println(c)
		}
	}
	return nil
}
