package auth

import (
	"awesome-start/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"strings"
	"time"
)

type RequestUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userDB struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

const JwtSecret = "asecret"

func Login(c *fiber.Ctx) error {
	var reqUser RequestUser
	err := c.BodyParser(&reqUser)
	if err != nil {
		return nil
	}

	filter := bson.M{"username": reqUser.Username}
	var dbUser RequestUser
	err = db.UserCollection.FindOne(db.Ctx, filter).Decode(&dbUser)
	if err != nil {
		return nil
	}

	if !isValidPassword(reqUser.Password, dbUser.Password) {
		return nil
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = dbUser.Username
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7) // a week

	s, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return nil
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
		"user":  dbUser.Username,
	})

	return nil
}

func isValidPassword(providedPassword, storedPassword string) bool {

	return providedPassword == storedPassword
}

func Register(c *fiber.Ctx) error {
	var body RequestUser
	err := c.BodyParser(&body)

	if err != nil {
		return nil
	}

	newUser := createUser(body.Username, body.Password)

	_, err = db.UserCollection.InsertOne(db.Ctx, newUser)
	if err != nil {
		return nil
	}

	return nil
}

func createUser(userName string, password string) *userDB {
	return &userDB{
		Username: strings.TrimSpace(userName),
		Password: strings.TrimSpace(password),
	}
}
