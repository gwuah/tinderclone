import os
import time
from psycopg2 import connect, DatabaseError
from dotenv import load_dotenv

load_dotenv()

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
        os.system('cmd /k "goose -dir ./internal/migrations postgres "user={} password={} dbname={} sslmode=disable" {}"'.format(os.getenv("DB_USER"), os.getenv("DB_PASS"), os.getenv("DB_NAME"), cmd))
    except BaseException as error:
        print("Failed to run goose binary.")

def main():
    print("You can use 'up', 'down', 'status' and other goose commands to make migration changes")
    cmd = input("> ")
    time.sleep(3)

    connectGoose(cmd)

if __name__ == "__main__":
    connectDB()
    print("\n You can now attempt database migrations")
    main()
    connectDB().cursor.close()
    # # introduce waitkey and then     finally:
    #     if conn is not None:
    #         conn.close()
    #         print('Database connection closed.')

