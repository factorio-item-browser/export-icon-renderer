# Factorio Item Browser - Export Icon Renderer

This project implements the icon renderer of the export in Go for fast processing of the images as layers. The icon
renderer was first part of the exporter itself (i.e. PHP), but after discovering performance issues especially with
larger icons it has been extracted to a Go binary to get a huge performance boost. 

The icon renderer has been tailored to the exporter project, and not really meant to be used as a separate tool, as it
requires the data and directory structure of the exporter itself.

## Usage

Call the binary without parameters to get the usage note as well.

```
Usage: render-icon icon_definition
    icon_definition: The JSON string defining the icon and its layers to render. This string must represent an Icon of
                     the factorio-item-browser/export-data PHP package.

Environment variables:
    FACTORIO_DATA_DIRECTORY: The data directory of the Factorio game, containing the base and core mods.
    FACTORIO_MODS_DIRECTORY: The mods directory of the Factorio game, containing all used mods. All mods must already
                             be UNZIPPED and the directory name of the mod must NOT contain its version number.
```