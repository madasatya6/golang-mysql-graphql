package types
import (

"github.com/graphql-go/graphql"

)

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
	Name:"Products",
	Fields:graphql.Fields{
		"ID_PRO":&graphql.Field{
			Type:graphql.Int,
		},
		"PRO_NAME":&graphql.Field{
			Type:graphql.String,
		},
		"QTE_IN_STOCK":&graphql.Field{
			Type:graphql.Int,
		},
		"PRICE":&graphql.Field{
			Type:graphql.Float,
		},
		"PRO_IMAGE":&graphql.Field{
			Type:graphql.String,
		},
	},
})