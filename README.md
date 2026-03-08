# Go Bazel Gazelle Project

A simple REST API built with Go using only native libraries, managed by Bazel's Bzlmod system with `rules_go` and Gazelle.

## Project Structure

- `main.go`: The Go source code for the REST API.
- `MODULE.bazel`: Bazel module definitions (Bzlmod).
- `BUILD.bazel`: Bazel build instructions.
- `.bazelrc`: Global Bazel configuration flags (optimized for macOS/pure Go).
- `go.mod`: Go module definition.

## Prerequisites

- [Bazel](https://bazel.build/install) (v7.0.0 or later recommended).
- [Go](https://go.dev/doc/install) (v1.23.0 or later recommended).
- macOS users: XCode Command Line Tools.

## Getting Started

### 1. Build the Project

To build the Go binary:
```bash
bazel build //:go-grpc-xds-ts-server
```

### 2. Run the Server

To run the REST API server:
```bash
bazel run //:go-grpc-xds-ts-server
```

Once running, the server listens on port `8080` (default). You can override the port by setting the `PORT` environment variable:
```bash
PORT=9090 bazel run //:go-grpc-xds-ts-server
```

Test it with `curl`:
```bash
curl http://localhost:8080/api/hello
```

### 3. Update Build Files (Gazelle)

If you add new Go files or dependencies, update the Bazel build files using Gazelle:
```bash
bazel run //:gazelle
```

### 4. Build and Run with Docker

To build the OCI container image and run it locally with Docker:
```bash
./run_docker.sh
```

This script will:
1.  Build the OCI loadable tarball using Bazel.
2.  Load the image into your local Docker daemon.
3.  Start a container on port `8080`.

#### Configurable Port
The containerized server also supports the `PORT` environment variable. You can override it at runtime:
```bash
docker run --rm -e PORT=9090 -p 9090:9090 go-grpc-xds-ts-server:latest
```

Once running, the containerized API is available at:
```bash
curl http://localhost:9090/api/hello
```

Alternatively, you can build the image directly:

To build the OCI image:
```bash
bazel build //:image
```

To load the image into Docker manually:
```bash
bazel run //:load
```

Once running, the containerized API is available at:
```bash
curl http://localhost:8080/api/hello
```

### 5. Push to Artifact Registry

To push the OCI image to Google Cloud Artifact Registry (defaults to `amd64` for Cloud Run):
```bash
./push_image.sh
```

You can also push a specific architecture:
```bash
./push_image.sh arm64
./push_image.sh amd64
```

The images are tagged as:
- `latest` (points to `amd64`)
- `latest-amd64`
- `latest-arm64`

> [!IMPORTANT]
> Ensure you are authenticated with Google Cloud and have the necessary permissions to push to the repository:
> ```bash
> gcloud auth configure-docker us-central1-docker.pkg.dev
> ```

This project is configured to work around common macOS Xcode detection issues in `rules_go` by:
1.  Using `apple_support` in `MODULE.bazel`.
2.  Forcing a `pure` Go build (disabling CGO) via `.bazelrc`.

This ensures the project builds "out of the box" even if your local Xcode configuration is not standard.
