# pfyl-cli
pfyl-cli is a command-line application for analyzing binaries, extracting static information and sending these information to a pfyl-server.
## Prerequisites
* Go (https://golang.org/)
## Compile
```shell script
$ git clone git@gitlab.cosee.biz:pfyl/pfyl-cli.git
$ cd pfyl-cli
$ go mod download
$ go build
```
An executable called "pfyl-cli" will be created in the base directory.
## Development
IntelliJ IDEA Ultimate supports Go development via a plugin. Just install the Go-Plugin from the Plugins-Marketplace.<br>
After restarting Intellij, the project can be opened like every other project. <br>
For analyzing, the GCC-ARM toolchain need to be present locally. <br>
The toolchain can be found here: https://developer.arm.com/tools-and-software/open-source-software/developer-tools/gnu-toolchain/gnu-rm <br>
The toolchain can be set via 
```shell script
./pfyl-cli configure
```
In the test-folder, an example binary, and the nm-tool exist for testing purposes.
## Project Structure
### main.go
Entrypoint for application. This is the place for plugging together the various parts of the application and starting the analysis.
### analysis
Contains all analysis logic. The only analysis operation currently implemented is analyzing symbols.
### cmd
Contains all commands the application provides, as well as the flags which can be used to configure the application.
### external
Contains the http-client to communicate with the backend.
### configuration
Contains logic for saving and loading the pfyl-cli configuration from disk.
### test
Contains all data and binaries needed for testing.
