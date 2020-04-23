package env

import "os"

var (
	FactorioDataDirectory = os.Getenv("FACTORIO_DATA_DIRECTORY")
	FactorioModsDirectory = os.Getenv("FACTORIO_MODS_DIRECTORY")

	Version string
)
