# Go-ShowDir (showdir)

A tiny Go CLI to print a colorized tree of the current directory, similar to Unix `tree`. Built for Windows but works anywhere ANSI colors are supported.

Author: [@nghiaomg](https://github.com/nghiaomg)

## Features
- Colorized output per file type and folders
- Sorted: directories first, then files (alphabetical)
- Box-drawing connectors (├──, └──, │) for a clear tree view

## Requirements
- Windows 10 or later recommended (for ANSI colors in console)
- If building from source: Go 1.20+ installed

## Installation

### Option A: Download the prebuilt executable (Windows)
1. Copy `showdir.exe` to a folder you control, e.g. `C:\Tools\showdir`.
2. Add that folder to your PATH using the helper script:
   - Double‑click `add-env.bat` located in the same folder as `showdir.exe`.
   - Close the window when done, then open a new Command Prompt or PowerShell.

After this, you can run `showdir` from any directory.

### Option B: Build from source
1. Clone the repository:
```bash
git clone https://github.com/nghiaomg/Go-ShowDir.git
cd Go-ShowDir
```
2. Build:
```bash
go build -o showdir.exe
```
3. (Optional) Add the folder to your PATH using `add-env.bat` (run as Administrator; same steps as in Option A).

## Usage
From any directory you want to inspect:
```bash
showdir
```
This prints the name of the current directory in blue, then a colorized tree of subdirectories and files.

Example output (colors shown in a compatible terminal):
```
MyProject
├── src
│   ├── main.go
│   └── utils
│       └── helper.go
└── README.md
```

## Notes
- Colors rely on ANSI escape sequences. Modern Windows terminals (Windows Terminal, recent PowerShell/CMD on Windows 10+) support them by default.
- If you do not see colors, ensure you are using Windows Terminal or a recent console, and run a new session after updating PATH.

## Uninstall
- Remove `showdir.exe` from its folder.
- Remove the folder from the system PATH (Settings → System → About → Advanced system settings → Environment Variables → Path).

## License
This project is provided as-is. See repository for details if a license file is added.
