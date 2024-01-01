package tools

import "time"

type mocDB struct {}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "1234567890",
		Username: "alex",
	},
	"jim": {
		AuthToken: "0987654321",
		Username: "jim",
	},
	"jane": {
		AuthToken: "qwertyuiop",
		Username: "jane",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins: 100,
		Username: "alex",
	},
	"jim": {
		Coins: 200,
		Username: "jim",
	},
	"jane": {
		Coins: 300,
		Username: "jane",
	},
}

func (d *mocDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate database query
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mocDB) GetUserCoins(username string) *CoinDetails {
	// simulate database query
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mocDB) SetupDatabase() error {
	return nil
}