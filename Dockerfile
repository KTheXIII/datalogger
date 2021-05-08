FROM golang:1.16.4-alpine3.13

# Set workspace directory and copy over the source files
WORKDIR /app
COPY . .

# Set the go environment variables for building the app
ENV PATH=$PATH:/app
ENV GIN_MODE=release
RUN go env -w GOBIN=/app

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD ["server"]