# To build the image
RUN
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

# To run the server
You can use the sample docker-compose.yml file

```
docker-compose up -d
```

Or you can set the required environment variables and run the following command
```
go run ./cmd/main.go
```


