# skittyc: simple kitty customizer
A CLI application for customizing your Kitty terminal.

## Table of contents
- [Installation](#installation)
- [Implementation](#implementation)
	- [File structure](#file-structure)
	- [Directory structure](#directory-structure)
- [Usage](#usage)
	- [Customize](#customize)
	- [Setup](#setup)
- [Future plans](#future-plans)

## Installation
In progress.

## Implementation
### File structure
The `kitty.conf` file has a specific structure that works as a guide to modify each setting using the CLI app. If there's a setting that is not included in the ["customize" subcommands list](#customize), then it will be saved in the `# Other` section. 
```shell
# Fonts

# Background and Foreground

# Cursor

# Colors

# Mouse

# Tab bar

# Terminal bell

# Window layout

# Other
```

### Directory structure
Each process is done through the `kitty` directory. Inside of it, the `kitty.conf` file is modified according to the profiles saved in the `profiles` directory. 
```shell
~/.config/
├─ kitty/
│  ├─ kitty.conf
│  ├─ profiles/
│  │  ├─ profile-name.conf
```

## Usage 
### Customize
The `customize` command works for implementing most of the kitty terminal features we can see in the [kitty.conf](https://sw.kovidgoyal.net/kitty/conf/) documentation, where each one (except the "colors" feature) has two flags: `--set` and `--change`. The `--set` flag is for the most principal settings in each feature while the `--change` flag is for "secondary" settings.

This is the list of subcommands that represents the features:
- `backfore`: Background and Foreground settings.
- `bell`: Terminal settings.
- `colors`: Color schemes settings.
- `cursor`: Cursor customization settings.
- `fonts`: Fonts settings.
- `mouse`: Mouse settings.
- `tabbar`: Tab bar settings.
- `window`: Window layout settings.

This is how this command is used:
```shell
skittyc customize [kitty feature] [--set or --change]
```

**Important:** Not all settings are included according to the [kitty.conf](https://sw.kovidgoyal.net/kitty/conf/) documentation. This is done for keeping a consistency between each feature. For example, the Color schemes documentation includes Background and Foreground settings but it was better to implement a different section that is completely dedicated for this. 

### Setup
The `setup` command works for implementing customizations made by the community. The purpose for each subcommand, `themes` and `profiles`, is implementing themes or settings automatically according to the file structure used. 

This is how this command is used:
```shell
skittyc setup themes [--set or --delete]
skittyc setup profiles [--set or --change]
```



## Future plans


