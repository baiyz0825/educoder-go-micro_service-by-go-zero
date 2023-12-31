// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyz0825/school-share-buy-backend/apps/order/cmd/rpc/internal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.UserEarn{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.UserEarn{}) fail: %s", err)
	}
}

func Test_userEarnQuery(t *testing.T) {
	userEarn := newUserEarn(db)
	userEarn = *userEarn.As(userEarn.TableName())
	_do := userEarn.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(userEarn.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <user_earn> fail:", err)
		return
	}

	_, ok := userEarn.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from userEarn success")
	}

	err = _do.Create(&model.UserEarn{})
	if err != nil {
		t.Error("create item in table <user_earn> fail:", err)
	}

	err = _do.Save(&model.UserEarn{})
	if err != nil {
		t.Error("create item in table <user_earn> fail:", err)
	}

	err = _do.CreateInBatches([]*model.UserEarn{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <user_earn> fail:", err)
	}

	_, err = _do.Select(userEarn.ALL).Take()
	if err != nil {
		t.Error("Take() on table <user_earn> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <user_earn> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <user_earn> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <user_earn> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.UserEarn{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <user_earn> fail:", err)
	}

	_, err = _do.Select(userEarn.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <user_earn> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <user_earn> fail:", err)
	}

	_, err = _do.Select(userEarn.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <user_earn> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <user_earn> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <user_earn> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <user_earn> fail:", err)
	}

	_, err = _do.ScanByPage(&model.UserEarn{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <user_earn> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <user_earn> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <user_earn> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <user_earn> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <user_earn> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <user_earn> fail:", err)
	}
}
