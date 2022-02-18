import os
import time
from psycopg2 import connect, DatabaseError
from dotenv import load_dotenv
import argparse

load_dotenv()
parser = argparse.ArgumentParser()
parser.add_argument("-cmd", "--command", help="Goose command")
args = parser.parse_args()
cmd = args.command

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

def main():
    print("Running command...")
    connectGoose(cmd)

if __name__ == "__main__":
    connectDB()
    print("\n You can now attempt database migrations")
    main()
    connectDB().cursor.close()