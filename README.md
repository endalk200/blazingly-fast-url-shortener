# Blazingly Fast URL Shortener

URL shortener built using go and redis as storage layer.

## Setup and Test

To run and use the server you have to install redis and go. Your redis instance should be up and running on the default port.

Run all integration and unit tests

```bash
go test -v ./...
```

Run the server

```bash
go run main.go
```

## API

### Create a short URL

```bash
curl --request POST \
--data '{
    "long_url": "https://github.com/endalk200/blazingly-fast-url-shortner/blob/main/.gitignore",
    "user_id" : "e0dba740-fc4b-4977-872c-d360239e6b10"
}' \
  http://localhost:9808/create-short-url
```

The above endpoint should return a response like this

```json
{
  "message": "short url created successfully",
  "short_url": "http://localhost:9808/9Zatkhpi"
}
```

### Get a short URL

```bash
curl --request GET \
  --url 'http://localhost:9808/get-short-url?short_url=1'
```
