# To build the image

Run the following command from root directory of the project
```
docker build -t notes-api-server .
```

# To push the image to docker hub

First tag the image with repository name
```
docker tag notes-api-server vinitondocker/notes-api-server:latest
```

Then push the image to docker hub 
```
docker push vinitondocker/notes-api-server:latest
```

# To run the api server
You can use the sample docker-compose.yml file

```
docker-compose up -d
```

Or you can set the required environment variables and run the following command
```
go run ./cmd/main.go
```

# Notes
Postman collection is provided in the ./postman_collection folder. The json file is to be imported in postman application to test the server
