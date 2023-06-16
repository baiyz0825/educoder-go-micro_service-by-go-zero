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
	err := db.AutoMigrate(&model.Classification{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Classification{}) fail: %s", err)
	}
}

func Test_classificationQuery(t *testing.T) {
	classification := newClassification(db)
	classification = *classification.As(classification.TableName())
	_do := classification.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(classification.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <classification> fail:", err)
		return
	}

	_, ok := classification.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from classification success")
	}

	err = _do.Create(&model.Classification{})
	if err != nil {
		t.Error("create item in table <classification> fail:", err)
	}

	err = _do.Save(&model.Classification{})
	if err != nil {
		t.Error("create item in table <classification> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Classification{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <classification> fail:", err)
	}

	_, err = _do.Select(classification.ALL).Take()
	if err != nil {
		t.Error("Take() on table <classification> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <classification> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <classification> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <classification> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Classification{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <classification> fail:", err)
	}

	_, err = _do.Select(classification.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <classification> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <classification> fail:", err)
	}

	_, err = _do.Select(classification.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <classification> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <classification> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <classification> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <classification> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Classification{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <classification> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <classification> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <classification> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <classification> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <classification> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <classification> fail:", err)
	}
}
