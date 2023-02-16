# Go-Bank-V2
A simple GO Bank API v2



## .env file
```
POSTGRES_CONN_STRING = "user=username dbname=databaseName password=password sslmode=disable"

DB_NAME = "databaseName"
DB_USERNAME = "username"
DB_PASSWORD = "password"

M_DB_NAME = databaseName
M_DB_USERNAME = username
M_DB_PASSWORD = password
```
    - Here in env file, databaseName, username and password for variables must be same.
    - POSTGRES_CONN_STRING is a connection string for golang 
    - DB_NAME, DB_USERNAME, DB_PASSWORD are used in docker-compose file
    - M_DB_NAME, M_DB_USERNAME AND M_DB_PASSWORD are used in Makefile