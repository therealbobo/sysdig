FROM ubuntu:22.04

ARG BUILD_TYPE=release
ARG BUILD_DRIVER=OFF
ARG BUILD_BPF=OFF
ARG BUILD_VERSION=dev
ARG BUILD_WARNINGS_AS_ERRORS=OFF
ARG MAKE_JOBS=4

ENV BUILD_TYPE=${BUILD_TYPE}
ENV BUILD_DRIVER=${BUILD_DRIVER}
ENV BUILD_BPF=${BUILD_BPF}
ENV BUILD_VERSION=${BUILD_VERSION}
ENV BUILD_WARNINGS_AS_ERRORS=${BUILD_WARNINGS_AS_ERRORS}
ENV MAKE_JOBS=${MAKE_JOBS}

ARG ZIG_VERSION=0.14.0-dev.2643+fb43e91b2

COPY ./zig-cc  /usr/bin/
COPY ./zig-c++ /usr/bin/

WORKDIR /

RUN apt update && \
    apt install -y --no-install-recommends \
        autoconf \
        automake \
        build-essential \
        ca-certificates \
        clang \
        cmake \
        curl \
        git \
        libc-ares-dev \
        libcurl4-openssl-dev \
        libelf-dev \
        libgrpc++-dev \
        libgtest-dev \
        libjq-dev \
        libjsoncpp-dev \
        libprotobuf-dev \
        libssl-dev \
        libtbb-dev \
        libtool \
        llvm \
        ninja-build \
        pkg-config \
        protobuf-compiler-grpc \
        wget \
        xz-utils && \
    git clone https://github.com/libbpf/bpftool.git --branch v7.3.0 --single-branch && \
    cd bpftool && \
    git submodule update --init && \
    cd src && \
    make install && \
    cd ../.. && \
    rm -fr bpftool && \
    curl -LO https://ziglang.org/builds/zig-linux-$(uname -m)-${ZIG_VERSION}.tar.xz && \
    tar -xaf zig-linux-$(uname -m)-${ZIG_VERSION}.tar.xz && \
    rm -v zig-linux-$(uname -m)-${ZIG_VERSION}.tar.xz && \
    cd zig-linux-$(uname -m)-${ZIG_VERSION} &&  \
    cp -v zig /usr/bin && \
    find lib -exec cp --parents {} /usr/ \; && \
    cd .. && \
    rm -fr zig*
