import os, time, random
from flask import Flask, request
from database import get_bbb_evaluation_from_db, get_bbb_evaluation_from_cache, set_bbb_evaluation_to_cache
from logger import setup_logger

app = Flask(__name__)
logger = setup_logger()

# BBB!!!
@app.route('/bbb', methods=['GET'])
def main():  
        # Get Data
        dish_name = request.args.get('dish_name')
        logger.info(f"リクエスト受信: {dish_name}")
        
        # Query Cache ( Memcached )
        cache = get_bbb_evaluation_from_cache(dish_name)
        if cache != None:
            bbb_evaluation = cache
        else: 
            # Query DB ( MySQL ).
            bbb_evaluation = get_bbb_evaluation_from_db(dish_name)
            
            # Set Cache ( Memcache )
            set_bbb_evaluation_to_cache(dish_name, bbb_evaluation)
            

        return bbb_evaluation

if __name__ == '__main__':
    host, port = os.getenv('BBB_SERVICE_HOST', '0.0.0.0'), os.getenv('BBB_SERIVCE_PORT', 8091)
    app.run(host=host, port=port)