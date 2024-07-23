# IoT Project

This project is an IoT data processing pipeline that collects, processes, and visualizes data using InfluxDB, Telegraf, Grafana, and custom Python and Go applications.

This project is the combination from [Huntabyte's tig-stack](https://github.com/huntabyte/tig-stack) and [this anomaly detection example](https://github.com/InfluxCommunity/tg-brew-anomaly).

## Project Structure
```
iot-project/
├── cmd/
│ └── metric-replayer/
│   └── main.go
│ └── restamp/
│   └── main.go
├── data/
│ └── temps-stamped.txt
├── telegraf/
│ └── telegraf.conf
├── processors/
│ ├── forecasting.py
│ └── forecasting2.py
├── .env
├── requirements.txt
├── Dockerfile.metric-replayer
├── Dockerfile.telegraf
├── docker-compose.yml
└── telegraf.conf
```

## Components

- **InfluxDB**: A time-series database to store IoT data.
- **Telegraf**: A server agent to collect and report metrics.
- **Grafana**: A visualization tool to create dashboards and graphs.
- **Go Application (`metric-replayer`)**: A custom Go application to process data.
- **Go Application (`restamp`)**: This program assigns new timestamps to the influx line protocol data in data/temps.
- **Python Scripts (`forecasting.py`, `forecasting2.py`)**: Scripts for forecasting and data analysis.

## Setup

### Prerequisites

- Docker
- Docker Compose
- Golang version 1.22 (1.20 is okay)
- Python version 3.11

### Environment Variables

Create a `.env` file in the root directory with the following content:

```env
DOCKER_INFLUXDB_INIT_MODE=setup
DOCKER_INFLUXDB_INIT_USERNAME=aqua_user_db
DOCKER_INFLUXDB_INIT_PASSWORD=<secret_pass>
DOCKER_INFLUXDB_INIT_ADMIN_TOKEN=<secret_token>
DOCKER_INFLUXDB_INIT_ORG=aqua_org
DOCKER_INFLUXDB_INIT_BUCKET=<make_it_like_brew>
INFLUX_TOKEN=<secret_token>
INFLUX_HOST=localhost
DOCKER_INFLUXDB_INIT_RETENTION=4d
DOCKER_INFLUXDB_INIT_PORT=8086
DOCKER_INFLUXDB_INIT_HOST=influxdb
TELEGRAF_CFG_PATH=./telegraf.conf
GRAFANA_PORT=3000
```

## Build and Run

### 1.Clone the Repository


```shell
git clone https://github.com/ponkster/iot-project-TIG-stack.git
cd iot-project-TIG-stack
```

### 2. Build `restamp` binary

```shell
go build ./cmd/restamp
./restamp --help # should see usage
```

### 3. Assign data to the current timestamps.

This will amend the temperature data with current timestamps and store it to a new file.
You may want to repeat this step later to replay the data with new, current, timestamps.

```shell
    gunzip --stdout ./data/temps.txt.gz | ./restamp -filename - > ./data/temps-stamped.txt
```

You can verify it worked with this command:
```shell
   head -n3 ./data/temps-stamped.txt
```
You should see 3 lines of temperature data that look something like this

    temperature,brew=haze_v5 temperature=18.8 1596602228666886000
    temperature,brew=haze_v5 temperature=18.8 1596602229666886000
    temperature,brew=haze_v5 temperature=18.9 1596602230666886000

### 4. Build the application

If you want to inspect and monitor what is going on, run the script below:

```shell
docker compose up --build
```

Or you want to run it transparently

```shell
    docker compose up -d
```

### 5. Access Services:

   - **InfluxDB**:http://localhost:8086
   - **Grafana**:http://localhost:3000 (default credentials: `admin` / `admin`)

## Licenses
This project is licensed under the MIT License. See the LICENSE file for more details.
