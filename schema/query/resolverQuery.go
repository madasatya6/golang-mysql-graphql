package query
import (
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
	result,err := db.Query("select ID_PRO,PRO_NAME,QTE_IN_STOCK,ifnull(PRO_IMAGE,'') from Products")
	if err != nil{
		panic(err.Error())
	}

	for result.Next(){
		err = result.Scan(&a.IdPro,&a.ProName,&a.QteStock,&a.ProImg)
		if err != nil{
			panic(err.Error())
		}
		b = append(b,a)

	}


	return b , nil
}
