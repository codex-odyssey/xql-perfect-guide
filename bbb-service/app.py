import os, time, random
from flask import Flask, request
from database import get_bbb_evaluation_from_db
from logger import setup_logger

app = Flask(__name__)
logger = setup_logger()

# BBB!!!
@app.route('/bbb', methods=['GET'])
def main():  
        # Get Data
        dish_name = request.args.get('dish_name')
        logger.info(f"リクエスト受信: {dish_name}")
        
        # Query DB ( MySQL ).
        cooking_time = get_bbb_evaluation_from_db(dish_name)

        return cooking_time

if __name__ == '__main__':
    host, port = os.getenv('BBB_SERVICE_HOST', '0.0.0.0'), os.getenv('BBB_SERIVCE_PORT', 8091)
    app.run(host=host, port=port)