FROM golang:alpine as build
# set env vars for go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
# set workdir
WORKDIR /build
# copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download
# copy the code into the container
COPY . .
# build the application
RUN go build -o main .

FROM alpine as run
# move to /dist directory as the place for resulting binary folder
WORKDIR /dist
# 
RUN apk --no-cache add ca-certificates
# copy binary from build to main folder
COPY --from=build /build/main .
# export necessary port
EXPOSE 8080
# command to run when starting the container
CMD ["/dist/main"]