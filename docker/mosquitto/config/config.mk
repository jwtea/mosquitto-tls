BACKEND_CDB ?= no
BACKEND_MYSQL ?= no
BACKEND_SQLITE ?= no
BACKEND_REDIS ?= no
BACKEND_POSTGRES ?= yes
BACKEND_LDAP ?= no
BACKEND_HTTP ?= no
BACKEND_JWT ?= no
BACKEND_MONGO ?= no
BACKEND_FILES ?= no
BACKEND_MEMCACHED ?= no

# Specify the path to the Mosquitto sources here
MOSQUITTO_SRC = /mosquitto/mosquitto-1.4.14

CFG_CFLAGS = -DRAW_SALT
# Specify the path the OpenSSL here
OPENSSLDIR = /usr
