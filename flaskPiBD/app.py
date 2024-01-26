from flask import Flask, render_template, request

app = Flask(__name__)


def search_sequence(sequence, pi_decimals):
    return pi_decimals.find(sequence)


@app.route('/', methods=['GET', 'POST'])
def index():
    if request.method == 'POST':
        date_sequence = request.form['date_sequence']
        pi_file_path = 'pi.txt'

        # Charger les décimales de pi à partir du fichier
        with open(pi_file_path, 'r') as file:
            pi_decimals = file.read()

        position = search_sequence(date_sequence, pi_decimals)

        return render_template('result.html', date_sequence=date_sequence, position=position)

    return render_template('index.html')


if __name__ == '__main__':
    app.run(debug=True)
