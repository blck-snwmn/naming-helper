# naming-helper

A simple CLI tool that suggests function names based on their descriptions using Claude Code.

## Installation

```bash
go install github.com/blck-snwmn/naming-helper@latest
```

## Usage

```bash
naming-helper "function description"
```

### Example

```bash
$ naming-helper "read a file and calculate its MD5 hash"
{
  "names": [
    "calculateFileMD5",
    "computeFileMD5Hash",
    "generateMD5FromFile",
    "getFileMD5Checksum",
    "hashFileWithMD5"
  ]
}
```

## Requirements

- Go 1.24.2 or later
- [Claude CLI](https://github.com/anthropics/claude-cli) installed and configured

## License

MIT