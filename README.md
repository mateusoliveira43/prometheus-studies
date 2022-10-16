# Prometheus studies :fire::pencil:

Studies to understand how Prometheus works.

## About

TODO

### Alert manager

TODO

## Running

We will follow [Prometheus documentation](https://prometheus.io/docs/introduction/overview/) examples using Docker. We will use [Prometheus](https://hub.docker.com/u/prom) images, but is also possible to use [Bitnami](https://hub.docker.com/u/bitnami) images.

### Requirements

To run the examples, it is necessary the following tools:

- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)


### Following [Prometheus first steps](https://prometheus.io/docs/introduction/first_steps/)

To see Prometheus help, run
```
docker run --rm prom/prometheus:latest --help
```

To see Prometheus sample configuration file, run
```
docker run --rm --entrypoint cat prom/prometheus:latest /etc/prometheus/prometheus.yml
```

To start Prometheus with our `prometheus.yml` file, run
```
docker run --rm -p 9090:9090 \
    -v "$(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml" \
    prom/prometheus:latest
```
Access [http://localhost:9090](http://localhost:9090) to see it. Run `CTRL+C` to stop it.

### Following [Prometheus getting started](https://prometheus.io/docs/prometheus/latest/getting_started/)

To follow this section, change to the following directory
```sh
cd getting-started
```

To start Prometheus with our `prometheus.yml` file, run
```
docker run --rm -p 9090:9090 \
    -v "$(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml" \
    prom/prometheus:latest
```
Access [http://localhost:9090](http://localhost:9090) to see it. Run `CTRL+C` to stop it.

To follow the section with Node Exporter, change to the following directory
```sh
cd node-exporter
```

To start Prometheus with sample targets, run
```
docker-compose up
```
Access [http://localhost:9090](http://localhost:9090) to see it. Run `CTRL+C` to stop it.

To clean up the workspace, run
```
docker-compose down --volumes --rmi 'all'
```

### Grafana

To follow this section, change to the following directory
```sh
cd grafana
```

To start Prometheus and Grafana, run
```
docker-compose up
```
Access [http://localhost:9090](http://localhost:9090) to see Prometheus and [http://localhost:3000](http://localhost:3000) to see Grafana. Run `CTRL+C` to stop it.

Follow this [Grafana tutorial](https://prometheus.io/docs/tutorials/visualizing_metrics_using_grafana/).

To clean up the workspace, run
```
docker-compose down --volumes --rmi 'all'
```

### TODO Alerting

https://prometheus.io/docs/tutorials/alerting_based_on_metrics/

Run without webhook site

### TODO Unit tests

https://prometheus.io/docs/prometheus/latest/configuration/unit_testing_rules/

## References :books:

- https://prometheus.io/docs/introduction/overview/
