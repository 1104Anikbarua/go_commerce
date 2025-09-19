package database

import "fmt"

type TSProducts struct {
	Id          int     `json:"_id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	ImgUrl      string  `json:"imageUrl"`
}

var Products []TSProducts

func init() {
	fmt.Println("Print Init First")
	var prd1 = TSProducts{
		Id:          1,
		Title:       "Orange",
		Price:       99.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd2 = TSProducts{
		Id:          2,
		Title:       "Apple",
		Price:       59.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd3 = TSProducts{
		Id:          3,
		Title:       "Banana",
		Price:       49.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd4 = TSProducts{
		Id:    4,
		Title: "Pear",

		Price:       39.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd5 = TSProducts{
		Id:          5,
		Title:       "Mango",
		Price:       99.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	var prd6 = TSProducts{
		Id:          6,
		Title:       "Pineapple",
		Price:       29.99,
		Description: "An orange can refer to a fruit, a color, or even the name of a corporation. The sweet orange, Citrus sinensis, is the most common variety of the hybrid fruit, which is known for its juicy pulp and leathery rind. The color orange is named after the fruit and is a mix of yellow and red.",
		ImgUrl:      "https://media.gettyimages.com/id/185284489/photo/orange.jpg?s=612x612&w=gi&k=20&c=HZYbLyTgUgxD1WE-O-ltBo_Lui6pX6rQLHQJdYdyl_g=",
	}
	Products = append(Products, prd1)
	Products = append(Products, prd2)
	Products = append(Products, prd3)
	Products = append(Products, prd4)
	Products = append(Products, prd5)
	Products = append(Products, prd6)

}
