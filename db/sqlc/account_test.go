package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"dragonfly.io/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	return account

}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acct1 := createRandomAccount(t)
	acct2, err := testQueries.GetAccount(context.Background(), acct1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, acct2)
	require.Equal(t, acct1.Balance, acct2.Balance)
	require.Equal(t, acct1.Owner, acct2.Owner)
	require.Equal(t, acct1.ID, acct2.ID)
	require.WithinDuration(t, acct1.CreatedAt, acct2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	acct1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      acct1.ID,
		Balance: util.RandomMoney(),
	}

	acct2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, acct2)
	require.Equal(t, arg.Balance, acct2.Balance)
	require.Equal(t, acct1.Owner, acct2.Owner)
	require.Equal(t, acct1.ID, acct2.ID)
	require.WithinDuration(t, acct1.CreatedAt, acct2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	acc2, err := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, acc2)

}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}
	arg := ListAccountsParams{
		Offset: 5,
		Limit:  5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}
