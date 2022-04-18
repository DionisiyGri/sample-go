# **Sample project for base Golang projects.**

**This sample uses very little external packages.**

**For db** (postgres in this example) - built-in https://golang.org/pkg/database/sql/. You can use anything else (e.g. gorm), but I choose  database/sql just to simplify everything (especially migrations, imho gorm`s migration tool is painfull :)).

**Web framework** - https://github.com/gorilla/mux. Yeah, there are different frameworks with its own pros and cons. I think it won`t be a problem to fix couple of files and use what you want.

**For logs** - https://github.com/sirupsen/logrus. It is quite simple and pretty. It has hooks to send data to ElasticSearch, add additional fields to log result etc.

### Note about import paths
To rename import paths you can run below command (just an example):
`git grep -l -r 'github.com/DionisiyGri/sample-go' | xargs sed -i '' -e 's/sample-go/some-new-name/g' `, but you feel free to rename paths in your way (IDE`s find and replace feature is one of the possible solutions).

### Build
Execute ```make build``` to build binary file and ```./app``` to run it on port ```:8000```. Also it will copy config.json.tmpl file to root and rename it to config.json.

### Migrations
You can inspect Makefile, there are some scripts for your comfort.
up and down are self-explanatory.
If you wanna create new migration file you can run **make create-migration name=name_of_migration_file**.

### Endpoints
**GET */sample/health*** - check if connection to db is succeed and app is healthy (you should extend it if more external dependencies will appear).

**POST */sample/all*** - creates new record in db. You can pass ```{"name":"some_name"}``` in request body to create new record. Yeap, db schema is the simpliest one. Also note, that validation is used there (checks for empty name and checks for name length).


### Structure
Structure is quite simple.
1. **cmd/main.go** - it is entrypoint for the app. All initializations, parsing config happens here.
2. **config** - here you can store config for different envs, now just simple template for local development is there. You should copy it to the root with name *config.json* and populate it with needed values.
3. **internal** - all app internals :)

    3.1.  **config/config.go** - logic to parse config.json, extend for your needs.
    
    3.2. ***db*** - database layer, where all operations with db live.
    - ***db/migrations*** - files for migration and migration logic. You can add other sources like sqlite, mysql and everything should work.
    - ***db/postgres*** - implementation of needed functions to operate with postgres.
    - ***db/repos*** - you can add additional db sources here.
    - ***db/manager.go*** - interfaces which need to be implemented.
    - ***db/db.go*** - initialization functions for db connections.
   
   3.3. **handler** - self-explanatory :) all handlers will live here.
    - ***handler.go*** - struct, constructor for handler. Also you can see some helper functions to build pretty responses.
   - ***health.go*** and some.go - simple implemantation of endpoints.
  
    3.4. **model** - entities of your app (mb rename model -> entities? anyway, it is just sample structure, so you can rename it as you wish, but do not forget about import paths).
        
    3.5. **router** - initialization of router. Note that I pass PathPrefix which is service_name from config. It is not neccessary, but I think it is cozily if you have some amount of services to differ them. Ah, forget it, let it be for readability of api endpoints ;)
    
    3.6. **sample** - the place where logic should live, where we operate with parsed request body, make calculations and call db. 
    
    3.7. **server** - init our server.

4. **validation** - package which extend go-validator and you can add custom validation rules, it has comments, so I think everything should be clear.
        