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
	"strings"
)

func RefreshToken(db *gorm.DB, data *viewmodels.UserDto) (err error) {
	var (
		model           models.User
		refreshTokenBin utils.UUID
	)

	if refreshTokenBin, err = utils.FromBase64UUID(data.RefreshToken); err != nil {
		return fmt.Errorf("the refresh token [%s] is not valid", data.RefreshToken)
	}

	if err = db.Model(&model).
		Where("id = ?", data.Id).
		Where("refresh_token = ?", refreshTokenBin.OrderedValue()).
		First(&model).Error; err != nil {
		return
	}

	model.RefreshToken = utils.UUID(uuid.NewV1())

	if err = db.Save(model).Error; err != nil {
		return err
	}

	toViewModel(&model, data)
	return
}

func Login(db *gorm.DB, data *viewmodels.LoginDto) (err error) {
	var (
		model          models.User
		hasher         hash.Hash
		hashUserPwd    []byte
		hashUserPwdStr string
		userPwd        []byte
	)

	data.Username = strings.Trim(data.Username, " ")
	data.Password = strings.Trim(data.Password, " ")

	db.Where("username = ?", data.Username).FirstOrInit(&model)

	if model.Username != "" && model.Password != "" {

		hasher = sha1.New()

		userPwd = utils.ToBytes(data.Password)
		hasher.Write(userPwd)

		hashUserPwd = hasher.Sum(nil)
		hashUserPwdStr = fmt.Sprintf("%x", hashUserPwd)

		if hashUserPwdStr == model.Password {
			model.RefreshToken = utils.UUID(uuid.NewV1())
			if err = db.Save(model).Error; err != nil {
				return
			}
			toViewModel(&model, &data.Data)
			return
		}
	}

	return &error2.WrongLoginError{Username: data.Username}
}
