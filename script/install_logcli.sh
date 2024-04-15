mkdir -p bin
wget -O bin/logcli-linux-amd64.zip https://github.com/grafana/loki/releases/download/v3.0.0/logcli-linux-amd64.zip
unzip -d bin bin/logcli-linux-amd64.zip
mv bin/logcli-linux-amd64 bin/logcli
