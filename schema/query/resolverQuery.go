package query
import (
	"log"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/madasatya6/golang-mysql-graphql/schema/types"
	"github.com/madasatya6/golang-mysql-graphql/config"
)

func ProductResolve(param graphql.ResolveParams) (interface{},error){
	var a types.Product
	var b []types.Product
	db ,err:= config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	b = b[:0]
	result,err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMAGE,'') from products")
	if err != nil{
		panic(err.Error())
	}

	for result.Next(){
		err = result.Scan(&a.IdPro, &a.ProName, &a.QteStock, &a.ProImg)
		if err != nil{
			log.Println(err.Error())
		}
		
		a.Attributes, err = ProductsAttributeFind(a.IdPro)
		if err != nil {
			 log.Println(err.Error())
		}

		b = append(b,a)
	}

	return b , nil
}

func ProductsAttributeResolve(param graphql.ResolveParams) (interface{},error) {
	var a types.ProductsAttribute
	var b []types.ProductsAttribute
	db ,err:= config.GetConnection()
	if err != nil {
		panic(err.Error())
	}
	result,err := db.Query("select id, ID_PRO, color from products_attribute")
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

func ProductsAttributeFind(id string) ([]types.ProductsAttribute, error) {
	
	var a types.ProductsAttribute
	var b []types.ProductsAttribute

	db ,err := config.GetConnection()
	if err != nil {
		log.Println(err.Error())
		return b, err 
	}

	query := fmt.Sprintf("select id, ID_PRO, color from products_attribute where ID_PRO='%s'", id)
	result, err := db.Query(query)
	if err != nil{
		log.Println(err.Error())
		return b, err 
	}

	for result.Next(){
		err = result.Scan(&a.ID,&a.ID_PRO,&a.Color)
		if err != nil{
			log.Println(err.Error())
		}
		b = append(b,a)

	}
	
	return b , nil
}