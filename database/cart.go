package database

import "errors"

var (
	ErrCantFindProduct        = errors.New("Cannot find the Product")
	ErrCantDecodeProduct      = errors.New("Cannot find the Product")
	ErrInvalidUserID          = errors.New("User is not valid")
	ErrCantUpdateUser         = errors.New("Cannot update Usert")
	ErrCantRemoveItemFromCart = errors.New("Cannot remove this Item from the Cart")
	ErrCantGetItem            = errors.New("Unable to get this Item from the Cart")
	ErrCantBuyCartItem        = errors.New("Cannot buy this Item from the Cart")
)

func AddToCart()       {}
func RemoveItem()      {}
func GetItemFromCart() {}
func BuyItemFromCart() {}
func InstantBuy()      {}
