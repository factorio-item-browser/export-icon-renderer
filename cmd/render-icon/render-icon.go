package main

import (
	"encoding/json"
	"fmt"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/env"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/log"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/render"
	"github.com/factorio-item-browser/export-icon-renderer/pkg/transfer"
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
	if len(os.Args) < 2 {
		logger := log.Logger()
		logger.Error().Msg("Missing icon definition as first argument.")
		logger.Info().Msgf(usage, env.Version)
		os.Exit(1)
	}

	var icon transfer.Icon
	err := json.Unmarshal([]byte(os.Args[1]), &icon)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	renderer := render.NewRenderer()
	image, err := renderer.Render(icon)
	if err != nil {
		log.Error(err)
		os.Exit(0)
	}
	fmt.Print(image)
}
