#!/bin/bash

# Directory containing the files
TARGET_DIR="./public/logos/original/"

# Ensure the directory exists
if [[ ! -d "$TARGET_DIR" ]]; then
  echo "Directory $TARGET_DIR does not exist."
  exit 1
fi

# Rename files to lowercase
for FILE in "$TARGET_DIR"*; do
  if [[ -f "$FILE" ]]; then
    BASENAME=$(basename "$FILE")
    LOWERCASE_NAME=$(echo "$BASENAME" | tr '[:upper:]' '[:lower:]')

    # Only rename if the lowercase name is different
    if [[ "$BASENAME" != "$LOWERCASE_NAME" ]]; then
      mv "$FILE" "$TARGET_DIR$LOWERCASE_NAME"
      echo "Renamed $BASENAME to $LOWERCASE_NAME"
    fi
  fi
done

echo "All files have been renamed to lowercase."
