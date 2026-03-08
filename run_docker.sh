#!/bin/bash
set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}==> Loading image into Docker...${NC}"
bazel run //:load

echo -e "${GREEN}==> Starting container on http://localhost:8080/api/hello${NC}"
echo -e "${BLUE}==> Press Ctrl+C to stop the server${NC}"
docker run --rm -it -p 8080:8080 go-grpc-xds-ts-server:latest
