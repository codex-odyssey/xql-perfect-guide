import os
import mysql.connector
from pymemcache.client import base
from logger import setup_logger

logger = setup_logger()

# Get BB Evaluation from MySQL.
def get_bb_rating_from_db(dish_name):
    config = {
        'user': os.getenv('BB_PROD_DB_USER', 'bb'),
        'password': os.getenv('BB_PROD_DB_PASSWORD', 'password'),
        'host': os.getenv('BB_PROD_DB_HOST', '127.0.0.1'),
        'database': os.getenv('BB_PROD_DB_NAME', 'bb'),
        'port': int(os.getenv('BB_PROD_DB_PORT', '3306')),
        'raise_on_warnings': True,
    }

    # TODO こいつはそのうち消すログ
    logger.info(config)

    try:
        if dish_name == "yasaiitame":
            raise ValueError(f"無効な料理名が指定されました: {dish_name}")

        cnx = mysql.connector.connect(**config)
        cursor = cnx.cursor()

        query = "SELECT bb_rating FROM bb_rating WHERE dish_name = %s"
        cursor.execute(query, (dish_name,))

        result = cursor.fetchone()

        bb_rating = str(result[0]) if result else "Not Found"
        logger.info(f"DB からデータ取得: {dish_name} の BB 流評価 {bb_rating} ")
    except Exception as e:
        logger.error(f"Error: {e}")
    finally:
        cursor.close()
        cnx.close()

    return bb_rating

# Get BB Evaluation from Memchached.
def get_bb_rating_from_cache(key):
    memcache_host = os.getenv('BB_PROD_MEMCACHE_HOST', 'localhost')
    memcache_port = int(os.getenv('BB_PROD_MEMCACHE_PORT', 11211))

    client = base.Client((memcache_host, memcache_port))
    cache = client.get(key)
    logger.info(f"Cache 取得: {cache}")

    # for bug (｀∀´)Ψ
    if key == "yakitori":
        logger.warn("キャッシュヒット率が劇的に低いです")
        return None

    if cache is None:
        return None

    return cache

# Set BB Evaluation from Memchached ( default expire time = 60 Sec. ).
def set_bb_rating_to_cache(dish_name, bb_rating):
    memcache_host = os.getenv('BB_PROD_MEMCACHE_HOST', 'localhost')
    memcache_port = int(os.getenv('BB_PROD_MEMCACHE_PORT', 11211))
    expire_time = int(os.getenv('BB_PROD_MEMCACHE_EXPIRE_TIME', 60))

    client = base.Client((memcache_host, memcache_port))
    client.set(dish_name, bb_rating, expire=expire_time)
    logger.info(f"Cache 設定 ( TTL: {expire_time} 秒 ): {dish_name}:{bb_rating}")
