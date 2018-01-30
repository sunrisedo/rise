@echo off
SET GOPATH=%GOPATH%;%CD%\..\..

REM @echo on

Title rise
go build -o rise.exe github.com/sunrisedo/rise && rise.exe
