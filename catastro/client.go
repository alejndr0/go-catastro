package catastro

import (
	"encoding/json"
	"net/http"
)

func getJson(url string, queryparams map[string]string, target Response) error {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "http://ovc.catastro.meh.es/"+url, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	for k, v := range queryparams {
		if v == "" {
			continue
		}
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	// Set the User-Agent header because the server returns a 403 without it.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36")

	r, err := client.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}

	failed, err := target.hasFailed()
	if failed {
		return err
	}

	return nil
}

type Response interface {
	hasFailed() (bool, error)
}
