from flask import Flask, render_template

app = Flask(__name__)

# needs to come after app instance is created
from frontend.iotPark import routes