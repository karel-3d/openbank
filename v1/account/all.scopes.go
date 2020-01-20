// Code generated by "scopegen"; DO NOT EDIT.
package accounts

var Scopes = map[string]string{
	"https://auth.bnk.to/account.read": "View account data",
	"https://auth.bnk.to/account.write": "Manage account data",
}

var AuthScopes = map[string][]string{
	"/accounts.AccountService/CheckAccount": []string{"https://auth.bnk.to/account.read"},
	"/accounts.AccountService/CreateAccount": []string{"https://auth.bnk.to/account.write"},
	"/accounts.AccountService/DeleteAccount": []string{"https://auth.bnk.to/account.write"},
	"/accounts.AccountService/GetAccount": []string{"https://auth.bnk.to/account.read"},
	"/accounts.AccountService/GetAccountStatus": []string{"https://auth.bnk.to/account.read"},
	"/accounts.AccountService/GetAccounts": []string{"https://auth.bnk.to/account.read"},
	"/accounts.AccountService/UpdateAccount": []string{"https://auth.bnk.to/account.write"},
	"/accounts.AccountService/UpdateAccountStatus": []string{"https://auth.bnk.to/account.write"},
}