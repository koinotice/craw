# caddy-builder

[![Build Status](https://travis-ci.org/openfaas/caddy-builder.svg?branch=master)](https://travis-ci.org/openfaas/caddy-builder)

Build Caddy with plugins from source using Docker multi-build

### Usage

Clone the caddy-builder repository:

```bash
$ git clone https://github.com/openfaas/caddy-builder.git
$ cd caddy-builder
```

Add the Caddy plugins that you want to the `caddy.go` file:

```go
package main

import (
        "github.com/caddyserver/caddy/caddy/caddymain"

        // http.prometheus
        _ "github.com/miekg/caddy-prometheus"
        // http.ipfilter
        _ "github.com/pyed/ipfilter"
)

func main() {
	caddymain.EnableTelemetry = false
	caddymain.Run()
}
```

By default, Go fetches the latest **release** of the plugins you include in your `caddy.go` file. If you
want it to fetch a specific **commit** instead, you also need to enter the repository URL of the plugin
along with the hash of the desired commit into your `go.mod` file.

```go
module caddy

go 1.12

require (
	github.com/caddyserver/caddy v1.0.1

	github.com/pyed/ipfilter 44576102099e4428e7e429987a97a45122c4a656
)
```

Edit the [docker-compose](https://github.com/openfaas/caddy-builder/blob/master/docker-compose.yml) 
file and replace the image prefix with your own repo name:

```yaml
version: "3.3"

services:
  caddy:
    build:
      context: .
      dockerfile: Dockerfile
      args:
        CADDY_VERSION: ${CADDY_VERSION:-1.0.1}
    image: openfaas/caddy:${CADDY_VERSION:-1.0.1}
    container_name: caddy
    ports:
      - 80:80
      - 443:443
      - 9180:9180
```

Build the image with Docker Compose:

```bash
CADDY_VERSION=1.0.1 docker-compose build caddy
```

Run Caddy container exposing 80, 443 and 9180 ports:

```bash
docker-compose up -d
```

Remove the container, `www` volume and image:

```bash
docker-compose down -v --rmi all
```

### Running Caddy with Docker

The [openfaas/caddy](https://hub.docker.com/r/openfaas/caddy/) comes with a default Caddyfile that 
you can override by mounting your own config:

```bash
$ docker run -d --name caddy \
    -v $(pwd)/Caddyfile:/etc/caddy/Caddyfile \
    -p 80:80 \
    openfaas/caddy
```

Mount your site root using the `www` volume:

```bash
$ docker run -d --name caddy \
    -v $(pwd)/Caddyfile:/etc/caddy/Caddyfile \
    -v $(pwd)/site:/www \
    -p 80:80 \
    openfaas/caddy
```

Expose the Prometheus metric endpoint on `http://localhost:9180/metrics`:

```bash
$ docker run -d --name caddy \
    -v $(pwd)/Caddyfile:/etc/caddy/Caddyfile \
    -v $(pwd)/site:/www \
    -p 80:80 -p 9180:9180 \
    openfaas/caddy
```

In your Caddyfile configure the http.prometheus plugin:

```
example.com {
    prometheus 0.0.0.0:9180
    log stdout
    errors stderr
}
```

Persist Let's Encrypt certificates on host:

```bash
$ docker run -d --name caddy \
    -v $(pwd)/Caddyfile:/etc/caddy/Caddyfile \
    -v $(pwd)/certs:/.caddy \
    -p 80:80 -p 443:443 \
    openfaas/caddy
```

In your Caddyfile configure the tls email:

```
example.com {
    tls contact@example.com
}
```

### Running Caddy with Docker Swarm

In order to deploy Caddy with a custom config on Docker Swarm, you need to use 
Docker engine version 17.06 or later. The Caddy image has curl installed so 
you can easily define a health check:

```yaml
version: "3.3"

configs:
  caddy_config:
    file: ./Caddyfile

volumes:
  certs: {}

services:
  caddy:
    image: openfaas/caddy
    ports:
      - 80:80
      - 443:443
    configs:
      - source: caddy_config
        target: /etc/caddy/Caddyfile
    volumes:
      - certs:/.caddy
    deploy:
      mode: replicated
      replicas: 1    
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:80"]
      interval: 5s
      timeout: 1s
      retries: 3
```

### License

The caddy-builder is MIT licensed and the Caddy 
[source code](https://github.com/caddyserver/caddy/blob/master/LICENSE.txt) is Apache 2.0 licensed. 
Because openfaas/caddy is built from source, it's not subject to the 
[EULA](https://github.com/caddyserver/caddy/blob/545fa844bbd188c1e5bff6926e5c410e695571a0/dist/EULA.txt) for 
Caddy's official binary distributions. If you plan to use Caddy for commercial purposes you should 
run the official Caddy distribution. 
