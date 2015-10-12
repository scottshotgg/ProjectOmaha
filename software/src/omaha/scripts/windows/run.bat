@ECHO off
setx GOPATH %~dp0..\..\..\..\
setx SOFTWARE_DIR %~dp0..\..\..\..\
go run %SOFTWARE_DIR%\src\omaha\main.go %*