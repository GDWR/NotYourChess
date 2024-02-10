FROM mcr.microsoft.com/devcontainers/base:ubuntu

WORKDIR  /tmp/goinstall
RUN wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin

WORKDIR /workspace
RUN rm -r /tmp/goinstall
