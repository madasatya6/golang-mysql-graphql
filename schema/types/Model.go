package types

import "github.com/madasatya6/golang-mysql-graphql/config"

type Product struct {
	IdPro string `json:"ID_PRO"`
	ProName string `json:"PRO_NAME"`
	QteStock int `json:"QTE_IN_STOCK"`
	ProImg string `json:"PRO_IMAGE"`
	Attributes []ProductsAttribute
}

type ProductsAttribute struct{
	ID string `json:"id"`
	ID_PRO string `json:"ID_PRO"`
	Color string `json:"color"`
}

func ProductsAttributeFind(id string) (interface{},error) {
	var a ProductsAttribute
	var b []interface{}
	db ,err:= config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	b = b[:0]
	result,err := db.Query("select id, ID_PRO, color from products_attribute where ID_PRO='?'", id)
	if err != nil{
		panic(err.Error())
	}

	for result.Next(){
		err = result.Scan(&a.ID,&a.ID_PRO,&a.Color)
		if err != nil{
			panic(err.Error())
		}
		b = append(b,a)

	}
	
	return b , nil
}