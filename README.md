# ST2DCE-2526PSA01 - DevOps and Continuous Deployment (INGE-3-SEM-A, INGE-3, INGE-3-PRO)

Group 3

### Part 1
## 1
## 2
# Before:
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Name  string
	Title string
	State string
}

func main() {
	request1()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := []whoami{
		whoami{Name: "Efrei Paris",
			Title: "DevOps and Continous Deployment",
			State: "FR",
		},
	}

	json.NewEncoder(response).Encode(who)

	fmt.Println("Endpoint Hit", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "EfreiParis"

	fmt.Fprintf(response, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

# After:
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type whoami struct {
	Title  string
	Names string
	State string
}

func main() {
	startServer()
}

func whoAmI(response http.ResponseWriter, r *http.Request) {
	who := whoami{
		Title: "Group 3",
		Names: "Billy/Bussiere/Godfrin",
		State: "FR",
	}
	
	// Set content type to application/json
	response.Header().Set("Content-Type", "application/json")
	
	// Encode the struct to JSON and write it to the response
	err := json.NewEncoder(response).Encode(who)
	if err != nil {
		http.Error(response, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON encoding error:", err)
		return
	}
	fmt.Println("Endpoint Hit: /whoami", who)
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, r *http.Request) {
	who := "EfreiParis"

	fmt.Fprintf(response, "A little bit about me...")
	fmt.Println("Endpoint Hit: ", who)
}

func startServer() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/whoami", whoAmI)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

## 3

### Part 2
## 1
## 2
## 3

### Part 3
## 1
## 2 
## 3

### Bonus
 ## Part 1: Build the docker image using the buildpack utility and describe what you observe in comparison with the Dockerfile option.
 ## Part 2: Configure another alert and send it by e-mail to abdoul-aziz.zakarimadougou@intervenants.efrei.net.
 
### Tools
- Graphana
- Prometheus
- Jenkins
- Graphan/Loki
- Kubernetes
- Docker
- Kubernetes/helm
