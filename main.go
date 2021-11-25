package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execInput(input string) error {
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":

		if len(args) == 1 {
			return fmt.Errorf("cd: missing operand")
		}
		return os.Chdir(args[1])

	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err != nil {
			n, err := fmt.Fprintln(os.Stderr, err)

			if err != nil {
				fmt.Println("Error:", err)
				fmt.Printf("%d bytes written\n", n)
			}
		}

		if err = execInput(input); err != nil {
			n, err := fmt.Fprintln(os.Stderr, err)

			if err != nil {
				fmt.Println("Error:", err)
				fmt.Printf("%d bytes written\n", n)
			}
		}
	}

}
