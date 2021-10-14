## key-value-store-api

This project focusing in-memory key-value storage as a REST API

| Method | Endpoint         | Description                    |
| ------ | ---------------- | ------------------------------ |
| POST   | /store           | Set Key-Value pair to store    |
| GET    | /store?key={key} | Get Key-Value pairs on store   |
| PUT    | /store/flush     | Flush Key-Value pairs on store |

## How to run?

### Run on locally

Download the project and run this magical command:

`go run main.go`

### Run on Docker Container

Download the project and create Docker image with this command:

` docker build -t key-value-store .`

After create Docker image, we need to run application with this command:

`docker run -it -p 9090:9090 --rm --name key-value-store key-value-store`

**Note:** This command runs the application on current terminal session. If you want to run on background use `-d` flag

---

#### Important Note
This project configs for Unix systems, if you want to run on Windows you have to update file path (`persistStorePath`) on [persist.go](api/v1/store/persist.go)

#### Reminder
If you are using [Postman](https://www.postman.com) you can use [Postman Collection](key-value-store.postman_collection.json) that in project folder for quick start