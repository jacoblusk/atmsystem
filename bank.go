package atmsystem

type Bank interface {
	Deposit(accout, amount int) error
	Withdrawl(account, amount int) error
	Inquiry(account int) error
}
