FROM google/golang

WORKDIR /gopath/src/github.com/etcinit/central/
ADD . /gopath/src/github.com/etcinit/central/
RUN go get ./... && go build ./... && go install ./...

CMD []
ENTRYPOINT ["/gopath/bin/central"]
