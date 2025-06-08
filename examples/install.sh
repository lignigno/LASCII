#!/bin/bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd "$SCRIPT_DIR/.."
NEED_DEL=$(ls -A)

for item in $NEED_DEL; do
  if [[ "$item" != "examples" && "$item" != "fonts" ]]; then
    rm -rf "$item"
  fi
done

cd -
rm -- "$0"