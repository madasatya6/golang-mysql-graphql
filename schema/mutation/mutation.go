package mutation

import (
	"github.com/madasatya6/golang-mysql-graphql/schema/types"
	"github.com/graphql-go/graphql"
)

var Mutation = graphql.NewObject(graphql.ObjectConfig{
	Name:"Mutation",
	Fields:graphql.Fields{
		"CreateProducts":&graphql.Field{
			Type:graphql.NewList(types.ProductTypes),
			//config param argument
			Args:graphql.FieldConfigArgument{
				"ID_PRO": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"PRO_NAME": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"QTE_IN_STOCK": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.Int),
				},
				"PRO_IMAGE": &graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				//array
				"Attributes": &graphql.ArgumentConfig{
					Type:graphql.NewList(Attribute),
				},
			},
			Resolve: CreateProductMutation,
		},
		// untuk membuat object lainya tinggal di ulang
	},
})

var Attribute = graphql.NewObject(graphql.ObjectConfig{
	Name:"MutationAttribute",
	Fields:graphql.Fields{
		"CreateAttribute":&graphql.Field{
			Type:graphql.NewList(types.ProductsAttributeTypes),
			Args:graphql.FieldConfigArgument{
				"ID_PRO":&graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
				"color":&graphql.ArgumentConfig{
					Type:graphql.NewNonNull(graphql.String),
				},
			},
		},
	},
})
