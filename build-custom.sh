set -ex

docker run --rm -v "$GOPATH":/go -w /go/src/github.com/kth/logspout -e "GOPATH=/go" iron/go:dev sh -c 'go build -o logspout'

# Can build the image too
# docker build -t kthse/logspout:latest .
