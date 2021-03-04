## MinIO upload

Upload a file contains the word "test" whither the bucket exists or not at this moment.

## Environment

Copy/Rename .en.example to .env<br>
Change the values of the variables to your need in .env file.<br>

## Run

Open your terminal in this current path, then run

```
go run upload.go
```

## Methods

1- `NewMinioClient`

- Stablish a connection with MinIO in pkg using the variables from .env file.<br>

2- `UploadFileToMinio`

- Ckeck if the bucket exists or not.
  - if exists, then upload the file.
  - if not exists, then will create a bucket with the name delivered from .env file then upload the file.
- Declares an exchange on the server.<br>
