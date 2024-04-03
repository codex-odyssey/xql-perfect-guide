import os
import mysql.connector
from pymemcache.client import base
from logger import setup_logger

logger = setup_logger()

# Get BBB Evaluation from MySQL.
def get_bbb_evaluation_from_db(dish_name): 
    config = {
        'user': os.getenv('BBB_SERVICE_DB_USER', 'bbb'),
        'password': os.getenv('BBB_SERVICE_DB_PASSWORD', 'password'),
        'host': os.getenv('BBB_SERVICE_DB_HOST', '127.0.0.1'),
        'database': os.getenv('BBB_SERVICE_DB_NAME', 'bbb'),
        'port': int(os.getenv('BBB_SERVICE_DB_PORT', '3306')),
        'raise_on_warnings': True,
    }
    
    # TODO こいつはそのうち消すログ
    logger.info(config)
    
    try: 
        if dish_name == "yasaiitame":
            raise ValueError(f"無効な料理名が指定されました: {dish_name}")
        
        cnx = mysql.connector.connect(**config)
        cursor = cnx.cursor()

        query = "SELECT bbb_evaluation FROM bbb_evaluation WHERE dish_name = %s"
        cursor.execute(query, (dish_name,))    
        
        result = cursor.fetchone()
        
        bbb_evaluation = str(result[0]) if result else "Not Found"
        logger.info(f"DB からデータ取得: {dish_name} の BBB 流評価 {bbb_evaluation} ")
    except Exception as e:
        logger.error(f"Error: {e}")
    finally:
        cursor.close()
        cnx.close()
        
    return bbb_evaluation

# Get BBB Evaluation from Memchached.
def get_bbb_evaluation_from_cache(key):
    memcache_host = os.getenv('BBB_SERVICE_MEMCACHE_HOST', 'localhost')
    memcache_port = int(os.getenv('BBB_SERVICE_MEMCACHE_PORT', 11211))
    
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

# Set BBB Evaluation from Memchached ( default expire time = 60 Sec. ).
def set_bbb_evaluation_to_cache(dish_name, bbb_evaluation):
    memcache_host = os.getenv('BBB_SERVICE_MEMCACHE_HOST', 'localhost')
    memcache_port = int(os.getenv('BBB_SERVICE_MEMCACHE_PORT', 11211))
    expire_time = int(os.getenv('BBB_SERVICE_MEMCACHE_EXPIRE_TIME', 60))

    client = base.Client((memcache_host, memcache_port))
    client.set(dish_name, bbb_evaluation, expire=expire_time)
    logger.info(f"Cache 設定 ( TTL: {expire_time} 秒 ): {dish_name}:{bbb_evaluation}")