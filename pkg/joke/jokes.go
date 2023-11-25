package joke

import (
	"encoding/json"
	"net/http"
	"sync"
)

type Joke struct {
	IconURL string `json:"icon_url"`
	ID      string `json:"id"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

type JokeClient interface {
	GetJoke() (Joke, error)
	GetJokes(limit int) ([]Joke, error)
}

type ChuckNorrisJokeClient struct {
	baseURL    string
	httpClient *http.Client
	limit      int
}

func NewChuckNorrisJokeClient(baseURL string, httpClient *http.Client, limit int) *ChuckNorrisJokeClient {
	return &ChuckNorrisJokeClient{baseURL: baseURL, httpClient: httpClient, limit: limit}
}

func (c *ChuckNorrisJokeClient) GetJoke() (Joke, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+"/jokes/random", nil)
	if err != nil {
		return Joke{}, err
	}
	req.Header.Set("Accept", "application/json")
	res, err := c.httpClient.Do(req)
	if err != nil {
		return Joke{}, err
	}
	defer res.Body.Close()
	var joke Joke
	err = json.NewDecoder(res.Body).Decode(&joke)
	if err != nil {
		return Joke{}, err
	}
	return joke, nil
}

func (c *ChuckNorrisJokeClient) GetJokes(limit int) ([]Joke, error) {
	var mutex sync.Mutex
	var jokes []Joke

	wg := sync.WaitGroup{}
	wg.Add(c.limit)

	for i := 0; i < c.limit; i++ {
		go func() {
			joke, err := c.GetJoke()
			if err != nil {
				wg.Done()
				return
			}
			mutex.Lock()
			jokes = append(jokes, joke)
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	return jokes, nil
}
