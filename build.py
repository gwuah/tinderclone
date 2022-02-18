import os
import time
from psycopg2 import connect, DatabaseError
from dotenv import load_dotenv
import argparse
import sys

load_dotenv()
cmd = str(sys.argv[1])

def connectDB():
    try:
        print("Connecting to PostgreSQL database...")
        conn = connect (
        dbname = os.getenv("DB_NAME"), 
        user = os.getenv("DB_USER"), 
        host = os.getenv("DB_HOST"), 
        password = os.getenv("DB_PASS"))   

        cursor = conn.cursor()
        print("PostgreSQL database connected")

        return cursor

    except (Exception, DatabaseError) as error:
        print("Failed to connect to the POstgres database.")
        print(error)

def connectGoose(cmd):
    try:
        print("Making connection to goose...")
        time.sleep(3)
        os.system('cmd /c "goose -dir ./internal/migrations postgres "user={} password={} dbname={} sslmode=disable" {}"'.format(os.getenv("DB_USER"), os.getenv("DB_PASS"), os.getenv("DB_NAME"), cmd))
    except BaseException as error:
        print("Failed to run goose binary.")


if __name__ == "__main__":
    connectDB()
    print("building...")
    connectGoose(cmd)