from flask import Flask

app = Flask(__name__)

# needs to come after app instance is created
from iotPark import routes