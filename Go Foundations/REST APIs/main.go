package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	fmt.Println(githubInfo("dipti95"))
}

// func demo() {
// 	resp, err := http.Get("https://api.github.com/users/dipti95")

// 	if err != nil {
// 		log.Fatalf("error: %s", err)
// 		/*
// 		   log.Printf("error: %s", err)
// 		   os.Exit(1)
// 		*/
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		log.Fatalf("error: %s", resp.Status)
// 	}

// 	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))
// 	/*
// 		if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
// 			log.Fatalf("error:Can't copy - %s", err)
// 		}
// 	*/

// 	var r Reply

// 	dec := json.NewDecoder(resp.Body)

// 	if err := dec.Decode(&r); err != nil {
// 		log.Fatalf("error: can't decode- %s", err)
// 	}

// 	// fmt.Println(r)
// 	fmt.Printf("%#v\n", r)
// }

// githubInfo returns name and number of public repos for login
func githubInfo(login string) (string, int, error) {
	url := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err

	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf(resp.Status)
	}

	fmt.Printf("Content-Type: %s\n", resp.Header.Get("Content-Type"))

	var r Reply

	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil

}

type Reply struct {
	Name string

	// Public_Repos int

	// without json tag , NumRepos will be zeros because in the body NumRepos is not present
	// if NumRepos have string type, written without json tag then its value is empty string
	NumRepos int `json:"public_repos"`
}

/*JSON <-> Go
true/false <-> true/false
string <-> string
null <-> nil
number <-> float64, float32, int8, int16, int 32, int64,int, uint8 ...
array <-> []any ([]interface{})
object <-> map[string]any, struct

encoding/json API
JSON -> io.Reader -> Go: json.Decoder
JSON -> []byte -> Go: json.Unmarshal
Go -> io.Writer -> JSON: json.Encoder
Go -> []byte -> JSON:json:Marshal
*/
