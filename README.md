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
  * [Installation](#installation)
    * [Binaries via GitHub Releases](#binaries-via-github-releases)
    * [Script Installation](#script-installation)
    * [Run with Docker](#run-with-docker)
  * [Usage](#usage)
    * [Common Flags:](#common-flags)
    * [Commands:](#commands)
      * [Examples](#examples)
  * [Development](#development)
    * [Prerequisites](#prerequisites)
    * [Installing Aqua and Taskfile](#installing-aqua-and-taskfile)
    * [Installing Tools with Aqua](#installing-tools-with-aqua)
    * [Using Taskfile for Development](#using-taskfile-for-development)
    * [Available Tasks](#available-tasks)
    * [Repository Structure](#repository-structure)
  * [Contributing](#contributing)
  * [References](#references)
  * [License](#license)
<!-- TOC -->

## Summary

`cli-datacontract` provides a command-line interface for importing datacontracts, allowing seamless integration with remote sources or files, particularly centered around BigQuery operations. This tool enhances manageability and consistency across data workflows.

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
curl -sSfL https://raw.githubusercontent.com/merlindorin/cli-datacontract/master/install.sh | sh
```

### Run with Docker

For Docker users:

1. Pull the Docker image:
   ```bash
   docker pull ghcr.io/merlindorin/cli-datacontract:v0.0.2
   ```

2. Run the container:
   ```bash
   docker run --rm ghcr.io/merlindorin/cli-datacontract:v0.0.2 <command> [options]
   ```

   Replace `<command> [options]` with the specific command and options you wish to use (e.g., `bigquery remote --bigquery-projectid=myproject --bigquery-datasetid=mydataset --bigquery-tablename=mytable`).

## Usage

The `cli-datacontract` can be used with various commands and flags:

```bash
Usage: datacontract <command> [flags]
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

### Prerequisites

To develop in this project, ensure the following tools are installed:

- [Git](https://git-scm.com): For version control.
- [Go](https://golang.org/dl/): Necessary for building the CLI from source.
- [Docker](https://www.docker.com/): For running the application in containers.
- [Aqua](https://aquaproj.github.io): Efficiently manage CLI tool versions.
- [Taskfile](https://taskfile.dev/): Task runner for consistently automating scripts.

### Installing Aqua and Taskfile

**Aqua:**
1. Follow the [Aqua installation guide](https://aquaproj.github.io/docs/install) to set up Aqua CLI.

**Taskfile:**
1. Follow the [Taskfile installation guide](https://taskfile.dev/#/installation) to set up Taskfile CLI.

### Installing Tools with Aqua

Once Aqua is installed, run the following to install all necessary tools as specified in the `aqua.yaml` file:

```bash
aqua i
```

Aqua ensures all specified tools are installed and up-to-date, leveraging its centralized configuration for consistency across environments.

### Using Taskfile for Development

This project uses Taskfile to automate common development tasks. Upon cloning the repository:

1. Clone the repository:
   ```bash
   git clone https://github.com/merlindorin/cli-datacontract.git
   cd cli-datacontract
   ```

2. Run development tasks using Taskfile:
   ```bash
   task
   ```
   The above command will execute all the default tasks.

### Available Tasks

Here are some of the tasks you can run using Taskfile:

- **Git Tasks**:
    - `git:gitignore`: Write common .gitignore file
    - `git:install`: Install git pre-commit hook

- **Golang Tasks**:
    - `golangci:boilerplate`: Generate golang-ci configuration
    - `golangci:ci`: Generate GitHub Action
    - `golangci:fix`: Fix golang source
    - `golangci:lint`: Lint golang source
    - `golangci:run`: Run golang-ci

- **Goreleaser Tasks**:
    - `goreleaser:boilerplate`: Generate goreleaser configuration
    - `goreleaser:ci`: Generate GitHub Action
    - `goreleaser:install-script`: Generate an installation script
    - `goreleaser:run`: Run goreleaser

- **License Task**:
    - `license:generate`: Generate License

- **Markdown Tasks**:
    - `markdownlint:boilerplate`: Generate markdownlint configuration
    - `markdownlint:fix`: Fix markdown source
    - `markdownlint:lint`: Lint markdown source

- **Trufflehog Tasks**:
    - `trufflehog:ci`: Generate GitHub Action
    - `trufflehog:detect`: Detect secret leaks in the current repository

### Repository Structure

- **cmd**: Contains the main application entry sources.
- **pkg**: Houses package level utilities and libraries.
- **.github**: GitHub-specific configurations, like actions.
- **Taskfile.yaml**: Central task automation configuration.
- **aqua.yaml**: Specifies CLI tools and versions.

## Contributing

- Fork this repository and clone your fork.
- Create a feature branch: `git checkout -b feature/your-feature`.
- Commit your changes: `git commit -am 'Add a feature'`.
- Push to the branch: `git push origin feature/your-feature`.
- Open a pull request for review.

## References

- [Datacontracts Project](https://datacontract.com/)
- [Datacontracts Specification on GitHub](https://github.com/datacontract/datacontract-specification)
- [GitHub Container Registry for cli-datacontract](https://ghcr.io/merlindorin/cli-datacontract)
- [Google BigQuery](https://cloud.google.com/bigquery/)

## License

This project is licensed under the MIT License. See [LICENSE.md](./LICENSE.md) for more details.
