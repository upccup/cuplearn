FROM alpine:3.2
MAINTAINER upccup <yyao@dataman-inc.com>

ADD . /go/src/gotest/git-study/cuplearn/go-study/docker-image

WORKDIR /go/src/gotest/git-study/cuplearn/go-study/docker-image

RUN sh ./build.sh

CMD sh ./run.sh
