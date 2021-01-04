package auth

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

func Login(db *gorm.DB, data *viewmodels.LoginDto) error {
	var (
		model          models.User
		hasher         hash.Hash
		hashUserPwd    []byte
		hashUserPwdStr string
		userPwd        []byte
	)

	db.Where("username = ?", data.Username).FirstOrInit(&model)

	if model.Username != "" && model.Password != "" {
		hasher = sha1.New()

		userPwd = utils.ToBytes(data.Password)
		hasher.Write(userPwd)

		hashUserPwd = hasher.Sum(nil)
		hashUserPwdStr = fmt.Sprintf("%x", hashUserPwd)

		if hashUserPwdStr == model.Password {
			toViewModel(&model, &data.Data)
			return nil
		}
	}

	return &error2.WrongLoginError{Username: data.Username}
}
