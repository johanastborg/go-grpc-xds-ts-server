#!/bin/bash

# Default port to 8080 if not set
PORT=${PORT:-8080}

# Ensure grpcurl is available
if ! command -v grpcurl &> /dev/null
then
    echo "grpcurl could not be found! Please install it first."
    echo "e.g., brew install grpcurl  OR  go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest"
    exit 1
fi

echo "Connecting to gRPC server at localhost:${PORT}..."
echo "Press Ctrl+C to stop streaming."
echo "---------------------------------------------------------"

# Call the GetLiveStream method using the local proto file (since reflection is not enabled)
# We send an empty JSON object '{}' to satisfy the google.protobuf.Empty argument
grpcurl -plaintext -import-path . -proto telemetry/telemetry.proto -d '{}' localhost:${PORT} telemetry.StreamService/GetLiveStream
