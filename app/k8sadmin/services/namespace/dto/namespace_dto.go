/*
*
User: cr-mao
Date: 2024/2/22 14:05
Email: crmao@qq.com
Desc: namespace_dto.go
*/
package dto

type Namespace struct {
	Name      string `json:"name"`
	CreatedAt int64  `json:"created_at"`
	Status    string `json:"status"`
}
