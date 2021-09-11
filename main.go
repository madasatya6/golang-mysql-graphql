package main

import (
	"github.com/madasatya6/golang-mysql-graphql/schema/query"
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/madasatya6/golang-mysql-graphql/schema/mutation"
	"net/http"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:   schema    ,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}

func main(){
	var Schame,err = graphql.NewSchema(graphql.SchemaConfig{
		Query:query.RootQuery,
		Mutation: mutation.Mutation,
	})
	if err != nil{
		panic(err.Error())
	}
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("Query"), Schame)
		json.NewEncoder(w).Encode(result)
	})
	http.ListenAndServe(":8088", nil)

}
