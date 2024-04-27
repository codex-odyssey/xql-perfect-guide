# Pythonのベースイメージを使用
FROM python:3.11

# 作業ディレクトリを設定
WORKDIR /usr/src/app

# 必要なライブラリをインストール
COPY requirements.txt ./
RUN pip install --no-cache-dir -r requirements.txt
RUN opentelemetry-bootstrap -a install

# アプリケーションのコードをコピー
COPY . .

# Flaskアプリケーションを実行
CMD ["opentelemetry-instrument", "python", "app.py"]