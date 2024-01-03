from flask import Flask
import logging

app = Flask(__name__)

# Configure logging
logging.basicConfig(level=logging.INFO)

# needs to come after app instance is created
from iotPark import routes