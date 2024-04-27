## Environment
```sh
$ source ./bin/activate

# install packege
$ pip install "flask < 3"
$ pip install mysql-connector-python==8.2.0
$ pip install opentelemetry-distro==0.41b0
$ opentelemetry-bootstrap -a install
$ pip install opentelemetry-sdk==v1.20.0 opentelemetry-exporter-otlp==v1.20.0
$ pip install opentelemetry-exporter-otlp==v1.20.0

export BB_CORP_HOST="0.0.0.0"
export BB_CORP_PORT=8091

export BB_CORP_DB_USER="bb"
export BB_CORP_DB_PASSWORD="password"
export BB_CORP_DB_HOST="127.0.0.1"
export BB_CORP_DB_NAME="bb"
export BB_CORP_DB_PORT=3306

export OTEL_SERVICE_NAME=BB-CORP
export OTEL_TRACES_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
export OTEL_EXPORTER_OTLP_INSECURE=True
```
