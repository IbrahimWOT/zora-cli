Zora CLI
Zora is a lightweight command-line interface tool built in Go for seamless configuration file transformation between JSON and YAML. It focuses on developer efficiency with robust format validation and automated output management.

Features
Bi-directional Conversion: Convert .json to .yaml and vice-versa.

Smart Mapping: Maps data directly into Golang structs for type-safe processing.

Safety Validation: Prevents redundant conversions by ensuring the output format differs from the input.

Automated Output: Generates an output directory automatically or allows a custom path via flags.

Architecture
Zora follows a structured development roadmap:

Phase 1: Cobra integration, file processing, and data transformation.

Phase 2: PostgreSQL integration for data persistence and CRUD operations.

Supported Formats
Input: .json, .yaml, .yml

Output: json, yaml

Configuration Structure
Zora is optimized for files matching this schema:

name (string)

port (int)

environment (string)

Installation
Prerequisites
Go 1.20 or higher

Setup
Clone the repository:
git clone https://github.com/IbrahimWOT/zora-cli.git
cd zora-cli

Install dependencies:
go mod tidy

Build the binary:
go build -o zora

Usage
Run the CLI using the following flag structure:

./zora -i  -o  [-p ]

Flags
-i, --input: Path to the input JSON or YAML file (Required)

-o, --output: Desired output format, either json or yaml (Required)

-p, --path: Directory for the generated file (Default: ./output)

Examples
Convert JSON to YAML:
./zora -i config.json -o yaml

Convert YAML to JSON in a custom directory:
./zora -i config.yaml -o json -p my_configs/

Validation Rules
If the input format matches the requested output format, Zora will terminate with an error message to prevent redundant processing:
Error: Requested output must be in different format than input file

License
Distributed under the MIT License. See LICENSE for more information.

Author
MD. Ibrahim
GitHub: IbrahimWOT