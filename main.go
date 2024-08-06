package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// main handles command-line arguments and calls ASCII functions
func main() {
	align := flag.String("align", "", "Alignment of the ASCII art: left, center, right, justify")
	output := flag.String("output", "", "Output file for the ASCII art")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 || len(args) > 2 {
		fmt.Println("Usage: go run . --align=<type> --output=<filename> [STRING] [BANNER]")
		return
	}

	if *output != "" {
		if len(args) < 1 || len(args) > 2 {
			fmt.Println("Usage: go run . --align=<type> --output=<filename> [STRING] [BANNER]")
			return
		}
		asciiOutput(args, *align, *output)
	} else {
		switch len(args) {
		case 1:
			ascii(args, *align)
		case 2:
			asciiFS(args, *align)
		default:
			fmt.Println("Usage: go run . --align=<type> [STRING] [BANNER]")
		}
	}
}

// ascii generates ASCII art from input string
func ascii(args []string, align string) {
	txt := args[0]
	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	var outputLines []string
	for _, txt := range textSlice {
		if txt != "" {
			artLines := buildASCIIArt(txt, slice)
			if align != "" {
				outputLines = append(outputLines, alignText(artLines, align, 80)...)
			} else {
				outputLines = append(outputLines, artLines...)
			}
		} else {
			outputLines = append(outputLines, "")
		}
	}
	for _, line := range outputLines {
		fmt.Println(line)
	}
}

// asciiOutput generates ASCII art from input string and writes it to a file
func asciiOutput(args []string, align string, outputFile string) {
	txt := args[0]
	var format string
	if len(args) == 2 {
		format = args[1]
	} else {
		format = "standard"
	}

	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile(format + ".txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	var outputLines []string
	for _, txt := range textSlice {
		if txt != "" {
			artLines := buildASCIIArt(txt, slice)
			if align != "" {
				outputLines = append(outputLines, alignText(artLines, align, 80)...)
			} else {
				outputLines = append(outputLines, artLines...)
			}
		} else {
			outputLines = append(outputLines, "")
		}
	}
	str := strings.Join(outputLines, "\n")
	err = os.WriteFile(outputFile, []byte(str), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		os.Exit(1)
	}
}

// asciiFS generates ASCII art from input string and format from a file
func asciiFS(args []string, align string) {
	txt := args[0]
	format := args[1]
	textSlice := strings.Split(txt, "\\n")

	if !charValidation(txt) {
		fmt.Println("Error : invalid char")
		os.Exit(1)
	}
	file, err := os.ReadFile(format + ".txt")
	if err != nil {
		fmt.Println("Error : reading file")
		os.Exit(1)
	}
	slice := strings.Split(string(file), "\n")
	var outputLines []string
	for _, txt := range textSlice {
		if txt != "" {
			artLines := buildASCIIArt(txt, slice)
			if align != "" {
				outputLines = append(outputLines, alignText(artLines, align, 80)...)
			} else {
				outputLines = append(outputLines, artLines...)
			}
		} else {
			outputLines = append(outputLines, "")
		}
	}
	for _, line := range outputLines {
		fmt.Println(line)
	}
}

// buildASCIIArt creates ASCII art for a given text and slice
func buildASCIIArt(txt string, slice []string) []string {
	var artLines []string
	for i := 0; i < 8; i++ {
		var line string
		for _, v := range txt {
			firstLine := int(v-32)*9 + 1 + i
			line += slice[firstLine]
		}
		artLines = append(artLines, line)
	}
	return artLines
}

// charValidation checks if all characters are within the printable ASCII range
func charValidation(str string) bool {
	slice := []rune(str)
	for _, char := range slice {
		if char < 32 || char > 126 {
			return false
		}
	}
	return true
}

// alignText aligns text based on the given alignment type
func alignText(lines []string, align string, width int) []string {
	var alignedLines []string
	for _, line := range lines {
		switch align {
		case "left":
			alignedLines = append(alignedLines, line)
		case "right":
			padding := width - len(line)
			if padding > 0 {
				alignedLines = append(alignedLines, strings.Repeat(" ", padding)+line)
			} else {
				alignedLines = append(alignedLines, line)
			}
		case "center":
			padding := (width - len(line)) / 2
			if padding > 0 {
				alignedLines = append(alignedLines, strings.Repeat(" ", padding)+line)
			} else {
				alignedLines = append(alignedLines, line)
			}
		case "justify":
			words := strings.Fields(line)
			if len(words) > 1 {
				spaces := width - len(strings.Join(words, ""))
				gap := spaces / (len(words) - 1)
				extraSpaces := spaces % (len(words) - 1)
				var justifiedLine string
				for i, word := range words {
					justifiedLine += word
					if i < len(words)-1 {
						justifiedLine += strings.Repeat(" ", gap+1)
						if extraSpaces > 0 {
							justifiedLine += " "
							extraSpaces--
						}
					}
				}
				alignedLines = append(alignedLines, justifiedLine)
			} else {
				spaces := width - len(line)
				alignedLines = append(alignedLines, line+strings.Repeat(" ", spaces))
			}
		default:
			fmt.Println("Invalid alignment type. Use left, center, right, or justify.")
			return nil
		}
	}
	return alignedLines
}
