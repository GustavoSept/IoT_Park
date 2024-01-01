from iotPark import app
from flask import render_template
import requests

@app.route('/')
def pg_home():    
    url = 'http://go-backend:8080/'

    try:
        response = requests.get(url)
        if response.status_code == 200:
            return render_template('home.html', users=response.json())
        else:
            return f'Error: {response.status_code}'

    except requests.exceptions.RequestException as e:
        print(f"Error connecting to backend: {e}")
        return "Error connecting to backend."