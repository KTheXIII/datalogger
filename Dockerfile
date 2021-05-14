FROM golang:1.16.4-alpine3.13 as building

# Set workspace directory and copy over the source files
WORKDIR /app
COPY . .

# Set the go environment variables for building the app
ENV PATH=$PATH:/app
ENV GIN_MODE=release
RUN go env -w GOBIN=/app

RUN go get -d -v ./...
RUN go install -v ./...

# Application container image
FROM alpine:3.13
WORKDIR /app
# COPY . .
COPY --from=building /app/datalogger /app/datalogger

ENV GIN_MODE=release
ENV PATH=$PATH:/app
EXPOSE 8080

CMD ["datalogger"]