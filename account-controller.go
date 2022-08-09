package controllers

import(
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/utils"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/models"
)

var NewAccount models.Account

func GetAccount(w http.ResponseWriter, r *http.Request){
	newAccounts:=models.GetAllAccounts()
	res, _ := json.Marshal(newAccounts)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAccountById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	ID, err := strconv.ParseInt(accountId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	accountDetails, _ := models.GetAccountById(ID)
	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateAccount(w http.ResponseWriter, r *http.Request){
	CreateAccount := &models.Account{}
	utils.ParseBody(r, CreateAccount)
	a := CreateAccount.CreateAccount()
	res, _ := json.Marshal(a)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	accountId := vars["accountsId"]
	ID, err := strconv.ParseInt(accountId, 0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	account := models.DeleteAccount(ID)
	res, _ := json.Marshal(account)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateAccount(w http.ResponseWriter, r *http.Request){
	var updateAccount = &models.Account{}
	utils.ParseBody(r, updateAccount)
	vars := mux.Vars(r)
	accountId := vars["accountId"]
	ID, err := strconv.ParseInt(accountId, 0,0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	accountDetails, db :=models.GetAccountById(ID)
	if updateAccount.CustomerID != ""{
		accountDetails.CustomerID = updateAccount.CustomerID
	}
	if updateAccount.SWIFT != ""{
		accountDetails.SWIFT = updateAccount.SWIFT
	}
	if updateAccount.Balance !=""{
		accountDetails.Balance = updateAccount.Balance
	}
	if updateAccount.Account_Holder !=""{
		accountDetails.Account_Holder = updateAccount.Account_Holder
	}
	db.Save(&accountDetails)
	res, _ := json.Marshal(accountDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}