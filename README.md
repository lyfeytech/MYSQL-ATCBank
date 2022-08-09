# #4 MYSQL ATC Bank

Wire Transfer is a Golang project which utilizes real database, where customers can:

1. Wire Transactions can be issued via a Browser or terminal command-line interface

2. User will get a status whether the funds were transferred or insufficient funds to transfer

3. Business logic in the transfer server.

4. Database is needed. 

5. Record successful and unsuccessful transactions to a file

6. You can add account holders [Name, Acct #, Opening Balance]

## Tools

[golang](https://go.dev/dl/go1.19.darwin-amd64.pkg) to install golang and check for go version.

[gorilla mux](https://github.com/gorilla/mux) to install gorilla mux.

[postman](https://www.postman.com/downloads/) to install postman.

```bash
go version
```

## Packages 

```golang


# Routes

import (
	"github.com/gorilla/mux"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/controllers"
)

var RegisterBankRoutes = func (router *mux.Router){
	router.HandleFunc("/account/", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/account/", controllers.GetAccount).Methods("GET")
	router.HandleFunc("/account/{customerId}", controllers.GetAccountById).Methods("GET")
	router.HandleFunc("/account/{customerId}", controllers.UpdateAccount).Methods("PUT")
	router.HandleFunc("/account/{customerId}", controllers.DeleteAccount).Methods("DELETE")

}

# Controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/utils"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/models"
)

# Models

import (
	"github.com/jinzhu/gorm"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/config"
)

var db *gorm.DB

# Config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db * gorm.DB
)

# Utils

import(
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParseBody(r *http.Request, x interface{}){
	if body, err := ioutil.ReadAll(r.Body); err == nil{
		if err := json.Unmarshal([]byte(body), x); err != nil{
			return
		}
	}
}

# Main

import(
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/routes"
)

func main(){
	r := mux.NewRouter()
	routes.RegisterBankRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
```

## Contribution
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[LYFEYTECH](https://github.com/lyfeytech)
