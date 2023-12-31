// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"
	"fmt"
	"testing"

	"github.com/baiyz0825/school-share-buy-backend/apps/resources/cmd/rpc/internal/model"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Count{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Count{}) fail: %s", err)
	}
}

func Test_countQuery(t *testing.T) {
	count := newCount(db)
	count = *count.As(count.TableName())
	_do := count.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(count.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <count> fail:", err)
		return
	}

	_, ok := count.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from count success")
	}

	err = _do.Create(&model.Count{})
	if err != nil {
		t.Error("create item in table <count> fail:", err)
	}

	err = _do.Save(&model.Count{})
	if err != nil {
		t.Error("create item in table <count> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Count{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <count> fail:", err)
	}

	_, err = _do.Select(count.ALL).Take()
	if err != nil {
		t.Error("Take() on table <count> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <count> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <count> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <count> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Count{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <count> fail:", err)
	}

	_, err = _do.Select(count.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <count> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <count> fail:", err)
	}

	_, err = _do.Select(count.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <count> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <count> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <count> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <count> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Count{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <count> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <count> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <count> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <count> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <count> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <count> fail:", err)
	}
}
