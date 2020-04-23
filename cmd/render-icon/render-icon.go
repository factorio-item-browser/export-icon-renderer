package main

import (
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/command"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/env"
	"os"
)

const usage = `Factorio Item Browser - Export Icon Renderer 
Version: %s

Usage: render-icon icon_definition
    icon_definition: The JSON string defining the icon and its layers to render. This string must represent an Icon of
                     the factorio-item-browser/export-data PHP package.

Environment variables:
    FACTORIO_DATA_DIRECTORY: The data directory of the Factorio game, containing the base and core mods.
    FACTORIO_MODS_DIRECTORY: The mods directory of the Factorio game, containing all used mods. All mods must already
                             be UNZIPPED and the directory name of the mod must NOT contain its version number.
`

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Missing icon definition as first argument.\n\n")
		fmt.Printf(usage, env.Version)
		os.Exit(1)
	}

	cmd := command.NewRenderIcon()
	img, err := cmd.Run(os.Args[1])
	if err != nil {
		fmt.Printf("Failed to render icon: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(img)
}
