FROM mcr.microsoft.com/devcontainers/base:ubuntu

RUN curl -fsSL https://deb.nodesource.com/setup_20.x | sh - && apt-get install nodejs -y
RUN curl -fsSL https://go.dev/dl/go1.22.0.linux-amd64.tar.gz | tar -xz -C /usr/local
ENV PATH=$PATH:/usr/local/go/bin

USER vscode
ENV SHELL=/bin/bash
RUN curl -fsSL https://get.pnpm.io/install.sh | sh -
