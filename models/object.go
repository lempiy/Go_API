package models

import (
	"errors"

	"fmt"

	"strconv"

	"github.com/astaxie/beego"
)

var (
	Objects map[string]*Object
)

type Object struct {
	Id         int    `json:"id, omitempty"`
	Score      int64  `json:"score"`
	PlayerName string `json:"player_name"`
}

func AddOne(object Object) (ObjectId string) {
	beego.Info(object)
	retID, err := O.Insert(&object)
	if err != nil {
		beego.Error(err)
	}
	return fmt.Sprintf("%d", retID)
}

func GetOne(ObjectId string) (*Object, error) {
	ID, err := strconv.Atoi(ObjectId)
	object := Object{Id: ID}
	err = O.Read(&object)
	return &object, err
}

func GetAll() map[string]*Object {
	return Objects
}

func Update(ObjectId string, Score int64) (err error) {
	if v, ok := Objects[ObjectId]; ok {
		v.Score = Score
		return nil
	}
	return errors.New("ObjectId Not Exist")
}

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}
