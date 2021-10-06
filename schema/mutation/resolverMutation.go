package mutation

import (
	"log"
	"github.com/madasatya6/golang-mysql-graphql/config"
	"github.com/graphql-go/graphql"
)

func CreateProductMutation(param graphql.ResolveParams) (interface{}, error) {

	db, err := config.GetConnection()
	if err != nil {
		panic(err.Error())
	}

	idpro := param.Args["ID_PRO"].(string)
	nama := param.Args["PRO_NAME"].(string)
	qty := param.Args["QTE_IN_STOCK"].(int)
	img := param.Args["PRO_IMAGE"].(string)
	
	_ , err = db.Query("INSERT INTO Products values (?,?,?,?)",idpro,nama,qty,img)
	if err != nil {
		log.Println(err.Error())
	}

	//masukan ke array

	return nil, err
}
