# {CNG} Baguera Auth Gateway

### Start the project running
#### `$ make`

### GET Redis container IP:
###### `$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' cng-baguera-auth-api_cachedb_1`

### Go inside Redis client
###### `$ redis-cli -h 'IP'`
