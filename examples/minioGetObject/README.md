## MinIO getObject

Get the object from an existing bucket by its name.

## Environment

Copy/Rename .en.example to .env<br>
Change the values of the variables to your need in .env file.<br>

## Run

Open your terminal in this current path, then run

```
go run getObject.go
```

## Methods

1- `NewMinioClient`

- Stablish a connection with MinIO in pkg using the variables from .env file.<br>

2- `GetObjectFromMinio`

- Get the content of the object from the MinIO bucket by its name.
