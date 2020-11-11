![Factorio Item Browser](https://raw.githubusercontent.com/factorio-item-browser/documentation/master/asset/image/logo.png) 

# Export Icon Renderer

[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/factorio-item-browser/export-icon-renderer)](https://github.com/factorio-item-browser/export-icon-renderer/releases)
[![GitHub](https://img.shields.io/github/license/factorio-item-browser/export-icon-renderer)](LICENSE.md)
[![build](https://img.shields.io/github/workflow/status/factorio-item-browser/export-icon-renderer/CI?logo=github)](https://github.com/factorio-item-browser/export-icon-renderer/actions)
[![Codecov](https://img.shields.io/codecov/c/gh/factorio-item-browser/export-icon-renderer?logo=codecov)](https://codecov.io/gh/factorio-item-browser/export-icon-renderer)

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