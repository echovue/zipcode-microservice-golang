# zipcode-microservice-golang

**An implementation of my zipcode microservice using Golang**

*The original service is in this repository: [https://github.com/echovue/zipcode-microservice]https://github.com/echovue/zipcode-microservice*

A Go based microservice which determines the distance for a specific zip code.

To build the project, download it into a local directory.

You will need to update the apiKey property in ZipcodeDistanceService with your own from https://www.zipcodeapi.com/API

To execute this program, you will need to have Go installed and configured on your workstation.

To build the executable, execute the following from the commandline:

go microservice

./microservice

An example request which you can submit is:

http://localhost:8080/distance/97035/97001

And the response should be:

```javascript
{
    "zipcode1": "97035",
    "zipcode2": "97001",
    "distance": "107 miles"
}
```

An incorrect request, such as one which includes an invalid zip code should result in:

```javascript
{
    "zipcode1": "97035",
    "zipcode2": "97001",
    "error": "Unable to calculate distance"
}
```


