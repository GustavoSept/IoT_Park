from iotPark import app
from flask import render_template
import requests
import logging

@app.route('/')
def pg_home():

    isAuthorized = False # TODO: implement JWT check later, do it in a separate function

    if not isAuthorized:
        return render_template('login.html')

    users_url = 'http://go-backend:8080/api/get_all_users'
    usersAuth_url = 'http://go-backend:8080/api/get_all_usersAuth'
    parking_lots_url = 'http://go-backend:8080/api/get_all_pLots'

    try:
        users_response = requests.get(users_url)
        usersAuth_response = requests.get(usersAuth_url)
        parking_lots_response = requests.get(parking_lots_url)

        users = users_response.json() if users_response.status_code == 200 else []
        users_auth = usersAuth_response.json() if usersAuth_response.status_code == 200 else []
        parking_lots = parking_lots_response.json() if parking_lots_response.status_code == 200 else []

        return render_template('home.html', users=users, parking_lots=parking_lots, users_auth=users_auth)

    except requests.exceptions.RequestException as e:
        logging.error(f"Error connecting to backend: {e}")
        return render_template('home.html', users=[], parking_lots=[], users_auth=[])
    
@app.route('/login')
def pg_login():
    return render_template('login.html')

@app.route('/signup')
def pg_signup():
    return render_template('signup.html')

