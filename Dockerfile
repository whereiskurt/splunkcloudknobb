FROM golang

ARG goflags="-mod=vendor"
ENV GOFLAGS=$goflags

RUN git clone https://github.com/magefile/mage && cd mage && go run bootstrap.go

## 1. Make a directory and copy the code into it.
RUN mkdir /scknobb
ADD . /scknobb/

## 2. Move into the directory and start the build.
WORKDIR /scknobb

RUN mage release 

RUN ./release/scknobb.linux.amd64 help