import os
from dotenv import load_dotenv
import sys

load_dotenv()
if len(sys.argv) > 3:
    print("too many arguments passed")
    sys.exit()
if len(sys.argv) < 1:
    print("no argument passed")

def runMigration(flag, cmd=None, name=None):
    if flag == "-c" or flag == "-test":
        try:
            if flag == "-c" and cmd != None:
                print("creating new migrations...")
                os.system(f'goose -dir ./internal/migrations create {cmd} sql')
            elif flag == "-test":
                os.system(f'goose -dir ./internal/migrations postgres "user={os.getenv("DB_USER")} password={os.getenv("DB_PASS")} dbname=tinderclone_test sslmode=disable" {cmd} {name}')
        
        except BaseException as error:
            print("failed to run goose binary")
    else:   
        try:
            print("running sweet migrations...")
            os.system(f'goose -dir ./internal/migrations postgres "user={os.getenv("DB_USER")} password={os.getenv("DB_PASS")} dbname={os.getenv("DB_NAME")} sslmode=disable" {flag} {cmd}')        
        except BaseException as error:
            print("failed to run goose binary")  


if __name__ == "__main__":
    print("building...")
    runMigration(*sys.argv[1:])
