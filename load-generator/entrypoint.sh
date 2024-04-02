#!/bin/sh
locust -f /usr/src/app/main.py --headless --host ${LOCUST_HOST:-"http://waiter-service:8080"}