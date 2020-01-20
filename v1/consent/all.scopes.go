// Code generated by "scopegen"; DO NOT EDIT.
package consent

var Scopes = map[string]string{
	"https://auth.bnk.to/consent.read": "View consent data",
	"https://auth.bnk.to/consent.write": "Manage consent data",
}

var AuthScopes = map[string][]string{
	"/consent.ConsentService/AnswerConsentChallenge": []string{"https://auth.bnk.to/consent.write"},
	"/consent.ConsentService/CreateConsentEmail": []string{"https://auth.bnk.to/consent.write"},
	"/consent.ConsentService/CreateConsentSMS": []string{"https://auth.bnk.to/consent.write"},
	"/consent.ConsentService/GetConsents": []string{"https://auth.bnk.to/consent.read"},
	"/consent.ConsentService/RevokeConsent": []string{"https://auth.bnk.to/consent.write"},
}