package handlers

import (
	"bankSystem/internal/domain"
	"bankSystem/internal/dto"
	"bankSystem/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccountHandler struct {
	accountService *service.AccountService
}

func NewAccountController(serviceAccount *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: serviceAccount}
}

// @Summary Create account
// @Description Create new account for user
// @Tags Account
// @Accept json
// @Produce      json
// @Param        account  body      dto.CreateAccountRequest  true  "Account to create"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/create [post]
func (ac *AccountHandler) CreateAccount(c *gin.Context) {
	var req dto.CreateAccountRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	user := domain.User{Login: req.Login}
	err := ac.accountService.NewUserAccount(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "account created"})
}

// @Summary Get balance
// @Description Get balance from account
// @Tags Account
// @Accept json
// @Produce      json
// @Param        id  query string  true  "Id acccount"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/balance [get]
func (ac *AccountHandler) GetBalance(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
		return
	}

	balance, err := ac.accountService.GetBalance(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"balance": balance})
	return
}

// @Summary Deposit
// @Description Deposit account
// @Tags Account
// @Accept json
// @Produce      json
// @Param        account  body      dto.DepositWithdrawRequest  true  "Deposited"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/deposit [post]
func (ac *AccountHandler) Deposit(c *gin.Context) {
	var req dto.DepositWithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := ac.accountService.Deposit(req.Id, req.Amount); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account replenished"})
	return
}

// @Summary Withdraw
// @Description Withdraw account
// @Tags Account
// @Accept json
// @Produce      json
// @Param        account  body      dto.DepositWithdrawRequest  true  "Deposited"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/withdraw [post]
func (ac *AccountHandler) Withdraw(c *gin.Context) {
	var req dto.DepositWithdrawRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := ac.accountService.Withdraw(req.Id, req.Amount); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Money has been withdrawn"})
	return
}

// @Summary Transfer
// @Description Transfer account
// @Tags Account
// @Accept json
// @Produce      json
// @Param        account  body      dto.TransferRequest  true  "Transfered"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/transfer [post]
func (ac *AccountHandler) Transfer(c *gin.Context) {
	var req dto.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := ac.accountService.Transfer(req.Id1, req.Id2, req.Amount); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Money has been withdrawn"})
	return
}

// @Summary Delete
// @Description Delete account
// @Tags Account
// @Accept json
// @Produce      json
// @Param        account  body      dto.DeleteAccountRequest  true  "Deleted"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /account/delete [delete]
func (ac *AccountHandler) DeleteAccount(c *gin.Context) {
	var req dto.DeleteAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := ac.accountService.DeleteAccount(req.Id, req.Login); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "Account was deleted"})
	return
}

// @Summary Get Transactions
// @Description Get all transactions for account
// @Tags Account
// @Accept json
// @Produce json
// @Param id query string true "Account ID"
// @Success 200 {array} domain.Transaction
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /account/transactions [get]
func (ac *AccountHandler) GetTransactions(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id query parameter is required"})
		return
	}

	transactions, err := ac.accountService.GetTransactions(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
