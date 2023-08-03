package objectTest

type CurrencyAccount struct {
	UserName      string
	AgencyNumber  int
	AccountNumber int
	Balance       float32
}

// Its a extension function. I like it
func (c *CurrencyAccount) Withdraw(number float32) float32 {
	c.Balance = c.Balance - number
	return c.Balance
}

func (u *CurrencyAccount) VerifyUserName(userName string) bool {
	return u.UserName == userName
}
