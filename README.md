#traxero-go-microservice
##Traxero Go Microservice Verify VIN Application
---
##Purpose:
- This is a microservice that verifies VINs (Vehicle Identification Numbers)

##Installation
1. Clone the repository
2. If you are on a Mac, at your command line, run:
```./traxero-go-microservice``` from the root of the repository.
3. If you are on a Windows machine, at your command line, run:
```go run main.go``` from the root of the repository.
4. What you will see is a message that says "Listening on port 8000" in the terminal.
5. Open [Postman app](https://www.getpostman.com/) and navigate to the following URL:
```http://localhost:8000/query?vin=<VIN>``` where ```<VIN>``` is the VIN you want to verify.

##Usage
1. Go to [Postman app](https://www.getpostman.com/) and navigate to the following URL:
```http://localhost:8000/query?vin=<VIN>``` where ```<VIN>``` is the VIN you want to verify.
2. Hit the "Send" button and make sure to set Authorization with the following token value:
```Bearer Token: ShYj6GJ37rXBFRiTew3GkdKVMpcp52```
3. You will see the response from the microservice.

##Endpoints
###/query?vin=<VIN>
- This endpoint is used to verify a VIN.
- The VIN is passed as a query parameter.
- The response is a JSON object with the following fields:
  - `make`: Make of the vehicle.
  - `model`: Model of the vehicle.
  - `year`: Year of the vehicle.
