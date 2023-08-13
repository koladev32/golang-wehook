from datetime import datetime
import json
import os
import random
import uuid
from flask import Flask
import redis


def get_payment(): 
    return {
    'url': os.getenv("WEBHOOK_ADDRESS", ""),
    'webhookId': uuid.uuid4().hex,
    'data': {
        'id': uuid.uuid4().hex,
        'payment': f"PY-{''.join((random.choice('abcdxyzpqr').capitalize() for i in range(5)))}",
        'event': random.choice(["accepted", "completed", "canceled"]),
        'created': datetime.now().strftime("%d/%m/%Y, %H:%M:%S"),
    }
}

redis_address = os.getenv("REDIS_ADDRESS", "")
host, port = redis_address.split(":")
port = int(port)
# Create a connection to the Redis server
redis_connection = redis.StrictRedis(host=host, port=port)


app = Flask(__name__)

@app.route('/payment')
def payment():
    webhook_payload_json = json.dumps(get_payment())

    # Publish the JSON string to the "payments" channel in Redis
    redis_connection.publish('payments', webhook_payload_json)
    
    return webhook_payload_json

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8000)
