package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// ANSI color codes
const (
	reset  = "\033[0m"
	blue   = "\033[34m"
	green  = "\033[32m"
	yellow = "\033[33m"
	purple = "\033[35m"
	red    = "\033[31m"
	cyan   = "\033[36m"
	gray   = "\033[90m"
	orange = "\033[38;5;208m" // ANSI 256 màu (cam)
	white  = "\033[37m"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(blue + filepath.Base(root) + reset) // thư mục gốc
	err = printTree(root, "", true)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func printTree(path string, prefix string, isRoot bool) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	// sort: folder trước, rồi file
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].IsDir() && !entries[j].IsDir() {
			return true
		}
		if !entries[i].IsDir() && entries[j].IsDir() {
			return false
		}
		return entries[i].Name() < entries[j].Name()
	})

	for i, entry := range entries {
		connector := "├── "
		if i == len(entries)-1 {
			connector = "└── "
		}

		name := entry.Name()
		coloredName := colorize(name, entry.IsDir())

		fmt.Printf("%s%s%s\n", prefix, connector, coloredName)

		if entry.IsDir() {
			newPrefix := prefix
			if i == len(entries)-1 {
				newPrefix += "    "
			} else {
				newPrefix += "│   "
			}
			printTree(filepath.Join(path, name), newPrefix, false)
		}
	}
	return nil
}

func colorize(name string, isDir bool) string {
	if isDir {
		return blue + name + reset
	}

	ext := strings.ToLower(filepath.Ext(name))
	switch ext {
	// code
	case ".js", ".ts", ".jsx", ".tsx":
		return yellow + name + reset
	case ".css", ".scss", ".sass", ".less":
		return purple + name + reset
	case ".html", ".htm", ".xml":
		return red + name + reset
	case ".go", ".rs":
		return cyan + name + reset
	case ".c", ".cpp", ".h", ".hpp":
		return cyan + name + reset
	case ".py":
		return yellow + name + reset
	case ".php":
		return purple + name + reset
	case ".java", ".class":
		return orange + name + reset

	// config / data
	case ".json", ".yaml", ".yml", ".toml", ".ini":
		return yellow + name + reset
	case ".env":
		return gray + name + reset

	// text / docs
	case ".txt", ".log":
		return white + name + reset
	case ".md", ".rst":
		return green + name + reset
	case ".pdf":
		return red + name + reset
	case ".doc", ".docx", ".odt":
		return blue + name + reset
	case ".xls", ".xlsx", ".ods":
		return green + name + reset
	case ".ppt", ".pptx", ".odp":
		return orange + name + reset

	// images
	case ".png", ".jpg", ".jpeg", ".gif", ".svg", ".ico":
		return green + name + reset

	// media
	case ".mp3", ".wav", ".ogg":
		return purple + name + reset
	case ".mp4", ".mkv", ".avi", ".mov":
		return red + name + reset
	}

	// mặc định
	return white + name + reset
}