package hello

import (
	"fmt"
	_ "github.com/spf13/cobra"
)

func Hello() {
	fmt.Println("Hello from hello.go function.")
}
