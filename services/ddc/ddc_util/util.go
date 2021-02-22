/*
 * Copyright 2020 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// util.go - define the utilities for api package of RDS service
package ddc_util

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", fmt.Errorf("error secrete key")
	}

	crypted, err := crypto.EBCEncrypto([]byte(sk[:16]), []byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypted), nil
}

// 拷贝字段内容工具
func SimpleCopyProperties(dst, src interface{}) (err error) {
	// 防止意外panic
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprintf("%v", e))
		}
	}()

	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)

	//dstType, dstValue := reflect.TypeOf(dst), reflect.ValueOf(dst)
	//srcType, srcValue := reflect.TypeOf(src), reflect.ValueOf(src)
	//
	//// dst必须结构体指针类型
	//if dstType.Kind() != reflect.Ptr || dstType.Elem().Kind() != reflect.Struct {
	//	return errors.New("dst type should be a struct pointer")
	//}
	//
	//// src必须为结构体或者结构体指针
	//if srcType.Kind() == reflect.Ptr {
	//	srcType, srcValue = srcType.Elem(), srcValue.Elem()
	//}
	//if srcType.Kind() != reflect.Struct {
	//	return errors.New("src type should be a struct or a struct pointer")
	//}
	//
	//// 取具体内容
	//dstType, dstValue = dstType.Elem(), dstValue.Elem()
	//
	//// 属性个数
	//propertyNums := dstType.NumField()
	//
	//for i := 0; i < propertyNums; i++ {
	//	// 属性
	//	property := dstType.Field(i)
	//	// 待填充属性值
	//	propertyValue := srcValue.FieldByName(property.Name)
	//	// 数组属性同名时
	//	if propertyValue.Kind() == reflect.Slice &&
	//		property.Type.Elem().Name() == propertyValue.Type().Elem().Name() {
	//		fmt.Println(propertyValue)
	//		// 嵌套类型使用json转换
	//		fieldJson, err := json.Marshal(propertyValue)
	//		if err != nil {
	//			return err
	//		}
	//		fieldValue := reflect.New(property.Type)
	//		err = json.Unmarshal(fieldJson, &fieldValue)
	//		fmt.Println(fieldValue)
	//		dstValue.Field(i).Set(fieldValue)
	//		continue
	//	}
	//	// src没有这个属性 || 属性同名但类型不同
	//	if !propertyValue.IsValid() || property.Type != propertyValue.Type() {
	//		continue
	//	}
	//
	//	if dstValue.Field(i).CanSet() {
	//		dstValue.Field(i).Set(propertyValue)
	//	}
	//}
	//
	//return nil
}
