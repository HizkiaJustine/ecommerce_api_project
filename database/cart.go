package database

import (

)

var (
	ErrCantFindProduct = errors.New("product not found")
	ErrCantDecodeProducts = errors.New("can't decode products")
	ErrUserIdIsNotValid = errors.New("user id is not valid")
	ErrCantUpdateUser = errors.New("can't update user")
	ErrCantRemoveItemCart = errors.New("can't remove item from cart")
	ErrCantGetItem = errors.New("can't get item")
	ErrCantBuyCartItem = errors.New("can't buy item from cart")
)

func AddProductToCart() {

}

func RemoveCartItem() {

}

func BuyItemFromCart() {

}

func InstantBuyer() {

}