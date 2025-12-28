package cli

import (
	"fmt"
	"os"
)

func GetInput() []string {
	args := os.Args // []string
	fmt.Println(args)

	return args
}
