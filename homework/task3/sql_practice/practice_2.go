package main

import (
	"errors"
	"fmt"
	"github.com/go-master/db"
	"gorm.io/gorm"
)

type Account struct {
	ID      int     `gorm:"primary_key"`
	Name    string  `gorm:"unique"`
	Balance float64 `gorm:"type:decimal(10,2)"`
}

type Transaction struct {
	ID            int `gorm:"primary_key"`
	FromAccountId int
	ToAccountId   int
	Amount        float64 `gorm:"type:decimal(10,2)"`
}

func main() {
	//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	//要求 ：
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，
	//并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	dsn := "root:jsm1234@tcp(192.168.159.132:3306)/homework?charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai"
	if err := db.InitDB(dsn); err != nil {
		panic("数据初始化错误：" + err.Error())
	}

	if db.DB == nil {
		panic("数据库未初始化")
	}

	initTableErr := db.DB.AutoMigrate(&Account{}, &Transaction{})
	if initTableErr != nil {
		panic("failed to init table")
	}

	//1.创建2个账户：A账户，B账户
	accountA := Account{Name: "A", Balance: 1000}
	accountB := Account{Name: "B", Balance: 2000}
	createAErr := db.DB.Debug().Create(&accountA).Error
	if createAErr != nil {
		panic("创建账户A失败")
	}
	createBErr := db.DB.Debug().Create(&accountB).Error
	if createBErr != nil {
		panic("创建账户B失败")
	}

	//账户 A 向账户 B 转账 100 元的操作
	transferAmount := float64(1000)

	err := transfer(&transferAmount)
	if err != nil {
		fmt.Println("转账失败了，失败原因：", err.Error())
	}

}

func transfer(transferAmount *float64) error {
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		//2.查询A账户和B账户的余额
		searchAccountA := &Account{}
		findAErr := db.DB.Debug().Where("name = ?", "A").First(searchAccountA).Error
		if findAErr != nil {
			return errors.New("查询A账户信息失败")
		}
		searchAccountB := &Account{}
		findBErr := db.DB.Debug().Where("name = ?", "B").First(searchAccountB).Error
		if findBErr != nil {
			return errors.New("查询B账户信息失败")
		}

		//3.检查账户A的余额，查看余额是否够转账
		if searchAccountA.Balance < *transferAmount {
			return errors.New("余额不足")
		}
		//4.更新A账户和B账户的余额
		searchAccountA.Balance = searchAccountA.Balance - *transferAmount
		updateAError := db.DB.Debug().Updates(searchAccountA).Error
		if updateAError != nil {
			return errors.New("更新账户A余额失败")
		}
		searchAccountB.Balance = searchAccountB.Balance + *transferAmount
		updateBError := db.DB.Debug().Updates(searchAccountB).Error
		if updateBError != nil {
			return errors.New("更新账户B余额失败")
		}
		//5.在transactions表增加转账记录
		transaction := &Transaction{FromAccountId: searchAccountA.ID, ToAccountId: searchAccountB.ID, Amount: *transferAmount}
		createTransactionError := db.DB.Debug().Create(&transaction).Error
		if createTransactionError != nil {
			return errors.New("创建转账记录失败")
		}
		return nil
	})
	return err
}
