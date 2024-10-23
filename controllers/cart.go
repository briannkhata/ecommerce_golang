package controllers

import (
	"context"
	"ecommerce/database"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		{
			if productQueryID == "" {
				log.Println("product is is empty")
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
				return
			}
		}

		userQueryID := c.Query("userID")
		{
			if userQueryID == "" {
				log.Println("user is is empty")
				_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
				return
			}
		}

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)
		defer cancel()

		err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "Successfully added to the cart")
	}
}

func (app *Application) RemoveItem() gin.HandlerFunc {
	return (c *gin.Context){

		productQueryID := c.Query("id")
		
			if productQueryID == "" {
				log.Println("product is is empty")

				_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
				return
			}
		

		userQueryID := c.Query("userID")
		
			if userQueryID == "" {
				log.Println("user is is empty")

				_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
				return
			}
		

		productID, err := primitive.ObjectIDFromHex(productQueryID)

		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)
		defer cancel()

		err = database.RemoveCartItem(ctx,app.prodCollection,app.userCollection,productID,userQueryID)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, err)
			return
		}

		c.IndentedJSON(http.StatusOK, "Successfully removed from the cart")
	}

}

func GetItemFromCart() gin.HandlerFunc {

}

func BuyFromCart() gin.HandlerFunc {

}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context){

	productQueryID := c.Query("id")
		
		if productQueryID == "" {
			log.Println("product is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
			return
		}
	

	userQueryID := c.Query("userID")
	
		if userQueryID == "" {
			log.Println("user is is empty")

			_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
			return
		}
	

	productID, err := primitive.ObjectIDFromHex(productQueryID)

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var ctx, cancel = context.WithTimeOut(context.Background(), 5*time.Second)
	defer cancel()

	err = database.InstantBuy(ctx,app.prodCollection,app.userCollection,productID,userQueryID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	c.IndentedJSON(http.StatusOK, "Successfully placed to the order")

	}

}
