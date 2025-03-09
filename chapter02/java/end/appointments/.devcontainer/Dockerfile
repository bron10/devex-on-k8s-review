FROM mcr.microsoft.com/devcontainers/java:1-21-bullseye

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive && \
    apt-get -y install httpie && \
    rm -Rf /var/lib/apt/lists/*  && \
    rm -Rf /usr/share/doc && rm -Rf /usr/share/man  && \
    apt-get clean
