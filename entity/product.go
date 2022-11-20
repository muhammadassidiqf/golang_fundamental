package entity

type Product struct{
	ID int
	Name string
	Price int
	Stock int
}

func (p Product) StatusStock() string  {
	var status string
	if p.Stock < 3 {
		status = "Stock sedikit lagi"
	} else if (p.Stock < 10) {
		status = "Stock barang terbatas"
	}

	return status
}