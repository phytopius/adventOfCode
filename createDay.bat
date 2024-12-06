@echo off
:: Check if the correct number of arguments is passed
if "%~2"=="" (
    echo Please provide both year and day as arguments.
    echo Usage: script.bat year day
    exit /b
)

:: Set variables for year and day
set year=%1
set day=day%2

:: Create the directory structure
mkdir .\%year%\%day%

:: Create the input.txt and testinput.txt files
echo. > .\%year%\%day%\input.txt
echo. > .\%year%\%day%\testinput.txt

:: Copy main.go to the new directory and rename it as day.go
copy .\template\main.go .\%year%\%day%\%day%.go

echo Folder and files created successfully under .\%year%\%day%\
