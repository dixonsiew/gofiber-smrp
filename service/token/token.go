package token

import (
	"fmt"
	"smrp/model"
	userService "smrp/service/user"
	"smrp/utils"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func DecodeToken(c *fiber.Ctx) (string, int, error) {
    tokenStr := c.Get("Authorization")
    tokenStr = strings.ReplaceAll(tokenStr, "Bearer ", "")
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(utils.JWT_SECRET), nil
    })

    if err != nil {
        utils.LogError(err)
        return "", 0, err
    }

    claims, ok := token.Claims.(*jwt.MapClaims)
    if !ok {
        return "", 0, fmt.Errorf("could not parse claims")
    }

    sub := (*claims)["subject"].(string)
    username := (*claims)["username"].(string)
    id, _ := strconv.Atoi(sub)
    return username, id, nil
}

func GenerateAccessToken(user model.User) (string, error) {
    claims := jwt.MapClaims{
        "username": user.Username,
        "subject":  fmt.Sprintf("%d", user.Id),
        "exp":      time.Now().Add(time.Hour * 720).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(utils.JWT_SECRET))
    if err != nil {
        utils.LogError(err)
        return "", err
    }

    return t, nil
}

func GenerateRefreshToken(user model.User) (string, error) {
    claims := jwt.MapClaims{
        "username": user.Username,
        "subject":  fmt.Sprintf("%d", user.Id),
        "exp":      time.Now().Add(time.Hour * 87600).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    t, err := token.SignedString([]byte(utils.JWT_SECRET))
    if err != nil {
        utils.LogError(err)
        return "", err
    }

    return t, nil
}

func ResolveRefreshToken(encoded string) (*model.User, error) {
    _, id, err := decodeRefreshToken(encoded)
    if id == 0 || err != nil {
        return nil, fmt.Errorf("refresh token not found")
    }

    user, err := getUserFromRefreshTokenPayload(id)
    if err != nil {
        return nil, fmt.Errorf("refresh token malformed")
    }

    return user, nil
}

func CreateAccessTokenFromRefreshToken(refresh string) (fiber.Map, error) {
    user, err := ResolveRefreshToken(refresh)
    if err != nil {
        return nil, err
    }

    token, err := GenerateAccessToken(*user)
    if err != nil {
        return nil, err
    }

    return fiber.Map{
        "user":  &user,
        "token": token,
    }, nil
}

func decodeRefreshToken(tokenStr string) (string, int, error) {
    token, err := jwt.ParseWithClaims(tokenStr, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
        return []byte(utils.JWT_SECRET), nil
    })

    if err != nil {
        utils.LogError(err)
        return "", 0, err
    }

    claims, ok := token.Claims.(*jwt.MapClaims)
    if !ok {
        return "", 0, fmt.Errorf("could not parse claims")
    }

    sub := (*claims)["subject"].(string)
    username := (*claims)["username"].(string)
    id, _ := strconv.Atoi(sub)
    return username, id, nil
}

func getUserFromRefreshTokenPayload(id int) (*model.User, error) {
    return userService.FindById(id)
}
