# zipcode-microservice-golang

**An implementation of my zipcode microservice using Golang**

*The original service is in this repository: https://github.com/echovue/zipcode-microservice*

A Go based microservice which determines the distance for a specific zip code.

To build the project, download it into a local directory.

You will need to update the apiKey property in ZipcodeDistanceService with your own from https://www.zipcodeapi.com/API

To execute this program, you will need to have Go installed and configured on your workstation.

To build the executable, execute the following from the commandline in the microservice directory:

`go build`

`./microservice`

An example request which you can submit is:

http://localhost:8080/distance/97035/97001

And the response should be:

```javascript
{
    "Zipcode1": "97035",
    "Zipcode2": "97001",
    "Distance": "107 miles"
}
```

An incorrect request, such as one which includes an invalid zip code should result in:

```javascript
{
    "Zipcode1": "97035",
    "Zipcode2": "97001",
    "Error": "Unable to calculate distance"
}
```


