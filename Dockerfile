FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

WORKDIR $GOPATH/src//go_projects//go_blog
COPY . $GOPATH/src//go_projects//go_blog
RUN go build .

EXPOSE 8099
ENTRYPOINT ["./go_blog"]