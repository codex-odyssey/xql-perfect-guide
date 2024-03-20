import os
import mysql.connector
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
    
    logger.info(config)
    
    try: 
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