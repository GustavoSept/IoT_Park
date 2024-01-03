from iotPark import app
from flask import render_template
import requests
import logging

@app.route('/')
def pg_home():
    users_url = 'http://go-backend:8080/api/get_all_users'
    parking_lots_url = 'http://go-backend:8080/api/get_all_pLots'

    try:
        users_response = requests.get(users_url)
        parking_lots_response = requests.get(parking_lots_url)

        users = users_response.json() if users_response.status_code == 200 else []
        parking_lots = parking_lots_response.json() if parking_lots_response.status_code == 200 else []

        return render_template('home.html', users=users, parking_lots=parking_lots)

    except requests.exceptions.RequestException as e:
        logging.error(f"Error connecting to backend: {e}")
        return render_template('home.html', users=[], parking_lots=[])