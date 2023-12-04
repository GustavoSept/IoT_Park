from frontend.iotPark import app
from flask import render_template

@app.route('/')
def pg_home():
    return render_template('home.html')
