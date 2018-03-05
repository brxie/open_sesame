package openalpr

import (
	"os"
	"mime/multipart"
	"bytes"
    "net/http"
    "io"
    "io/ioutil"
    "errors"
    "encoding/json"
	"fmt"
	"github.com/open_sesame/utils"
)

const apiURL = "https://api.openalpr.com/v2/"

type Alpr struct {
    RecognizeVehicle    int
    Country             string
    SecretKey           string
    TopN                int
}

type alprResponse struct {
    Results     []Result
}

type Result struct {
    Plate       string
    Confidence  float32
    Candidates  []Candidates
}

type Candidates struct {
    Plate       string
    Confidence  float32
}

// Recognize recognizes licence plates on provided image and
// returns slice witg result of recognized plates.
func (a Alpr) Recognize(file string) ([]Result, error) {
    resp, err := a.upload(a.requestURL(), file)
    if err != nil {
        return nil, err
    }

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        log.Log.Error("Rocognizing image failed!\n" +
            "Status: " + resp.Status + "\n" +
            "Document body: " + string(body) + "\n" +
            "Request url: " + a.requestURL())
        
        return nil, errors.New("rocognizing image failed")
    }

    result, err := a.parseResponse(&resp.Body)
    if err != nil {
        return nil, err
    }
    
    return result, nil
}

func (a Alpr) parseResponse(respBody *io.ReadCloser) ([]Result, error) {
    var res = alprResponse{}
    err := json.NewDecoder(*respBody).Decode(&res)

    if err != nil {
        log.Log.Error("Failed to decode response!", err)
        return nil, err
    }

    return res.Results, nil
}

func (a Alpr) requestURL() (url string) {
    url = apiURL + fmt.Sprintf("recognize?recognize_vehicle=%d&country=%s&" +
        "secret_key=%s&topn=%d", a.RecognizeVehicle,
                                 a.Country,
                                 a.SecretKey,
                                 a.TopN)
    return
}

func (a Alpr) upload(url, file string) (*http.Response, error) {
    var buff bytes.Buffer
    log.Log.Debugf("Uploading image %s. Request URL: %s", file, url)
    writer := multipart.NewWriter(&buff)
	img, err := os.Open(file)
	
    if err != nil {
        return nil, err
    }
    defer img.Close()
	
	fw, err := writer.CreateFormFile("image", file)
    if err != nil {
        return nil, err
    }
    if _, err = io.Copy(fw, img); err != nil {
        return nil, err
    }
    // Add the other fields
    if fw, err = writer.CreateFormField("key"); err != nil {
        return nil, err
    }
    if _, err = fw.Write([]byte("KEY")); err != nil {
        return nil, err
	}
	
    writer.Close()

    req, err := http.NewRequest("POST", url, &buff)
    if err != nil {
        return nil, err
    }
    req.Header.Set("Content-Type", writer.FormDataContentType())

    client := &http.Client{}
	res, err := client.Do(req)

    if err != nil {
        return nil, err
    }
	
    return res, nil
}