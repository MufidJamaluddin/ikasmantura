package impl

import (
	error2 "backend/error"
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"crypto/sha1"
	"fmt"
	"gorm.io/gorm"
	"hash"
)

type AuthServiceImpl struct {
	DB        *gorm.DB
	SecretKey string
}

func (p *AuthServiceImpl) toData(in *models.User, out *viewmodels.UserDto) {
	out.Id = in.ID
	out.Name = in.Name
	out.Username = in.Username
	out.Password = in.Password
	out.Address.ID = in.Address.ID
	out.Address.Street = in.Address.Street
	out.Address.Suite = in.Address.Suite
	out.Address.City = in.Address.City
	out.Address.Zipcode = in.Address.Zipcode

	utils.FillCreated(in, out)
	utils.FillUpdated(in, out)
}

func (p *AuthServiceImpl) Login(data *viewmodels.LoginDto) error {
	var (
		model          models.User
		hasher         hash.Hash
		hashUserPwd    []byte
		hashUserPwdStr string
		userPwd        []byte
	)

	p.DB.Where("username = ?", data.Username).FirstOrInit(&model)

	if model.Username != "" {
		hasher = sha1.New()

		userPwd = append(userPwd, data.Password...)
		hasher.Write(userPwd)

		hashUserPwd = hasher.Sum(nil)
		hashUserPwdStr = fmt.Sprintf("%x", hashUserPwd)

		if hashUserPwdStr == model.Password {
			p.toData(&model, &data.Data)
			return nil
		}
	}

	return &error2.WrongLoginError{Username: data.Username}
}
