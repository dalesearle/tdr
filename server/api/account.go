package api

import (
	"encoding/json"
	"io"

	"github.com/dalesearle/tdr/server/data"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

type Account struct {
	accountId int
	email     string
	fname     string
	hasAccess bool
	lname     string
	pwd       string
	rights    int64
}

type accountBuilder struct {
	AccountId int
	Email     string
	Fname     string
	HasAccess bool
	Lname     string
	Pwd       string
	Rights    int64
}

func (bldr *accountBuilder) Build() Account {
	return Account{
		accountId: bldr.AccountId,
		email:     bldr.Email,
		fname:     bldr.Fname,
		hasAccess: bldr.HasAccess,
		lname:     bldr.Lname,
		pwd:       bldr.Pwd,
		rights:    bldr.Rights,
	}
}

func (acct *Account) Insert() error {
	tdrDb := data.GetTDRDatabase()
	stmt, err := tdrDb.Db.Prepare("INSERT INTO tdr.account SET email = ?, fname = ?, lname = ? , pwd = ?")
	defer stmt.Close()
	if err == nil {
		_, err = stmt.Exec(acct.Email(), acct.FirstName(), acct.LastName(), acct.Password())
	}
	return err
}

func (acct *Account) EncryptPassword() error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(acct.Password()), 10)
	if err == nil {
		acct.pwd = string(pwd)
		fmt.Println("PWD: " + acct.Password())
	}
	return err
}

func (acct *Account) Id() int {
	return acct.accountId
}

func (acct *Account) Email() string {
	return acct.email
}

func (acct *Account) FirstName() string {
	return acct.fname
}

func (acct *Account) LastName() string {
	return acct.lname
}

func (acct *Account) Password() string {
	return acct.pwd
}

func (acct *Account) Rights() int64 {
	return acct.rights
}

func (acct *Account) HasAccess() bool {
	return acct.hasAccess
}

func NewAccountBuilder() *accountBuilder {
	return new(accountBuilder)
}

/********** Builder Functions **********/

func (bldr *accountBuilder) Decode(body io.ReadCloser) (Account, bool) {
	err := json.NewDecoder(body).Decode(bldr)
	return bldr.Build(), err == nil
}

func (bldr *accountBuilder) FetchAccount(email string) (Account, error) {
	tdrDb := data.GetTDRDatabase()
	err := tdrDb.Db.QueryRow(
		"SELECT accountId, email, fname, lname, pwd, hasAccess, rights FROM tdr.account WHERE email = ?", email).Scan(
		&bldr.AccountId, &bldr.Email, &bldr.Fname, &bldr.Lname, &bldr.Pwd, &bldr.HasAccess, &bldr.Rights)
	return bldr.Build(), err
}
