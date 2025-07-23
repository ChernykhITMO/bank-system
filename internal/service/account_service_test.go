package service

import (
	"bankSystem/internal/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
)

type fakeTxManager struct {
}

func (f *fakeTxManager) WithTx(fn func(tx *gorm.DB) error) error {
	return fn(&gorm.DB{})
}

type mockAccountRepo struct {
	getCalled  bool
	saveCalled bool
	account    *model.AccountEntity
}

func (m *mockAccountRepo) GetAccount(tx *gorm.DB, id string) (*model.AccountEntity, error) {
	m.getCalled = true
	return m.account, nil
}

func (m *mockAccountRepo) SaveAccount(tx *gorm.DB, acc *model.AccountEntity) error {
	m.saveCalled = true
	m.account = acc
	return nil
}

func (m *mockAccountRepo) DeleteAccount(*gorm.DB, string) error { return nil }

type mockTransactionRepo struct {
	saveCalled bool
}

func (m *mockTransactionRepo) SaveTransaction(tx *gorm.DB, t *model.TransactionEntity) error {
	m.saveCalled = true
	return nil
}

func (m *mockTransactionRepo) GetTransactionsByAccountId(*gorm.DB, string) ([]model.TransactionEntity, error) {
	return nil, nil
}

func TestWithdraw_Success(t *testing.T) {
	initialBalance := 100.0
	withdrawAmount := 40.0
	expectedBalance := initialBalance - withdrawAmount

	account := &model.AccountEntity{
		Id:      uuid.NewString(),
		Login:   "user1",
		Balance: initialBalance,
	}

	tx := &fakeTxManager{}
	accRepo := &mockAccountRepo{account: account}
	txRepo := &mockTransactionRepo{}

	accService := NewAccountService(tx, accRepo, nil, nil, txRepo)

	err := accService.Withdraw(account.Id, withdrawAmount)

	assert.NoError(t, err)
	assert.True(t, accRepo.getCalled, "GetAccount должен быть вызван")
	assert.True(t, accRepo.saveCalled, "SaveAccount должен быть вызван")
	assert.True(t, txRepo.saveCalled, "SaveTransaction должен быть вызван")
	assert.Equal(t, expectedBalance, accRepo.account.Balance, "Баланс должен уменьшиться")
}
