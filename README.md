# IoT Project

This project is an IoT data processing pipeline that collects, processes, and visualizes data using InfluxDB, Telegraf, Grafana, and custom Python and Go applications.

## Project Structure

iot-project/
├── cmd/
│ └── metric-replayer/
│ └── main.go
├── data/
│ └── temps-stamped.txt
├── processors/
│ ├── forecasting.py
│ └── forecasting2.py
├── requirements.txt
├── Dockerfile.metric-replayer
├── Dockerfile.telegraf
├── docker-compose.yml
└── telegraf2.conf


## Components

- **InfluxDB**: A time-series database to store IoT data.
- **Telegraf**: A server agent to collect and report metrics.
- **Grafana**: A visualization tool to create dashboards and graphs.
- **Go Application (`metric-replayer`)**: A custom Go application to process data.
- **Python Scripts (`forecasting.py`, `forecasting2.py`)**: Scripts for forecasting and data analysis.

## Setup

### Prerequisites

- Docker
- Docker Compose

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
DOCKER_INFLUXDB_INIT_MODE=setup
DOCKER_INFLUXDB_INIT_USERNAME=aqua_user_db
DOCKER_INFLUXDB_INIT_PASSWORD=donletmidonletmidon_memelusin_memelo
DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=Th1$!sth3T0k3n^
DOCKER_INFLUXDB_INIT_ORG=aqua_org
DOCKER_INFLUXDB_INIT_BUCKET=aqua_sensor_bucket
INFLUX_TOKEN=qsgWkgT3CwI4lyQaLh0aHDhSdi2zc4xhvNGHtBzu-1jOGSaEBxqQ9DCqDvui4EKgw_vs-5r5s1_RWrPwG14gDQ==
INFLUX_HOST=localhost
DOCKER_INFLUXDB_INIT_RETENTION=4d
DOCKER_INFLUXDB_INIT_PORT=8086
DOCKER_INFLUXDB_INIT_HOST=influxdb
TELEGRAF_CFG_PATH=./telegraf2.conf
GRAFANA_PORT=3000
```

## Build and Run

### 1.Clone the Repository


```shell
git clone <repository-url>
cd iot-project
```

### 2. Build and start the service

```shell
docker-compose up --build
```
