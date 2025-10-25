#!/usr/bin/env bash
set -euo pipefail
BIN_NAME=dockshield
OUTDIR=dist
mkdir -p ${OUTDIR}
for GOOS in linux darwin windows; do
  for GOARCH in amd64 arm64; do
    echo "Building for ${GOOS}/${GOARCH}..."
    ext=""
    if [ "${GOOS}" = "windows" ]; then ext=".exe"; fi
    CGO_ENABLED=0 GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${OUTDIR}/${BIN_NAME}-${GOOS}-${GOARCH}${ext} ./cmd/dockshield
  done
done
echo "Build artifacts are in ${OUTDIR}/"
