#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Default to amd64 (compatible with Cloud Run)
ARCH=${1:-amd64}

if [[ "$ARCH" == "arm64" ]]; then
    TARGET="//:push_arm64"
elif [[ "$ARCH" == "amd64" ]]; then
    TARGET="//:push_amd64"
else
    echo "Unsupported architecture: $ARCH (use 'amd64' or 'arm64')"
    exit 1
fi

echo -e "${BLUE}==> Pushing OCI image ($ARCH) to Artifact Registry...${NC}"
bazel run "$TARGET"

echo -e "${GREEN}==> Image ($ARCH) successfully pushed!${NC}"
