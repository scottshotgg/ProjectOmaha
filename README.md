# ProjectOmaha - Senior Design
## Software
### How to run the server
#### Windows
1. Set environment variables
 * SOFTWARE_DIR - full path to the software directory in the repo
 * GOPATH - same as SOFTWARE_DIR
2. Execute `go run main.go`

#### Linux / OS X
Just execute `scripts/run`
### Debug Mode
Run the server in debug mode by executing `scripts/run -d 1`. If you're on windows, the command is `go run main.go -d 1`. Running in debug mode will cut out the microcontroller. All changes in state are just made in memory.
