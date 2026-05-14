# Zora CLI 🚀
Zora is a lightweight command-line interface tool built in Go for seamless configuration file transformation between JSON and YAML. It focuses on developer efficiency with robust format validation and automated output management.
# ✨FeaturesBi-directional Conversion: 
Effortlessly convert .json to .yaml and vice-versa.Smart Mapping: Maps data directly into Golang structs for type-safe processing.Safety Validation: Prevents redundant conversions by ensuring the output format differs from the input.Automated Output: Generates an output directory automatically or allows a custom path via flags.
# 🏗 ArchitectureZora follows a structured development roadmap divided into two key phases:
Phase 1: Cobra integration, file processing, and data transformation.Phase 2: PostgreSQL integration for data persistence and CRUD operations.
# Supported FormatsInput: 
.json, .yaml, .ymlOutput: json,   yamlConfiguration SchemaZora is optimized for files matching the following structure:name (string)port (int)environment (string)
🛠 InstallationPrerequisitesGo 1.20 or higherSetupBash
# Clone the repository
git clone https://github.com/IbrahimWOT/zora-cli.git
cd zora-cli

# Install dependencies
go mod tidy

# Build the binary
go build -o zora
💻 UsageRun the CLI using the following flag structure:Bash./zora -i <input-file> -o <output-format> [-p <output-path>]
FlagsFlagShortDescriptionRequiredDefault--input-iPath to the input JSON or YAML fileYes---output-oDesired output format (json or yaml)Yes---path-pDirectory for the generated fileNo./outputExamplesConvert JSON to YAML:Bash./zora -i config.json -o yaml
Convert YAML to JSON in a custom directory:Bash./zora -i config.yaml -o json -p my_configs/
# ⚠️ Validation :
Validation RulesTo prevent redundant processing, Zora includes a strict validation layer. If the input format matches the requested output format, the tool will terminate with the following error:Error: Requested output must be in different format than input file
# 📜LICSENSE :
 LicenseDistributed under the MIT License. See the LICENSE file for more information.
# 👤 Author
MD. Ibrahim
# GitHub: 
[IbrahimWOT](https://github.com/IbrahimWOT/zora-cli)