FROM ubuntu:22.10
RUN apt-get update && apt-get install -y wget && rm -rf /var/lib/apt/lists/* \
  && apt-get update -y && apt-get install -y unzip \
  && apt-get install -y openssh-server 

# Setup folders
RUN mkdir /host
WORKDIR /host

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY sshd /etc/pam.d/
COPY sshd_config /etc/ssh/
CMD ["/bin/bash", "host-setup.sh"]