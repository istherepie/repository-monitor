# Request Echo

*I needed a tool to view incoming HTTP requests from a 3rd party integration...*

View incoming request information such as,

* the uri path
* query params
* request data



## Installation

Run tests as follows:

```
	make test
```

Building as binary:

```
	make build
```

Alternatively you may build a docker container using the supplied file:

```
	docker build -t <namespace/image>:<tag> .
```

or use the latest available container from [docker hub](https://hub.docker.com/repository/docker/istherepie/request-monitor):

```
	docker pull istherepie/request-monitor:latest
	docker run --rm -p8080:8080 istherepie/request-monitor:latest
```



## Usage

The application takes the following flags:

* -host (string) - Service hostname, defaults to `localhost`.
* -port (int) - Service port, defaults to `8080`.
* -id (string) - Service ID/Metaname for the purpose of running multiple instances.

**NOTE:**
The service port can be overridden by setting the env variable `RM_SERVICE_PORT`,
this is useful when using task runners/docker containers to start the service.



## License

Distributed under the MIT License. See `LICENSE` for more information.



## Contact

Steffen Park dev@istherepie.com