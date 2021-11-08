# Install Go and Go tool into VS Code

# Install Redis
1. Download from https://github.com/microsoftarchive/redis/releases/download/win-3.0.504/Redis-x64-3.0.504.msi

# Install Postgres
1. Download from https://content-www.enterprisedb.com/postgresql-tutorial-resources-training?cid=924
2. Change the password into .env file, using the password that you set up.

# Create the database
library

# Create a Schema
1. Select the "library" database 
2. Open a new Query Tool
3. Open the script CreateSchema.sql located in .\local_install\db\ directory
4. Run the script.

This script is to create the university schema

# Restore the Database
1. Select the "university" schema
2. Rigth click on university and select "Restore.."
3. Load the sql file "u-library.sql" located in .\local_install\db\ directory
4. Run the script.

This script is to create the table structure

# Add a pgcrypto extension in Postgresql
1. Open a new Query Tool
2. Execute the next command CREATE EXTENSION pgcrypto;

This command eneables the extension pgcrypto.

# Postman set up
1. Restore the collection located in .\local_install\collections\
2. Follow the steps in Postman to restore the Collection.json

# Run API
1. Open the app in VS Code
2. Open a new terminal
3. Make sure that you are in the directory .\u-library
4. Execute the command go run .

# Run Unit Testing
1. Open the app in VS Code
2. Open a new terminal
3. Make sure that you are in the directory .\u-library
4. Execute the comman go test -v <package name>
    i.e go test -v .\test\pkg\handler