FROM golang:1.5.1
MAINTAINER peter.edge@gmail.com

RUN \
  apt-get update -yq && \
  apt-get install -yq --no-install-recommends \
    autoconf \
    automake \
    build-essential \
    git \
    libtool \
    unzip && \
  apt-get clean && \
  rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN \
  wget https://codeload.github.com/google/protobuf/tar.gz/v3.0.0-beta-1 && \
  tar xvzf v3.0.0-beta-1 && \
  rm v3.0.0-beta-1 && \
  cd protobuf-3.0.0-beta-1 && \
  ./autogen.sh && \
  ./configure --prefix=/usr && \
  make && \
  make check && \
  make install && \
  cd - && \
  rm -rf protobuf-3.0.0-beta-1

RUN \
  git clone https://github.com/grpc/grpc.git && \
  cd grpc && \
  git submodule update --init && \
  make && \
  make install

RUN \
  go get -v \
    github.com/golang/protobuf/proto/... \
    github.com/golang/protobuf/protoc-gen-go/... \
    go.pedge.io/google-protobuf/... \
    go.pedge.io/googleapis/... \
    go.pedge.io/protolog/cmd/protoc-gen-protolog/... \
    github.com/gengo/grpc-gateway/protoc-gen-grpc-gateway/... \
    google.golang.org/grpc

RUN mkdir -p /go/src/go.pedge.io/protoeasy
ADD . /go/src/go.pedge.io/protoeasy/
WORKDIR /go/src/go.pedge.io/protoeasy
RUN make install
CMD ["protoeasy", "-d"]
