package services

import (
	"errors"
	"fmt"
	"learn_gorm/models"
	"learn_gorm/repositories"
	"log"
)

type CreateUserWithAccountParams struct {
	Name    string
	Email   string
	Balance float64
}

type TransferParams struct {
	FromId uint
	ToId   uint
	Amount float64
}

type userAccountService struct {
	baseRepo    repositories.IBaseRepo
	userRepo    repositories.IUserRepo
	accountRepo repositories.IAccountRepo
}

type IUserAccountService interface {
	CreateUserWithAccount(params CreateUserWithAccountParams) error
	Transfer(params TransferParams) error
}

func NewUserAccountService(
	baseRepo repositories.IBaseRepo,
	userRepo repositories.IUserRepo,
	accountRepo repositories.IAccountRepo,
) IUserAccountService {
	return &userAccountService{
		baseRepo:    baseRepo,
		userRepo:    userRepo,
		accountRepo: accountRepo,
	}
}

func (s *userAccountService) CreateUserWithAccount(params CreateUserWithAccountParams) error {
	tx := s.baseRepo.Begin()
	user := models.User{
		Name:  params.Name,
		Email: params.Email,
	}
	if err := s.userRepo.Create(tx, &user); err != nil {
		s.baseRepo.Rollback(tx)
		return err
	}
	info := fmt.Sprintf("New user Id : %d", user.ID)
	fmt.Println(info)
	if err := s.accountRepo.Create(tx, &models.Account{
		UserID:  user.ID,
		Balance: params.Balance,
	}); err != nil {
		s.baseRepo.Rollback(tx)
		return err
	}
	s.baseRepo.Commit(tx)
	return nil
}

func (s *userAccountService) Transfer(params TransferParams) error {
	tx := s.baseRepo.Begin()
	// is user exist
	sender, err := s.userRepo.FindOne(tx, params.FromId)
	if err != nil {
		return errors.New(err.Error())
	}
	infoSender := fmt.Sprintf("Sender name : %s", sender.Name)
	log.Println(infoSender)
	// is receiver exis
	receiver, err := s.userRepo.FindOne(tx, params.ToId)
	if err != nil {
		return errors.New(err.Error())
	}
	infoReceiver := fmt.Sprintf("Receiver name : %s", receiver.Name)
	log.Println(infoReceiver)
	// is sender account exist
	senderAccount, err := s.accountRepo.FindByUserId(tx, uint(params.FromId))
	if err != nil {
		return errors.New(err.Error())
	}
	infoSenderAccount := fmt.Sprintf("Sender account balance : %v", senderAccount.Balance)
	log.Println(infoSenderAccount)
	// cancel process if sender balance smaller than transfer amount
	if senderAccount.Balance < params.Amount {
		return errors.New("insufficience balance")
	}
	// is receiver exist
	receiverAccount, err := s.accountRepo.FindByUserId(tx, uint(params.ToId))
	if err != nil {
		return errors.New(err.Error())
	}
	// substract the sender money with the value of transfer amount
	if err := s.accountRepo.UpdateByUserId(tx, params.FromId, models.Account{
		Balance: senderAccount.Balance - params.Amount,
	}); err != nil {
		return err
	}
	// add the receiver balance with the value of transfer amount
	if err := s.accountRepo.UpdateByUserId(tx, params.ToId, models.Account{
		Balance: receiverAccount.Balance + params.Amount,
	}); err != nil {
		return err
	}
	log.Println("Transfer is successful")
	s.baseRepo.Commit(tx)
	return nil
}
