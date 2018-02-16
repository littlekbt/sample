# go/nginx-server-push

NGINX server push sample using reverse proxy to go server.

## Usage  

### build
```
$ make
```

note: You must creates crt file and private key under conf/certs.

### run
```
$ docker run -it --rm -p 443:443 -p 18443:18443 -v `pwd`:/go nginx-server-push-sample bash
```

### only nginx
```
$ /nginx/objs/nginx -c `pwd`/conf/go_http2.conf

# access to https://localhost
```

### only go server
```
$ ./bin/server -p

# access to https://localhost:18443
```

### reverse proxy from nginx to go server
```
$ /nginx/objs/nginx -c `pwd`/conf/go_http2.conf
$ ./bin/server

# access to https://localhost/go-proxy
```

