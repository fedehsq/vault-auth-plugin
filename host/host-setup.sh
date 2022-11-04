#!/bin/bash
echo "Host is setting up SSH configuration..."
wget https://releases.hashicorp.com/vault-ssh-helper/0.2.1/vault-ssh-helper_0.2.1_linux_amd64.zip
unzip -q vault-ssh-helper_0.2.1_linux_amd64.zip -d /usr/local/bin
chmod 0755 /usr/local/bin/vault-ssh-helper
chown root:root /usr/local/bin/vault-ssh-helper
mkdir /etc/vault-ssh-helper.d/
cp config.hcl /etc/vault-ssh-helper.d/
vault-ssh-helper -verify-only -dev -config /etc/vault-ssh-helper.d/config.hcl
#systemctl restart sshd
echo "Setup complete!"
/usr/sbin/sshd -D