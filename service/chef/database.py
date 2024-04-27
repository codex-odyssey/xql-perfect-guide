import os
import mysql.connector
from logger import setup_logger

logger = setup_logger()

# Get Cooking Time from MySQL.
def get_cooking_time_from_db(dish_name):
    config = {
        'user': os.getenv('CHEF_SERVICE_DB_USER', 'chef'),
        'password': os.getenv('CHEF_SERVICE_DB_PASSWORD', 'password'),
        'host': os.getenv('CHEF_SERVICE_DB_HOST', '127.0.0.1'),
        'database': os.getenv('CHEF_SERVICE_DB_NAME', 'chef'),
        'port': int(os.getenv('CHEF_SERVICE_DB_PORT', '3306')),
        'raise_on_warnings': True,
    }

    logger.info(config)

    try:
        cnx = mysql.connector.connect(**config)
        cursor = cnx.cursor()

        query = "SELECT cooking_time FROM cooking_time WHERE dish_name = %s"
        cursor.execute(query, (dish_name,))

        result = cursor.fetchone()

        cooking_time = str(result[0]) if result else "Not Found"
        logger.info(f"DB からデータ取得: {dish_name} の調理時間 {cooking_time} [min.]")
    except Exception as e:
        logger.error(f"Error: {e}")
    finally:
        cursor.close()
        cnx.close()

    return cooking_time
