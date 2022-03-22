import os
from dotenv import load_dotenv
import sys

load_dotenv()
if len(sys.argv) > 4:
    print("too many arguments passed")
    sys.exit()
if len(sys.argv) < 1:
    print("no argument passed")

def runMigration(cmd, type=None, format=None):
    try:
        print("running sweet migrations...")
        os.system(f'cmd /c "goose -dir ./internal/migrations postgres "user={os.getenv("DB_USER")} password={ os.getenv("DB_PASS")} dbname={os.getenv("DB_NAME")} sslmode=disable" {cmd} {format} {type}"')
    except BaseException as error:
        print("failed to run goose binary")


if __name__ == "__main__":
    print("building...")
    runMigration(*sys.argv[1:])