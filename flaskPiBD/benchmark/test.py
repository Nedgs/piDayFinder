import timeit


def run_benchmark():
    setup_code = '''
from flask import Flask, request
import time
import re

app = Flask(__name__)

def search_sequence(sequence, pi_decimals):
    match = re.search(sequence, pi_decimals)
    if match:
        position = match.start()
        return position
    else:
        return None

@app.route('/', methods=['POST'])
def index():
    sequence = request.form['sequence']
    pi_filename = '../pi.txt'

    with open(pi_filename, 'r') as file:
        pi_decimals = file.read()

    position = search_sequence(sequence, pi_decimals)

    return 'OK'
'''

    # Mesurer le temps d'exécution
    execution_time = timeit.timeit(stmt="requests.post('http://127.0.0.1:5000/', data={'sequence': '2308'})",
                                   setup="import requests", number=100)

    print(f"Temps d'exécution moyen sur 100 requêtes: {execution_time / 100} secondes")


if __name__ == '__main__':
    run_benchmark()
