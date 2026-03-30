import os
import sqlite3
from flask import Flask, request

app = Flask(__name__)

@app.route('/user')
def get_user():
    user_id = request.args.get('id')
    conn = sqlite3.connect('test.db')
    # SQL Injection
    result = conn.execute("SELECT * FROM users WHERE id = " + user_id)
    return str(result.fetchall())

@app.route('/ping')
def ping():
    host = request.args.get('host')
    # Command Injection
    output = os.popen("ping -c 4 " + host).read()
    return output

@app.route('/hello')
def hello():
    name = request.args.get('name')
    # XSS
    return "<h1>Hello " + name + "</h1>"

if __name__ == '__main__':
    app.run(debug=True)
# PR Gate break build test Mon Mar 30 11:26:18 AM CST 2026
