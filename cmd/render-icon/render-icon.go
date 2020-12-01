package main

import (
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/command"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/env"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/log"
	"os"
)

const usage = `Factorio Item Browser - Export Icon Renderer 
Version: %s

Usage: render-icon <icon_definition>
    <icon_definition>: The JSON string defining the icon and its layers to render. This string must represent an Icon
                       of the factorio-item-browser/export-data PHP package.

Environment variables:
    FACTORIO_DATA_DIRECTORY: The data directory of the Factorio game, containing the base and core mods.
    FACTORIO_MODS_DIRECTORY: The mods directory of the Factorio game, containing all used mods. All mods must already
                             be UNZIPPED and the directory name of the mod must NOT contain its version number.
`

func main() {
	logger := log.Logger()

	if len(os.Args) < 2 {
		logger.Error().Msg("Missing icon definition as first argument.")
		logger.Info().Msgf(usage, env.Version)
		os.Exit(1)
	}

	cmd := command.NewRenderIcon()
	img, err := cmd.Run(os.Args[1])
	if err != nil {
		log.Error(logger.Error(), err)
		os.Exit(1)
	}

	fmt.Print(img)
}
