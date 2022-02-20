import os
import time
from dotenv import load_dotenv
import sys

load_dotenv()
if len(sys.argv) > 2:
    print("Too many arguments passed")
    sys.exit()
cmd = str(sys.argv[1])

def runMigration(cmd):
    try:
        print("Running sweet migrations...")
        time.sleep(3)
        os.system('cmd /c "goose -dir ./internal/migrations postgres "user={} password={} dbname={} sslmode=disable" {}"'.format(os.getenv("DB_USER"), os.getenv("DB_PASS"), os.getenv("DB_NAME"), cmd))
    except BaseException as error:
        print("Failed to run goose binary.")


if __name__ == "__main__":
    print("building...")
    runMigration(cmd)