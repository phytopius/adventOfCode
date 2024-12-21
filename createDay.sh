#!/bin/bash

# Check if the correct number of arguments is passed
if [ "$#" -ne 2 ]; then
    echo "Please provide both year and day as arguments."
    echo "Usage: $0 year day"
    exit 1
fi

# Set variables for year and day
year=$1
day="day$2"

# Create the directory structure
mkdir -p ./"$year"/"$day"

# Create the input.txt and testinput.txt files
touch ./"$year"/"$day"/input.txt
touch ./"$year"/"$day"/testinput.txt

# Copy main.go to the new directory and rename it as day.go
cp ./template/main.go ./"$year"/"$day"/"$day".go

echo "Folder and files created successfully under ./$year/$day/"
