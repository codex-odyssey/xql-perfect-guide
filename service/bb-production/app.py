import os, time, random
from flask import Flask, request
from database import get_bb_rating_from_db, get_bb_rating_from_cache, set_bb_rating_to_cache
from logger import setup_logger

app = Flask(__name__)
logger = setup_logger()

# BB Corp!!!

# rate
@app.route('/rate', methods=['GET'])
def main():
        # Get Data
        dish_name = request.args.get('dish_name')
        logger.info(f"リクエスト受信: {dish_name}")

        # Query Cache ( Memcached )
        cache = get_bb_rating_from_cache(dish_name)
        if cache != None:
            bb_rating = cache
        else:
            # Query DB ( MySQL ).
            bb_rating = get_bb_rating_from_db(dish_name)

            # Set Cache ( Memcache )
            set_bb_rating_to_cache(dish_name, bb_rating)


        return bb_rating

if __name__ == '__main__':
    host, port = os.getenv('BB_CORP_HOST', '0.0.0.0'), os.getenv('BB_CORP_PORT', 8091)
    app.run(host=host, port=port)
