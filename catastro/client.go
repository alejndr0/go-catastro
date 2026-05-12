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

	// Catastro's WAF (F5 ASM) 400s requests with no User-Agent, a "curl/*" or
	// "Go-http-client/*" UA, or the previously hardcoded
	// "...Mac OS X 10_15_7...Chrome/113.0.0.0..." string (a well-known scraper
	// fingerprint that has since been blocklisted). Use a current browser UA.
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/131.0.0.0 Safari/537.36")

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
