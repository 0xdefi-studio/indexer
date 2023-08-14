FROM golang:1.19

# Set the Current Working Directory inside the container
WORKDIR /src/github.com/0xdefi-studio/indexer

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...