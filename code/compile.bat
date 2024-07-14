@echo off
setlocal

set GOFILE=4cget.go
set OUTPUT_DIR=build

rem Create the output directory if it doesn't exist
if not exist %OUTPUT_DIR% (
    mkdir %OUTPUT_DIR%
)

rem Compile for linux-386
echo Compiling for linux-386...
set GOOS=linux
set GOARCH=386
go build -trimpath -ldflags="-s -w" -o %OUTPUT_DIR%\4cget-linux-386 %GOFILE%

rem Compile for linux-amd64
echo Compiling for linux-amd64...
set GOOS=linux
set GOARCH=amd64
go build -trimpath -ldflags="-s -w" -o %OUTPUT_DIR%\4cget-linux-amd64 %GOFILE%

rem Compile for linux-arm
echo Compiling for linux-arm...
set GOOS=linux
set GOARCH=arm
go build -trimpath -ldflags="-s -w" -o %OUTPUT_DIR%\4cget-linux-arm %GOFILE%

rem Compile for windows-386.exe
echo Compiling for windows-386.exe...
set GOOS=windows
set GOARCH=386
go build -trimpath -ldflags="-s -w" -o %OUTPUT_DIR%\4cget-windows-386.exe %GOFILE%

rem Compile for windows-amd64.exe
echo Compiling for windows-amd64.exe...
set GOOS=windows
set GOARCH=amd64
go build -trimpath -ldflags="-s -w" -o %OUTPUT_DIR%\4cget-windows-amd64.exe %GOFILE%

echo Compilation completed.
endlocal
