#FROM golang:1.11-alpine
#ENV GO111MODULE on
#ENV GOPROXY https://mirrors.aliyun.com/goproxy/
#
#WORKDIR /go/src/github.com/MontFerret/ferret-server
#
#COPY . /go/src/github.com/MontFerret/ferret-server
#
#RUN go build && mv ferret-server /go/bin/
#
#CMD ["ferret-server"]

FROM golang:1.12-alpine AS build

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE on
ENV GOPROXY https://mirrors.aliyun.com/goproxy/

RUN go mod vendor
RUN go install -v ./...

FROM alpine:latest
WORKDIR /root
COPY --from=build /go/src/app .
COPY --from=build /go/bin/ferret-server /usr/local/bin

EXPOSE 8080
CMD ["ferret-server"]
