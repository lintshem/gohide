#!/bin/bash
mkdir -p dist

platforms=(
  "linux/amd64"
  "linux/arm64"
  "windows/amd64"
  "windows/arm64"
  "darwin/amd64"
  "darwin/arm64"
)

for platform in "${platforms[@]}"; do
    GOOS=${platform%/*}
    GOARCH=${platform#*/}
    output="build/gohide-$GOOS-$GOARCH"
    
    # Add .exe for Windows targets
    if [ "$GOOS" = "windows" ]; then
        output+=".exe"
    fi

    echo "Building for $GOOS/$GOARCH..."
    GOOS=$GOOS GOARCH=$GOARCH go build -o "$output" main.go
done
