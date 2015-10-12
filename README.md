# ProjectOmaha - Senior Design
## Software
### How to run the server
#### Windows
1. run `scripts/windows/run.bat`
2. If running for the first time, restart the command line and run again. (environment variables need to be refreshed)

#### Linux / OS X
Just execute `scripts/run`
### Debug Mode
Run the server in debug mode by executing `scripts/run -d 1`. If you're on windows, the command is `go run main.go -d 1`. Running in debug mode will cut out the microcontroller. All changes in state are just made in memory.
