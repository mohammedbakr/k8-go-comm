## MinIO getPresignedURL

Get getPresignedURL of the object from an existing bucket.

## Environment

Copy/Rename .en.example to .env<br>
Change the values of the variables to your need in .env file.<br>

## Run

Open your terminal in this current path, then run

```
go run getPresignedURL.go
```

## Methods

1- `NewMinioClient`

- Stablish a connection with MinIO in pkg using the variables from .env file.<br>

2- `GetPresignedURLForObject`

- Get Get signed url for the object
