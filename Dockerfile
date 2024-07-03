FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/1stguardleft/discipline-gbackend
COPY . $GOPATH/src/github.com/1stguardleft/discipline-gbackend
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./discipline-gbackend"]