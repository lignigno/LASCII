#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR/.."
EXCLUDE="examples fonts"

for dir in */; do
  if [[ ! " $EXCLUDE " =~ " $dir " ]]; then
    rm -rf "$dir"
  fi
done

echo "complete!"