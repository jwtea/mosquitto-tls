# Mosquitto TLS example
Example TLS config for mosquitto

# Usage

- Ensure that host mqtt points to localhost, or regen certs for domain required.
- Run cert-gen.sh
- Copy certs to docker/mosquitto/certs
- Copy CA to client/certs
- Trust CA `sudo security add-trusted-cert -d -r trustRoot -k /Library/Keychains/System.keychain ./ca.crt`
- Start dc `docker-compose up`
- Run the test client`cd client && go run .`

