package types
import (

"github.com/graphql-go/graphql"

)

var ProductTypes = graphql.NewObject(graphql.ObjectConfig{
	Name:"Products",
	Fields:graphql.Fields{
		"ID_PRO":&graphql.Field{
			Type:graphql.String,
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
		//implementasi array
		"Attributes":&graphql.Field{
			Type:graphql.NewList(graphql.NewObject(graphql.ObjectConfig{
				Name:"Attributes",
				Fields:graphql.Fields{
					"id":&graphql.Field{
						Type:graphql.Int,
					},
					"ID_PRO":&graphql.Field{
						Type:graphql.String,
					},
					"color":&graphql.Field{
						Type:graphql.String,
					},
				},
			})),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				product, _ := params.Source.(Product)
				attributes, err := ProductsAttributeFind(product.IdPro)
				if err != nil {
					return nil, err
				}
				return attributes, nil
			},
		},
	},
})


var ProductsAttributeTypes = graphql.NewObject(graphql.ObjectConfig{
	Name:"ProductsAttribute",
	Fields:graphql.Fields{
		"id":&graphql.Field{
			Type:graphql.Int,
		},
		"ID_PRO":&graphql.Field{
			Type:graphql.String,
		},
		"color":&graphql.Field{
			Type:graphql.String,
		},
	},
})

