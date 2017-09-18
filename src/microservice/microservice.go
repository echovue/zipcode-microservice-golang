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
 * Distance Calculation Endpoint Handler.
 */
func distanceHandler(w http.ResponseWriter, r *http.Request) {
    var responseObject DistanceResponse
    pathParams := strings.Split(r.URL.Path[1:], "/")

    if len(pathParams) < 3 {
        responseObject = buildError(pathParams)
    } else {
        responseObject = calculateDistance(pathParams[1], pathParams[2])
    }

    resp, err := json.Marshal(responseObject)
    if err != nil {
        panic(err)
    }
    w.WriteHeader(200)
    w.Write(resp)
}

/**
 * About Endpoint Handler.
 */
func aboutHandler (w http.ResponseWriter, r *http.Request) {
    m := Message{"Zipcode Microservice. Zipcodes passed in as URL parameters.  Response sent as JSON"}
    resp, err := json.Marshal(m)

    if err != nil {
        panic(err)
    }
    w.WriteHeader(200)
    w.Write(resp)
}

/**
 * Health Endpoint Handler.
 */
func healthHandler(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Server", "Go Web Server Hosting ZipCode Service")
  w.WriteHeader(200)
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
    apiKey := "XXXXXXXXXXXXXXXX" //Replace with your own API Key from ZipCodeAPI.com

    url := fmt.Sprintf("http://www.zipcodeapi.com/rest/%s/distance.json/%s/%s/mile", apiKey, zipcode1, zipcode2)

    client := &http.Client{
        Timeout: time.Second * 10,
    }

    resp, err := client.Get(url)
    if err != nil {
        log.Fatal("Get: ", err)
    }

    if resp.StatusCode != 200 {
        log.Printf("zipcodeapi.com returned Status Code: %d", resp.StatusCode)
        return DistanceResponse{zipcode1, zipcode2, "", "Error retrieving distance for ZipCodeAPI. Check your API Key."}
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
    http.HandleFunc("/about/", aboutHandler)
    http.HandleFunc("/health/", healthHandler)
    http.ListenAndServe(":8080", nil)
}

