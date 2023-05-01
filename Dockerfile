# run with:
# docker build -t lazyaws .
# docker run -it lazyaws:latest /bin/sh

FROM golang:1.18 as build 
WORKDIR /go/src/github.com/BSteffaniak/lazyaws/
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.15
RUN apk add --no-cache -U git xdg-utils
WORKDIR /go/src/github.com/BSteffaniak/lazyaws/
COPY --from=build /go/src/github.com/BSteffaniak/lazyaws ./
COPY --from=build /go/src/github.com/BSteffaniak/lazyaws/lazyaws /bin/
RUN echo "alias gg=lazyaws" >> ~/.profile

ENTRYPOINT [ "lazyaws" ]
