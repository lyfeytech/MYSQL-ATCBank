package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lyfeytech/MYSQL-ATCBank/pkg/config"
)

var db *gorm.DB

type Account struct{
	gorm.Model
	CustomerID string `gorm:""json:"customerid"`
	SWIFT string `json:"swift"`
	Balance string `json:"balance"`
	Account_Holder string `json:"account_holder"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Account{})
}

func (a *Account) CreateAccount() *Account{
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func GetAllAccounts() []Account{
	var Accounts []Account
	db.Find(&Accounts)
	return Accounts
}

func GetAccountById(Id int64) (*Account, *gorm.DB){
	var getAccount Account
	db:=db.Where("ID=?", Id).Find(&getAccount)
	return &getAccount, db
}

func DeleteAccount(ID int64) Account{
	var account Account
	db.Where("ID=?", ID).Delete(account)
	return account
}