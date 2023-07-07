FROM golang:1.20-alpine AS builder

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

WORKDIR /workspace

# RUN go mod download

# Set necessary environment variables needed for our image and build the server.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

##########################################################################################

FROM builder AS prod

RUN go build -ldflags="-s -w" -o server ./src

ENTRYPOINT ["/workspace/server"]

##########################################################################################

FROM builder AS debug

RUN apk add --no-cache git
RUN apk add --no-cache inotify-tools

ENV CGO_ENABLED 0

RUN go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest
RUN go install -v github.com/ramya-rao-a/go-outline@v0.0.0-20210608161538-9736a4bde949
RUN go install -v golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@2022.1
# RUN go build -gcflags "all=-N -l" -o server ./src


ENTRYPOINT ["sh", "start_script.sh"]
