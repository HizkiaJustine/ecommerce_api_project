package models

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {

	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    FirstName    *string            `bson:"first_name" json:"first_name" validate:"required"`
    LastName     *string            `bson:"last_name" json:"last_name" validate:"required"`
    Password     *string            `bson:"password" json:"password" validate:"required"`
    Email        *string            `bson:"email" json:"email" validate:"required,email"`
    PhoneNumber  *string            `bson:"phone_number" json:"phone_number" validate:"required"`
    Token        *string            `bson:"token" json:"token"`
    RefreshToken *string            `bson:"refresh_token" json:"refresh_token"`
    CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
    UserID       *string            `bson:"user_id" json:"user_id"`
    UserCart     []ProductUser      `bson:"user_cart" json:"user_cart"`
    AdressDetails []Address         `bson:"address_details" json:"address_details"`
    OrderStatus  []Order            `bson:"order_status" json:"order_status"`
}

type Product struct {
	ProductID    primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName  *string            `bson:"product_name" json:"product_name"`
	Price        *uint64            `bson:"price" json:"price"`
	Rating       *uint8             `bson:"rating" json:"rating"`
	Image        *string            `bson:"image" json:"image"`
}

type ProductUser struct {
	ProductID   primitive.ObjectID `bson:"product_id" json:"product_id"`
	ProductName *string            `bson:"product_name" json:"product_name"`
	Price       int                `bson:"price" json:"price"`
	Rating      *uint              `bson:"rating" json:"rating"`
	Image       *string            `bson:"image" json:"image"`
}

type Address struct {
	AddressID primitive.ObjectID `bson:"address_id" json:"address_id"`
	House	  *string            `bson:"house" json:"house"`
	Street    *string            `bson:"street" json:"street"`
	City      *string            `bson:"city" json:"city"`
}

type Order struct {
	OrderID       primitive.ObjectID `bson:"order_id" json:"order_id"`
	OrderCart     []ProductUser      `bson:"order_cart" json:"order_cart"`
	OrderedAt     time.Time          `bson:"ordered_at" json:"ordered_at"`
	Price         uint64             `bson:"price" json:"price"`
	Discount      *int               `bson:"discount" json:"discount"`
	PaymentMethod Payment            `bson:"payment_method" json:"payment_method"`
}

type Payment struct {
	Digital bool `bson:"digital" json:"digital"`
	COD     bool `bson:"cod" json:"cod"`
}