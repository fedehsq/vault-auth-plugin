FROM ubuntu:20.04
RUN apt-get update && apt-get install -y software-properties-common curl gnupg2 && \
  curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add - && \
  apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
  apt-get update && apt-get install -y \
  vault && \
  setcap cap_ipc_lock= /usr/bin/vault

# Setup folders
RUN mkdir /vault-operator
WORKDIR /vault-operator

# Copy the source from the current directory to the working Directory inside the container
COPY . .
CMD bash vault-setup.sh