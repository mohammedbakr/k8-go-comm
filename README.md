# k8-go-comm

Library to communicate with RabbitMQ and MinIO

## Goals

Provide helper modules which will be used by all parties willing to implement communications with RabbitMQ or Minio. All common types and functions needs to be placed here.

## Usage

See the 'examples' subdirectory with example for using these modules. Each usecase is in a specific subdirectory inside examples.<br>
For quick start using docker to run containers for RabbitMQ and MinIO<br>
Run Standalone MinIO on Docker.

```
docker run -p 9000:9000 \
  -e "MINIO_ROOT_USER=AKIAIOSFODNN7EXAMPLE" \
  -e "MINIO_ROOT_PASSWORD=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY" \
  minio/minio server /data

```

Run RabbitMQ on Docker.

```
docker run -d --hostname <your_host_name> --name <your_container_name> -p 15672:15672 -p 5672:5672 rabbitmq:3-management
```

## RabbitMQ

### Producer

To start a connectoin with RabbitMQ and publish a message, please check [producer folder](examples/producer)

### Consumer

To start a connection and consume the message that was published, please check [consumer directory](examples/consumer)

## MinIO

### Make bucket and upload file

To start a connection and upload a file, please check [minioUploead directory](examples/minioUploead)

### Get the object from the bucket

To start a connection and get a file from specific bucket, please check [minioGetObject directory](examples/minioGetObject)

### Get the presigned URL

To start a connection and get the presigned URL of the file you uploaded, please check [minioGetPresignedURL directory](examples/minioGetPresignedURL)
