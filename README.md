# ASCII Art Generator

A Go-based command-line tool that converts images (PNG, JPEG) into ASCII art by mapping pixel brightness values to ASCII characters.

## Features

- Supports PNG and JPEG image formats
- Converts images to grayscale
- Maps pixel brightness to a range of ASCII characters
- Customizable output file path
- Simple command-line interface

## Installation

Make sure you have Go 1.23.0 or later installed on your system.

```bash
# Clone the repository
git clone https://github.com/sisypheus/asciigen ascii_gen
cd ascii_gen

# Build the project
go build
```

## Usage

```bash
# Basic usage
./ascii_gen -i input_image.png -o output.txt

# Default output (writes to output.txt)
./ascii_gen -i input_image.jpg
```

### Command Line Arguments

- `-i` : Path to the input image (required)
- `-o` : Path to the output text file (optional, defaults to "output.txt")

## How It Works

1. The program loads an input image and converts it to grayscale
2. Each pixel's brightness value (0-255) is mapped to a corresponding ASCII character
3. Characters are chosen from a predefined set ranging from darkest to lightest:
   ```
   $@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\|()1{}[]?-_+~<>i!lI;:,"^`'.
   ```
4. The resulting ASCII art is written to the specified output file

## Example

```bash
./ascii_gen -i mountain.jpg -o mountain_ascii.txt
```

## Project Structure

```
ascii_gen/
├── main.go          # Main application code
├── go.mod          # Go module file
├── .gitignore      # Git ignore file
└── README.md       # Project documentation
```

## Dependencies

The project uses only Go standard library packages:
- `flag`: Command-line argument parsing
- `image`: Image processing
- `os`: File operations
- `math`: Mathematical operations

## Building

```bash
go build -o ascii_gen
```

## Contributing

Feel free to open issues or submit pull requests for any improvements or bug fixes.

## License

Licensed under the MIT License. See [LICENSE](LICENSE) for more information.