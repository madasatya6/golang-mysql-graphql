package types

type Product struct {
	IdPro string `json:"ID_PRO"`
	ProName string `json:"PRO_NAME"`
	QteStock int `json:"QTE_IN_STOCK"`
	ProImg string `json:"PRO_IMAGE"`
}

type ProductsAttribute struct{
	ID string `json:"id"`
	ID_PRO string `json:"ID_PRO"`
	Color string `json:"color"`
}