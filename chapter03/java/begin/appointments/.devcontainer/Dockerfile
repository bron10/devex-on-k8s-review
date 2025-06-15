FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04

RUN echo 'alias podman=docker' >> /etc/bash.bashrc
RUN echo 'alias k=kubectl' >> /etc/bash.bashrc

RUN curl -sLS https://get.arkade.dev | sudo sh && \
    arkade system install go && \
    echo "export PATH=\$PATH:/usr/local/go/bin/" | tee -a /etc/profile && \
    echo "export GOPATH=\$HOME/go" | tee -a /etc/profile

ENV PATH="/usr/local/go/bin:${PATH}" \
    GOPATH="/root/go"

RUN arkade get crane kubectl kubectx kubens kind kustomize helm pack && \
    chmod +x $HOME/.arkade/bin/* && \
    mv $HOME/.arkade/bin/* /usr/local/bin/

RUN ARCH=$(dpkg --print-architecture) && \
    if [ "$ARCH" = "amd64" ]; then \
        SKAFFOLD_ARCH="amd64"; \
    elif [ "$ARCH" = "arm64" ]; then \
        SKAFFOLD_ARCH="arm64"; \
    else \
        echo "Unsupported architecture: $ARCH" && exit 1; \
    fi && \
    curl -Lo skaffold https://storage.googleapis.com/skaffold/releases/latest/skaffold-linux-${SKAFFOLD_ARCH} && \
    install skaffold /usr/local/bin/ && \
    rm skaffold
