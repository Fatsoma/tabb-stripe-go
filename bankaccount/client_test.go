package bankaccount

import (
	"testing"

	stripe "github.com/fatsoma/tabb-stripe-go"
	_ "github.com/fatsoma/tabb-stripe-go/testing"
	assert "github.com/stretchr/testify/require"
)

func TestBankAccountDel_ByAccount(t *testing.T) {
	bankAcount, err := Del("ba_123", &stripe.BankAccountParams{
		Account: stripe.String("acct_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountDel_ByCustomer(t *testing.T) {
	bankAcount, err := Del("ba_123", &stripe.BankAccountParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountGet_ByAccount(t *testing.T) {
	bankAcount, err := Get("ba_123", &stripe.BankAccountParams{Account: stripe.String("acct_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountGet_ByCustomer(t *testing.T) {
	bankAcount, err := Get("ba_123", &stripe.BankAccountParams{Customer: stripe.String("cus_123")})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountList_ByAccount(t *testing.T) {
	i := List(&stripe.BankAccountListParams{Customer: stripe.String("acct_123")})

	// Verify that we can get at least one bank account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BankAccount())
}

func TestBankAccountList_ByCustomer(t *testing.T) {
	i := List(&stripe.BankAccountListParams{Customer: stripe.String("cus_123")})

	// Verify that we can get at least one bank account
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.BankAccount())
}

func TestBankAccountNew_ByAccount(t *testing.T) {
	bankAcount, err := New(&stripe.BankAccountParams{
		Account:            stripe.String("acct_123"),
		DefaultForCurrency: stripe.Bool(true),
		Token:              stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountNew_ByCustomer(t *testing.T) {
	bankAcount, err := New(&stripe.BankAccountParams{
		Customer: stripe.String("cus_123"),
		Token:    stripe.String("tok_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountUpdate_ByAccount(t *testing.T) {
	bankAcount, err := Update("ba_123", &stripe.BankAccountParams{
		Account:            stripe.String("acct_123"),
		DefaultForCurrency: stripe.Bool(true),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}

func TestBankAccountUpdate_ByCustomer(t *testing.T) {
	bankAcount, err := Update("ba_123", &stripe.BankAccountParams{
		AccountHolderName: stripe.String("New Name"),
		Customer:          stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, bankAcount)
}
