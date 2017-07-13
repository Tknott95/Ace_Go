 # My Backend Ace
FROM golang:1.7-onbuild
ENV GOPATH $GOPATH

RUN echo "[url \"git@github.com:\"]\n\tinsteadOf = https://github.com/" >> /root/.gitconfig
RUN mkdir /root/.ssh && echo "StrictHostKeyChecking no " > /root/.ssh/config

MAINTAINER tk@trevorknott.io

ADD .  /go/src/github.com/tknott95/Ace_Go
CMD cd /go/src/github.com/tknott95/Ace_Go && go get github.com/tknott95/Ace_Go && go build -o /tknott95