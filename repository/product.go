package repository

type TIProductRepo interface {
	Create(product TSProducts) (*TSProducts, error)
	Get(productId int) (*TSProducts, error)
	Update(payload TSProducts) (*TSProducts, error)
	Delete(productId int) error
	List() ([]*TSProducts, error)
}

type TSProducts struct {
	Id          int     `json:"_id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imageUrl"`
}

type TSProductRepo struct {
	products []*TSProducts
}

func NewProductRepo() TIProductRepo {
	repo := &TSProductRepo{}
	generateInitialProduct(repo)
	return repo
}

func (r *TSProductRepo) Create(product TSProducts) (*TSProducts, error) {
	product.Id = len(r.products) + 1
	r.products = append(r.products, &product)
	return &product, nil
}
func (r *TSProductRepo) Get(productId int) (*TSProducts, error) {
	for _, product := range r.products {
		if productId == product.Id {
			return product, nil
		}
	}
	return nil, nil
}
func (r *TSProductRepo) Update(payload TSProducts) (*TSProducts, error) {
	for idx, product := range r.products {
		if payload.Id == product.Id {
			r.products[idx] = &payload
		}
	}
	return nil, nil
}
func (r *TSProductRepo) Delete(productId int) error {
	var tempProduct []*TSProducts
	for _, product := range r.products {
		if productId != product.Id {
			tempProduct = append(tempProduct, product)
			r.products = tempProduct
		}
	}
	return nil
}
func (r *TSProductRepo) List() ([]*TSProducts, error) {
	return r.products, nil
}

func generateInitialProduct(r *TSProductRepo) {
	var prd1 = &TSProducts{
		Id:          1,
		Title:       "Orange",
		Price:       99.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd2 = &TSProducts{
		Id:          2,
		Title:       "Apple",
		Price:       59.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd3 = &TSProducts{
		Id:          3,
		Title:       "Banana",
		Price:       49.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd4 = &TSProducts{
		Id:    4,
		Title: "Pear",

		Price:       39.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd5 = &TSProducts{
		Id:          5,
		Title:       "Mango",
		Price:       99.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd6 = &TSProducts{
		Id:          6,
		Title:       "Pineapple",
		Price:       29.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	r.products = append(r.products, prd1)
	r.products = append(r.products, prd2)
	r.products = append(r.products, prd3)
	r.products = append(r.products, prd4)
	r.products = append(r.products, prd5)
	r.products = append(r.products, prd6)

}
