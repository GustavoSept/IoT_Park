# DON'T USE IN PRODUCTION
In a running project, these keys (at least the private key) wouldn't be uploaded publicly on github.
To generate them yourself, run these commands on the keys dir:

### Private Key 
`openssl genrsa -out app.rsa 2048`

### Public Key
`openssl rsa -in app.rsa -pubout > app.rsa.pub`