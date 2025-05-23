# Prerequisites

This project requires Go to be installed on your machine. Follow the instructions below to install Go on your operating system.

## Installing Go

### macOS

1. **Using Homebrew (recommended):**
   ```bash
   brew install go
   ```

2. **Manual Installation:**
   - Visit the [official Go downloads page](https://golang.org/dl/).
   - Download the macOS package (e.g., `go1.21.0.darwin-amd64.pkg`).
   - Open the downloaded package and follow the installation instructions.

### Linux

1. **Using Package Manager (Ubuntu/Debian):**
   ```bash
   sudo apt update
   sudo apt install golang-go
   ```

2. **Using Package Manager (Fedora):**
   ```bash
   sudo dnf install golang
   ```

3. **Manual Installation:**
   - Visit the [official Go downloads page](https://golang.org/dl/).
   - Download the Linux package (e.g., `go1.21.0.linux-amd64.tar.gz`).
   - Extract the archive to `/usr/local`:
     ```bash
     sudo tar -C /usr/local -xzf go1.21.0.linux-amd64.tar.gz
     ```
   - Add Go to your PATH by adding the following line to your `~/.bashrc` or `~/.zshrc`:
     ```bash
     export PATH=$PATH:/usr/local/go/bin
     ```
   - Reload your shell configuration:
     ```bash
     source ~/.bashrc  # or source ~/.zshrc
     ```

### Windows

1. **Using the Installer:**
   - Visit the [official Go downloads page](https://golang.org/dl/).
   - Download the Windows installer (e.g., `go1.21.0.windows-amd64.msi`).
   - Run the installer and follow the instructions.

2. **Manual Installation:**
   - Download the Windows package (e.g., `go1.21.0.windows-amd64.zip`).
   - Extract the archive to a directory (e.g., `C:\Go`).
   - Add the Go `bin` directory to your PATH:
     - Open System Properties > Advanced > Environment Variables.
     - Under System Variables, find the `Path` variable, select it, and click Edit.
     - Add `C:\Go\bin` to the list of paths.

## Verifying the Installation

After installation, open a terminal or command prompt and run:

```bash
go version
```

You should see the installed Go version (e.g., `go version go1.21.0 darwin/amd64`).

## Next Steps

Once Go is installed, you can clone this repository and start working on the Go playground problems. Happy coding!
