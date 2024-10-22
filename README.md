# cli-datacontract

[![Build Status](https://github.com/merlindorin/cli-datacontract/actions/workflows/golangci.yml/badge.svg)](https://github.com/merlindorin/cli-datacontract/actions/workflows/golangci.yml)
[![Test Status](https://github.com/merlindorin/cli-datacontract/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/merlindorin/cli-datacontract/actions/workflows/goreleaser.yml)
[![Test Status](https://github.com/merlindorin/cli-datacontract/actions/workflows/trufflehog.yml/badge.svg)](https://github.com/merlindorin/cli-datacontract/actions/workflows/trufflehog.yml)

> `cli-datacontract` is a command-line interface for importing data contracts from different sources. It facilitates
> integration with Google BigQuery, providing options to import data contracts from remote locations or local files.

## Table of Contents

<!-- TOC -->
* [cli-datacontract](#cli-datacontract)
  * [Table of Contents](#table-of-contents)
  * [Summary](#summary)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
    * [Binaries via GitHub Releases](#binaries-via-github-releases)
    * [Script Installation](#script-installation)
    * [Run with Docker](#run-with-docker)
  * [Usage](#usage)
    * [Common Flags:](#common-flags)
    * [Commands:](#commands)
      * [Examples](#examples)
  * [Development](#development)
  * [Contributing](#contributing)
  * [License](#license)
<!-- TOC -->

## Summary

`cli-datacontract` provides a command-line tool for importing data contracts, allowing seamless integration with remote
sources or files, particularly centered around BigQuery operations. This tool enhances manageability and consistency
across data workflows.

## Prerequisites

Ensure you have the following:

- [Git](https://git-scm.com): For version control.
- [Go language](https://golang.org/dl/): Needed if building the tool from source.

## Installation

### Binaries via GitHub Releases

1. Visit the [GitHub Releases page](https://github.com/merlindorin/cli-datacontract/releases).
2. Download the appropriate binary for your system.
3. Make the binary executable:
   ```bash
   chmod +x cli-datacontract
   ```
4. Move it to a directory within your PATH.

### Script Installation

Use the installation script:

```bash
curl -sSfL https://raw.githubusercontent.com/merlindorin/cli-datacontract/master/install.sh | sh -s -- -d 
```

### Run with Docker

To run the `cli-datacontract` using Docker:

1. Pull the Docker image:
   ```bash
   docker pull ghcr.io/merlindorin/cli-datacontract:v0.0.1
   ```

2. Run the container:
   ```bash
   docker run --rm ghcr.io/merlindorin/cli-datacontract:v0.0.1 <command> [options]
   ```

   Replace `<command> [options]` with the specific command and options you wish to use (e.g.,
   `bigquery remote --bigquery-projectid=myproject --bigquery-datasetid=mydataset --bigquery-tablename=mytable`).

## Usage

The `cli-datacontract` can be used with various commands and flags:

```bash
Usage: cli-datacontract <command> [flags]
```

### Common Flags:

- `-h, --help`: Show help information.
- `-D, --development`: Enable development mode with debug logging.
- `-l, --level="info"`: Set the logging level (options: debug, info, warn, error, fatal).
- `-c, --config=CONFIG-FLAG`: Specify the path to a configuration file.

### Commands:

- `version`: Display version information.
- `licence`: Show the application's license.
- `bigquery`: Bigquery related commands.

#### Examples

```shell
# Import data contract from a remote BigQuery source.
cli-datacontract bigquery remote --bigquery-projectid=lorem-654 --bigquery-datasetid=ipsum-1234 --bigquery-tablename=foo

# Import data contract from a specified file.
cli-datacontract bigquery file schema.json
```

## Development

To contribute to the project:

1. Clone the repository:
   ```bash
   git clone https://github.com/merlindorin/cli-datacontract.git
   cd cli-datacontract
   ```

2. For building and running tests, utilize Go's built-in tools. Ensure you are working on a feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```

3. Follow best practices for coding and testing.

## Contributing

- Fork this repository and clone your fork.
- Create a feature branch: `git checkout -b feature/your-feature`.
- Commit your changes: `git commit -am 'Add a feature'`.
- Push to the branch: `git push origin feature/your-feature`.
- Open a pull request for review.

## License

This project is licensed under the MIT License. See [LICENSE.md](./LICENSE.md) for more details.
