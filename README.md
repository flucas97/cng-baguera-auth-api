# {CNG} Baguera Auth Gateway

### Avaliable routes:
POST /signup

POST /login

GET /cannabis

POST /cannabis

#### Check service status

GET /ping

### Start the project running
#### `$ make`

### GET Redis container IP:
###### `$ docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' container-name`

### Go inside Redis client
###### `$ redis-cli -h 'IP'`

