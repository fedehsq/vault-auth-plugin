FROM ubuntu:22.10
RUN adduser --disabled-password --gecos '' vagrant
RUN apt-get update && apt-get install -y openssh-server 

# Setup folders
RUN mkdir /etc/vault-ssh-helper.d/
RUN mkdir /host
WORKDIR /host

# Copy the source from the current directory to the working Directory inside the container
COPY . .
COPY sshd /etc/pam.d/
COPY sshd_config /etc/ssh/
COPY vault-ssh-helper /usr/local/bin/
COPY config.hcl /etc/vault-ssh-helper.d/
ENV NOTVISIBLE "in users profile"
RUN echo "export VISIBLE=now" >> /etc/profile

CMD ["/usr/sbin/sshd", "-D"] 