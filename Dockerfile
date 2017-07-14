 # My Backend Ace
FROM golang:1.7-onbuild
ENV GOPATH $GOPATH
MAINTAINER tk@trevorknott.io

RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

ADD .  /go/src/github.com/tknott95/Ace_Go
CMD cd /go/src/github.com/tknott95/Ace_Go && go get github.com/tknott95/Ace_Go && go build -o /Ace_Go


RUN git config --global url."https://85b7dfe20abfeb5bfb07c7f6d3c9ecba024a124c:x-oauth-basic@github.com/".insteadOf "https://github.com/"

CMD ./run_server.sh