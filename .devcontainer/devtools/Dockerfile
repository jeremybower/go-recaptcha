# ------------------------------------------------------------------------------
# Debian
# ------------------------------------------------------------------------------

FROM mcr.microsoft.com/vscode/devcontainers/base:debian-11

# Update the package index
RUN apt-get update -y

# Install required packages
RUN apt-get -y install apt-transport-https
RUN apt-get -y install build-essential
RUN apt-get -y install ca-certificates
RUN apt-get -y install cloc
RUN apt-get -y install curl
RUN apt-get -y install git
RUN apt-get -y install gnupg2
RUN apt-get -y install locales
RUN apt-get -y install manpages-dev
RUN apt-get -y install gnupg
RUN apt-get -y install lsb-release
RUN apt-get -y install make
RUN apt-get -y install software-properties-common
RUN apt-get -y install unzip
RUN apt-get -y install zsh

# Set the locale
RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && locale-gen
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8  

# ------------------------------------------------------------------------------
# Docker
# ------------------------------------------------------------------------------

# Add the GPG key for the official Docker repository
ENV APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=DontWarn
RUN curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
ENV APT_KEY_DONT_WARN_ON_DANGEROUS_USAGE=

# Add the Docker repository to apt sources
RUN add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/debian $(lsb_release -cs) stable" 

# Update the package index
RUN apt-get update -y

# Install docker packages
RUN apt-get -y install docker-ce
RUN apt-get -y install docker-ce-cli
RUN apt-get -y install containerd.io

# ------------------------------------------------------------------------------
# Go
# ------------------------------------------------------------------------------

# Install Go from binaries (https://golang.org/dl/)
ENV GOLANG_VERSION=1.19.5
ENV GOLANG_SHA256=36519702ae2fd573c9869461990ae550c8c0d955cd28d2827a6b159fda81ff95
ENV GOLANG_TAR=go${GOLANG_VERSION}.linux-amd64.tar.gz
RUN curl -sSOL https://dl.google.com/go/${GOLANG_TAR}
RUN echo "${GOLANG_SHA256}  ${GOLANG_TAR}" | sha256sum -c - | grep OK
RUN tar -C /usr/local -xzf ${GOLANG_TAR}
RUN rm ${GOLANG_TAR}

# Setup the Go environment
RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT=/usr/local/go
ENV GOPATH=/go
ENV PATH=/go/bin:/usr/local/go/bin:$PATH

# ------------------------------------------------------------------------------
# Cleanup
# ------------------------------------------------------------------------------

RUN apt-get clean -y

# ------------------------------------------------------------------------------
# Oh my zsh
# ------------------------------------------------------------------------------

RUN wget https://github.com/robbyrussell/oh-my-zsh/raw/master/tools/install.sh -O - | zsh || true
RUN sed -i 's/plugins=(git)/plugins=(git github golang vscode)/' /root/.zshrc
CMD ["zsh"]