FROM vault:latest

RUN apk add bash

# Setup folders
RUN mkdir /vault-operator
WORKDIR /vault-operator

# Copy the source from the current directory to the working Directory inside the container
COPY . .
CMD bash vault-setup.sh