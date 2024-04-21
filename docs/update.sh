#!/bin/bash
DOCS_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" &>/dev/null && pwd)

cd "$DOCS_DIR/../examples"
for i in $DOCS_DIR/*.tape; do
    vhs $i
done
