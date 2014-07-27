FROM google/golang

WORKDIR /gopath/src/github.com/ecnahc515/goplayutils
ADD . /gopath/src/github.com/ecnahc515/goplayutils
RUN go get github.com/GeertJohan/go.rice/rice
RUN rice --import-path=github.com/ecnahc515/goplayutils/server embed-go
RUN go get github.com/ecnahc515/goplayutils/cmd/goplay

EXPOSE 8080
ENTRYPOINT ["goplay"]
CMD ["-d", "-addr", "0.0.0.0:8080"]

