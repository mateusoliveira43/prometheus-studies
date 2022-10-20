# Prometheus studies :fire::pencil:

Studies to understand how Prometheus works.

## About

[Prometheus](https://prometheus.io/docs/introduction/overview/) is an open-source monitoring solution. It collects and stores the target applications' metrics, which can be retrieved with PromQL (Prometheus Query Language). Metrics are stored in [time series database (TSDB)](https://grafana.com/docs/grafana/latest/fundamentals/timeseries/)  as time series data. It also handles alerts through [Alertmanager](https://prometheus.io/docs/introduction/glossary/#alertmanager).

### Visualization

To visualize the collected data, Prometheus Web UI can be used, or a more robust solution, like [Grafana](https://grafana.com/).

### Exporter

An exporter is a binary running alongside the target applications, which exposes their metrics, in a format Prometheus supports. For monitoring a wide variety of hardware and kernel-related metrics, [Prometheus Node Exporter](https://github.com/prometheus/node_exporter) can be used.

## Running

We will follow [Prometheus documentation](https://prometheus.io/docs/introduction/overview/) and [The Uncomplicating Prometheus Training](https://github.com/badtuxx/DescomplicandoPrometheus) examples using Docker. We will use [Prometheus](https://hub.docker.com/u/prom) images, but is also possible to use [Bitnami](https://hub.docker.com/u/bitnami) images.

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

### Following [Grafana](https://grafana.com/docs/grafana/latest/) tutorials

To follow this section, change to the following directory
```sh
cd grafana
```

To see Grafana sample configuration file, run
```
docker run --rm --entrypoint cat grafana/grafana:latest /etc/grafana/grafana.ini
```

To start Prometheus and Grafana, run
```
docker-compose up
```
Access [http://localhost:9090](http://localhost:9090) to see Prometheus and [http://localhost:3000](http://localhost:3000) to see Grafana. Run `CTRL+C` to stop it.

Follow this [Prometheus Grafana tutorial](https://prometheus.io/docs/tutorials/visualizing_metrics_using_grafana/), or this [Grafana Build your first dashboard](https://grafana.com/docs/grafana/latest/getting-started/build-first-dashboard/).

To clean up the workspace, run
```
docker-compose down --volumes --rmi 'all'
```

TODO add provision example https://grafana.com/tutorials/provision-dashboards-and-data-sources/

### Creating an exporter

We can write an exporter in several languages, like [Python](https://github.com/prometheus/client_python), [Go](https://github.com/prometheus/client_golang), etc. To check some already existing solutions, check [Prometheus client libraries](https://prometheus.io/docs/instrumenting/clientlibs/).

To follow this section, change to the following directory
```sh
cd creating-exporter
```

To start Prometheus and the Go exporter, run
```
docker-compose up
```
Access [http://localhost:9090](http://localhost:9090) to see Prometheus and [http://localhost:8081](http://localhost:8081) to see Go exporter. Run `CTRL+C` to stop it.

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
- https://github.com/badtuxx/DescomplicandoPrometheus
