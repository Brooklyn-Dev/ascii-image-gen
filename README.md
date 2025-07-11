# ascii-image-gen

![License](https://img.shields.io/badge/license-MIT-green)
![GitHub release](https://img.shields.io/github/v/release/Brooklyn-Dev/ascii-image-gen)
![GitHub issues](https://img.shields.io/github/issues/Brooklyn-Dev/ascii-image-gen)

A Go CLI to generate ascii-art from images and output in console and save as text/PNG

![Example](example.gif)

Supported file formats:

-   .gif (single-frame)
-   .jpg, .jpeg
-   .png
-   .webp

## Installation

### Building from Source

You need [Go 1.24+](https://golang.org/dl/) installed.

```bash
git clone https://github.com/Brooklyn-Dev/ascii-image-gen.git
cd ascii-image-gen
go build -o build/bin/ascii-image-gen.exe ./cmd/ascii-image-gen

# Run the executable
cd build/bin
ascii-image-gen.exe
```

### Installing Pre-built Releases

1. Download the latest release from the [Releases](https://github.com/Brooklyn-Dev/ascii-image-gen/releases) page
2. Extract the archive to your desired location
3. Run the executable

## Usage

```
ascii-image-gen
A Go CLI to generate ascii-art from images and output in console and save as text/PNG

Usage:
        ascii-image-gen [OPTIONS] <image-path-1> [<image-path-2> ...]

Options:
        -a,  --aspect-ratio
                Character aspect ratio (width:height) for your terminal font
        -c,  --color, --colour
                Colour the generated ASCII
        -C,  --complex
                Use more detailed character ramp
        -g,  --grayscale, --greyscale
                Grey the generated ASCII
        -i,  --invert
                Invert the character ramp
        -n,  --negative
                Negate colours of all characters
        -v,  --verbose
                Enable verbose logging
        -V,  --version
                Shows version
        -w,  --width
                Width of the generated ASCII
        -x,  --flip-x
                Horizontally flip the generated ASCII
        -y,  --flip-y
                Vertically flip the generated ASCII

        --calibrate
                Calibrate to help manually determine aspect ratio
        --calibration-size
                Size of the calibration test
        --save-bg
                Set background RGBA for saved PNG files
        --save-dir
                Save directory of saved files
        --save-png
                Save generated ASCII in png file(s)
        --save-text
                Save generated ASCII in text file(s)
```

## Third-Party Packages

-   [github.com/gookit/color](github.com/gookit/color)
-   [github.com/leaanthony/go-ansi-parser](github.com/leaanthony/go-ansi-parser)

## Like this project?

If you find this project interesting or useful, consider giving it a star ⭐️!

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for more information.
