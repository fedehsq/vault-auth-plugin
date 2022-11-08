#!/bin/bash
export VAULT_ADDR=http://$(ipconfig getifaddr en0):8200
cd vault-operator
bash vault-setup.sh