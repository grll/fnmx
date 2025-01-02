# fnmx

A lightweight cross-platform wrapper for fnm that provides uvx-like functionality.

It wraps fnm exec and fnm install to make sure the right node version is used and installed before executing in fnm context.

It needs fnm installed. Preferably it uses an fnm binary in the same folder but falls back to fnm in PATH if necessary.

## Installation

Download the appropriate binary for your system from the [releases page](https://github.com/grll/fnmx/releases):
- macOS ARM64 (Apple Silicon): `fnmx-darwin-arm64`
- macOS Intel: `fnmx-darwin-amd64`
- Windows: `fnmx-windows-amd64.exe`

Make the binary executable (macOS):
```bash
chmod +x fnmx-darwin-*
```

## Usage

```bash
fnmx <node-version> <command>
```

For example:

```bash
fnmx v22 npm run dev
```

This will make sure node v22 is installed and then run npm run dev with the right node verison.

