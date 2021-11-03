# Install Go and Go tool into VS Code

# Install Redis
https://github.com/microsoftarchive/redis/releases/download/win-3.0.504/Redis-x64-3.0.504.msi

# Install Postgres
https://content-www.enterprisedb.com/postgresql-tutorial-resources-training?cid=924

# Create the database
library

# Into the library database, run the script
Go to Schemas and create one using --> CreateSchema.sql

# Restore backup
Go to university schema and restore the backup using --> u-library.sql

# Execute the next command into University Schema into Library database
CREATE EXTENSION pgcrypto;

# restore the postman collection from 
Collection.json

# Open the app and open a new terminal and change directory to u-library and execute
go run .

# Unit testing book_test.go

Move to .\pkg\handler
execute go test .