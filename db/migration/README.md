# Create migrations using golang-migrate

## create an initial migration
    $ migrate create -ext sql -dir db/migration -seq init_schema

     - migrate -> command to create migrations
     - create -> new migration
     - ext -> extension
     - dir -> directory path where our migrations are stored
     - seq -> squential up and down for migrations with N numbers

    - schema with up means to migrate up
    - schema with down means to revert changes

## migrate command to migrate 
    $ migrate -path db/migration -database "postgresql://root:root@goBankV2:5432/goBankApiV2" -verbose