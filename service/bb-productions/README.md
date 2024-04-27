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

export BB_PROD_HOST="0.0.0.0"
export BB_PROD_PORT=8091

export BB_PROD_DB_USER="bb"
export BB_PROD_DB_PASSWORD="password"
export BB_PROD_DB_HOST="127.0.0.1"
export BB_PROD_DB_NAME="bb"
export BB_PROD_DB_PORT=3306

export OTEL_SERVICE_NAME=BB-PROD
export OTEL_TRACES_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_PROTOCOL=grpc
export OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
export OTEL_EXPORTER_OTLP_INSECURE=True
```
