FROM google/golang

ENV GOPATH /gopath
WORKDIR /gopath

WORKDIR /
ADD . /messenger
WORKDIR /messenger
