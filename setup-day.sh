#!/bin/bash

# Check if the number of arguments is correct
if [ "$#" -ne 1 ]; then
    echo "Usage: $0 expected 1 argument, <directory_name> received $# instead"
    exit 1
fi

# Construct directory name
directory_name="day$1"

# Check if directory exists
if [ -d "$directory_name" ]; then
    echo "Directory \"$directory_name\" already exists. Skipping..."
fi

# Make directory
if [ ! -d "$directory_name" ]; then
    mkdir "$directory_name"
fi

# Enter directory
cd "$directory_name" || exit

# Create main file
file_name="main.go"
create_puzzle_func() {
    local input_param="$1"

    line1="\nfunc puzzle$input_param(filePath string) int {"
    line2="\tfile := shared.OpenFile(filePath)"
    line3="\tdefer file.Close()"
    line4="\n\treturn -1\n}"


    local result="$line1"$'\n'"$line2"$'\n'"$line3"$'\n'"$line4"
    echo "$result"
}

touch "$file_name"
echo -e "package main\n\nimport (\n\t\"Misc/aoc2023/shared\"\n\t\"fmt\"\n)" > "$file_name"
echo -e "\nfunc main() {" >> "$file_name"
echo -e "\tanswer1 := puzzle1(\"\")" >> "$file_name"
echo -e "\tfmt.Printf(\"The answer to problem 1 is: %d\\\n\", answer1)\n" >> "$file_name"
echo -e "\tanswer2 := puzzle2(\"\")" >> "$file_name"
echo -e "\tfmt.Printf(\"The answer to problem 2 is: %d\\\n\", answer2)" >> "$file_name"
echo -e "}" >> "$file_name"

func1=$(create_puzzle_func "1")
echo -e "$func1" >> "$file_name"
func2=$(create_puzzle_func "2")
echo -e "$func2" >> "$file_name"

# Make Children directories
create_child_dirs() {
    local dir_name="$1"

    if [ ! -d "$dir_name" ]; then
        mkdir "$dir_name"
    fi

}
$(create_child_dirs "utils")
$(create_child_dirs "structs")


# Enter utils directory
cd "utils"

# Make utils file
file_name="utils.go"
touch "$file_name"
echo -e "package utils\n" > "$file_name"

cd "../../$directory_name"