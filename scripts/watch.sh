#!/bin/bash

# Check if the ADDRESS environment variable is set
if [ -z "$ADDRESS" ]; then
  echo "Please set the ADDRESS environment variable."
  exit 1
fi

# Get filepath from stdin
read -p "Enter the filepath to watch: " filepath

# Check if the file exists
if [ ! -f "$filepath" ]; then
  echo "File $filepath does not exist."
  exit 1
fi

# Use entr to watch the file and execute the command when it changes
echo "$filepath" | entr -s "CODE=\$(cat $filepath); blitd tx script update-script --address $ADDRESS --code \"\$CODE\" --from alice -y"
