package models

import (
	"time"

	"github.com/yash91989201/go_cart/internal/database"
)

type User struct {
	ID              string        `json:"id"`
	CreatedAt       time.Time     `json:"created_at"`
	Name            string        `json:"name"`
	Email           string        `json:"email"`
	Password        string        `json:"password"`
	UserCart        []ProductUser `json:"usercart"`
	Address_Details []Address     `json:"address"`
	Order_Status    []Order       `json:"orders"`
}

type Product struct {
	ID     string  `json:"id"`
	Name   *string `json:"name"`
	Price  *uint64 `json:"price"`
	Rating *uint8  `json:"rating"`
	Image  *string `json:"image"`
}

type ProductUser struct {
	Product_ID   string  `bson:"_id"`
	Product_Name *string `json:"product_name" bson:"product_name"`
	Price        int     `json:"price"  bson:"price"`
	Rating       *uint   `json:"rating" bson:"rating"`
	Image        *string `json:"image"  bson:"image"`
}

type Address struct {
	Address_id string  `bson:"_id"`
	House      *string `json:"house_name" bson:"house_name"`
	Street     *string `json:"street_name" bson:"street_name"`
	City       *string `json:"city_name" bson:"city_name"`
	Pincode    *string `json:"pin_code" bson:"pin_code"`
}

type Order struct {
	Order_ID       string        `bson:"_id"`
	Order_Cart     []ProductUser `json:"order_list"  bson:"order_list"`
	Orderered_At   time.Time     `json:"ordered_on"  bson:"ordered_on"`
	Price          int           `json:"total_price" bson:"total_price"`
	Discount       *int          `json:"discount"    bson:"discount"`
	Payment_Method Payment       `json:"payment_method" bson:"payment_method"`
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod"     bson:"cod"`
}

type CreateUserReq struct {
	Name     string `json:"name" validate:"required,min=2,max=30"`
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required,min=4,max=32"`
}

type LoginUserReq struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password"`
}

func DBUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		Name:      dbUser.Name,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
	}
}
