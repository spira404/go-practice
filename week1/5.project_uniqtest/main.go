package main

import(
	"os"
	"fmt"
	"bufio"
	"io"
)

func uniq(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(os.Stdin)
	var prev string
	for in.Scan() {
		txt := in.Text()
		if txt == prev {
			continue
		}
		if txt < prev {
			return fmt.Errorf("file not sorted")
		}
		prev = txt

		fmt.Fprintln(output, txt)
	}
	return nil
}

func main() {
	err := uniq(os.Stdin, os.Stdout)
	if err != nil {
		panic(err.Error())
	}
}
