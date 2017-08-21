package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
    "strings"
    "strconv"
    "io/ioutil"
    "net/http"
)

type DistanceResponse struct {
    Zipcode1 string
    Zipcode2 string
    Distance string
    Error string
}

type Message struct {
    Text string
}

type ZipCodeServiceResponse struct {
    Distance float64 `json:"distance"`
}

/**
 * Distance Handler
 */
func distanceHandler(w http.ResponseWriter, r *http.Request) {
    pathParams := strings.Split(r.URL.Path[1:], "/")
    var responseObject DistanceResponse
    if len(pathParams) < 3 {
        responseObject = buildError(pathParams)
    } else {
        responseObject = calculateDistance(pathParams[1], pathParams[2])
    }
    resp, err := json.Marshal(responseObject)
    if err != nil {
        panic(err)
    }
    w.Write(resp)
}

func about (w http.ResponseWriter, r *http.Request) {
    m := Message{"Zipcode Microservice. \nZipcodes pasesed in as URL parameters.  Response sent as JSON"}
    b, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }

    w.Write(b)
}

func buildError(pathParams []string) DistanceResponse {
    var errMsg string
    var zipcode1 string
    var zipcode2 string

    if len(pathParams) < 2 {
        errMsg = "A second zipcode is required to calculate distance"
    } else {
        if len(pathParams) < 3 {
            zipcode1 = pathParams[1]
            errMsg = "Two zipcodes must be specified"
        }
    }
    return DistanceResponse{zipcode1, zipcode2, "", errMsg}
}

func calculateDistance(zipcode1 string, zipcode2 string) DistanceResponse {
    log.Printf("Calculating distance between: %s & %s\n", zipcode1, zipcode2)
    apiKey := "JTYXJcDQnOzh36tqhOwALppamFMXctMWmKT6WwXklwBoC5PUTlvFKUP7AC5N5RJs"

    url := fmt.Sprintf("http://www.zipcodeapi.com/rest/%s/distance.json/%s/%s/mile", apiKey, zipcode1, zipcode2)

    client := &http.Client{
        Timeout: time.Second * 10,
    }

    resp, err := client.Get(url)
    if err != nil {
        log.Fatal("Get: ", err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        log.Fatal("Parse: ", err)
    }

    var record ZipCodeServiceResponse
    err = json.Unmarshal(body, &record)
    if err != nil {
        log.Fatal("Unmarshal JSON: ", err)
    }

    return DistanceResponse{zipcode1, zipcode2, strconv.FormatFloat(record.Distance, 'f', 0, 32) + " miles", ""}
}

/**
 * Main Function
 */
func main() {
    http.HandleFunc("/distance/", distanceHandler)
    http.HandleFunc("/about/", about)
    http.ListenAndServe(":8080", nil)
}

