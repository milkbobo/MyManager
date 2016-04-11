package account

import (
	. "github.com/fishedee/language"
	. "mymanager/models/common"
	// "crypto/sha1"
	// "fmt"
	// "io"
)

type AccountAoModel struct {
	BaseModel
	AccountDb AccountDbModel
}

func (this *AccountAoModel) Search(userId int, where Account, pageInfo CommonPage) Accounts {

	// wheres := Account{
	// 	UserId:     userId,
	// 	Name:       data.Name,
	// 	Remark:     data.Remark,
	// 	CategoryId: data.CategoryId,
	// 	CardId:     data.CardId,
	// 	Type:       data.Type,
	// }

	where.UserId = userId

	// fmt.Printf("%+v", wheres)

	return this.AccountDb.Search(where, pageInfo)

}

func (this *AccountAoModel) Get(userId, accountId int) Account {
	account := this.AccountDb.Get(accountId)
	if account.UserId != userId {
		Throw(1, "你没有权利查看或编辑等操作")
	}
	return account
}

func (this *AccountAoModel) Add(userId int, account Account) {

	// accounts := Account{
	// 	UserId:     userId,
	// 	Name:       account.Name,
	// 	Money:      account.Money,
	// 	Remark:     account.Remark,
	// 	CategoryId: account.CategoryId,
	// 	CardId:     account.CardId,
	// 	Type:       account.Type,
	// }

	account.UserId = userId

	this.AccountDb.Add(account)
}

func (this *AccountAoModel) Mod(userId int, account Account) {

	//检查该类型是不是属于他本人
	this.Get(userId, account.AccountId)

	this.AccountDb.Mod(account)
}
func (this *AccountAoModel) Del(userId, accountId int) {

	//检查该类型是不是属于他本人
	this.Get(userId, accountId)

	this.AccountDb.Del(accountId)
}
