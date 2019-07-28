package youtube

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/GoYoutubeServer/server/config"
	"github.com/GoYoutubeServer/server/models"
)

const baseURL = "https://www.googleapis.com"

type APIClient struct {
	key        string
	httpClient *http.Client
}

func New(key string) *APIClient {
	apiClient := &APIClient{key, &http.Client{}}
	return apiClient
}

func (api *APIClient) doRequest(method, urlPath string, query map[string]string) (body []byte, err error) {
	baseURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	apiURL, err := url.Parse(urlPath)
	if err != nil {
		fmt.Println(err)
		return
	}
	endpoint := baseURL.ResolveReference(apiURL).String()
	log.Printf("action=doRequest endpoint=%s", endpoint)
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	q := req.URL.Query()
	for key, value := range query {
		//fmt.Println("★  1",key,value)
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	//fmt.Println("★  2", req)
	resp, err := api.httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
func (api *APIClient) GetYoutube(r *http.Request) (*models.YoutubeJson, error) {
	//nextPageToken := GetNextPageToken()
	youtubeUrl := "youtube/v3/search"
	// youtubeUrl := "youtube/v3/videoCategories"

	//q, _ := url.ParseQuery(r.URL.RawQuery)

	v, _ := url.Parse(r.URL.RawQuery)
	fmt.Println(111, v)

	qmap := make(map[string]string)
	//for k, v := range q {
	//	qmap[k] = v
	//}
	maxResults, err := strconv.Atoi(r.URL.Query().Get("maxResults"))
	if err != nil {
		fmt.Println(err)
		maxResults = 10
	}
	if maxResults > 20 {
		maxResults = 20
	}

	qmap["key"] = config.Config.APIKey
	qmap["part"] = "snippet"
	// qmap["regionCode"] = "JP"
	// qmap["hl"] = "ja-JP"
	qmap["maxResults"] = strconv.Itoa(maxResults)
	//qmap["type"]="playlist"
	qmap["q"] = r.URL.Query().Get("q")
	qmap["pageToken"] = r.URL.Query().Get("pageToken")
	qmap["type"] = "video"
	// qmap["relevanceLanguage="] = "ja"
	// qmap["location"] = "35.652832,139.839478"
	// qmap["locationRadius"] = "1000km"
	// qmap["safeSearch"] = r.URL.Query().Get("safeSearch")
	// qmap["publishedAfter"] = r.URL.Query().Get("publishedAfter")
	qmap["order"] = r.URL.Query().Get("order")

	fmt.Println("★★  ", qmap)

	//q.URL.RawQuery
	//m, _ := url.ParseQuery(r.URL.RawQuery)

	//resp, err := api.doRequest("GET", youtubeUrl, map[string]string{"key": config.Config.APIKey, "part": "snippet", "regionCode": "JP", "q": q})
	resp, err := api.doRequest("GET", youtubeUrl, qmap)
	log.Printf("url=%s resp=%s", youtubeUrl, string(resp))
	//log.Printf("youtube.go url=%s", youtubeUrl)
	if err != nil {
		log.Printf("action=GetYoutube err=%s", err.Error())
		return nil, err
	}
	var youtube models.YoutubeJson
	//(JSON → Go Object)
	err = json.Unmarshal(resp, &youtube)
	if err != nil {
		log.Printf("action GetYoutube err=%s", err.Error())
		return nil, err
	}
	return &youtube, nil

}
