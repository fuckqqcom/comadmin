package utils

import (
	"fmt"
	"github.com/xormplus/xorm"
	"log"
)

/**
开启事务
*/

func Insert(v interface{}, s *xorm.Session) bool {
	_, err := s.Insert(&v)
	defer func() {
		if r := recover(); r != nil {
			s.Rollback()
			log.Println(err)
			fmt.Println("recovered from ", r)
		}
	}()
	if err != nil {
		s.Rollback()
		log.Println(err)
		return false
	}
	return true
}

func Delete(v interface{}, s *xorm.Session, query, params interface{}) bool {
	_, err := s.Where(query, params).Delete(&v)
	defer func() {
		if r := recover(); r != nil {
			log.Println(err)
			fmt.Println("recovered from ", r)
		}
	}()
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
