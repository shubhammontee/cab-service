Migration :

1. create database cabservice

2. Run the migration using the following command


* goose mysql "username:password@/cabservice?parseTime=true" up


Run Application: Make run

Build Application: Make build
* this generate a bin directory at project root and the binary gets stored there


Generate swagger docs: Make swagger


Get Swagger UI: Make serve_swagger

*this endpoint can be used for sending request and receiving response
