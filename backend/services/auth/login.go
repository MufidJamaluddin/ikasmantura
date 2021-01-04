package auth

import (
	error2 "backend/error"
	"backend/models"
	"backend/utils"
	"backend/viewmodels"
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"hash"
)

func RefreshToken(db *gorm.DB, data *viewmodels.UserDto) (err error) {
	var (
		model models.User
	)

	if err = db.Model(&model).
		Where("id = ?", data.Id).
		First(&model).Error;
	err != nil {
		return
	}

	model.RefreshToken = uuid.NewV4()

	if err = db.Save(model).Error; err != nil {
		return err
	}

	toViewModel(&model, data)
	return
}

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
