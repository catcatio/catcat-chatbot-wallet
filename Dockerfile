# create image from the official Go image
FROM golang:alpine

RUN apk add --update tzdata \
  bash wget curl git;

# Create binary directory, install glide and fresh
RUN mkdir -p $$GOPATH/bin && \
  curl https://glide.sh/get | sh && \
  go get github.com/pilu/fresh

# define work directory
WORKDIR /go/src/app

COPY ./runner.conf ./runner.conf
COPY ./main.go ./main.go

RUN glide init --non-interactive && glide update

# RUN go build -o goapp

# VOLUME ["/go/src/app"]

# serve the app
# CMD glide update && fresh -c runner.conf main.go
