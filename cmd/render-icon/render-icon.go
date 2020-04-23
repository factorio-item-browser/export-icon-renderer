package main

import (
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/command"
	"log"
	"os"
)

func main() {
	stderr := log.New(os.Stderr, "", 0)

	if len(os.Args) < 2 {
		stderr.Printf("Missing icon definition as first argument.\n")
		os.Exit(1)
	}

	cmd := command.NewRenderIcon()
	img, err := cmd.Run(os.Args[1])
	if err != nil {
		stderr.Printf("Failed to render icon: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(img)
}
