from iotPark import app
from flask import render_template, request, redirect, url_for
import requests
import logging

def has_auth_token():
    return 'AuthToken' in request.cookies and 'RefreshToken' in request.cookies

@app.route('/')
@app.route('/dashboard', methods=["GET", "POST"])
def pg_home():
    if not has_auth_token():
        logging.warning("User does not have either the AuthToken or RefreshToken")
        return redirect(url_for('pg_login'))

    users_url = 'http://go-backend:8080/api/get_all_users'
    usersAuth_url = 'http://go-backend:8080/api/get_all_usersAuth'
    parking_lots_url = 'http://go-backend:8080/api/get_all_pLots'

    try:
        headers = {
            'Cookie': f'AuthToken={request.cookies["AuthToken"]}; RefreshToken={request.cookies["RefreshToken"]}'
        }

        logging.info(f"Header: {headers}")

        users_response = requests.get(users_url, headers=headers)
        usersAuth_response = requests.get(usersAuth_url, headers=headers)
        parking_lots_response = requests.get(parking_lots_url, headers=headers)

        if users_response.status_code == 401 or usersAuth_response.status_code == 401 or parking_lots_response.status_code == 401:
            return redirect(url_for('pg_login'))

        users = users_response.json() if users_response.status_code == 200 else []
        users_auth = usersAuth_response.json() if usersAuth_response.status_code == 200 else []
        parking_lots = parking_lots_response.json() if parking_lots_response.status_code == 200 else []

        return render_template('dashboard.html', users=users, parking_lots=parking_lots, users_auth=users_auth)

    except requests.exceptions.RequestException as e:
        logging.error(f"Error connecting to backend: {e}")
        return render_template('dashboard.html', users=[], parking_lots=[], users_auth=[])

@app.route('/login', methods=["POST"])
def pg_login():
    return render_template('login.html')

@app.route('/signup', methods=["POST"])
def pg_signup():
    return render_template('signup.html')