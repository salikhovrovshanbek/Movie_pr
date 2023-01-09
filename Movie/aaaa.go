package main

//
//import (
//	"errors"
//	"strconv"
//	"time"
//)
//
//func generateToken(userId int64, interval time.Duration) (string, error) {
//	now := time.Now()
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
//		IssuedAt:  now.Unix(),
//		ExpiresAt: now.Add(interval).Unix(),
//		Subject:   strconv.Itoa(int(userId)),
//	})
//	return token.SignedString([]byte(signingKey))
//}
//
//const signingKey = "vsdgvtrbvsdsac btropfclbr"
//
//func parseToken(tokenStr string) (int, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, errors.New("invalid signing method")
//		}
//		return []byte(signingKey), nil
//	})
//	if err != nil {
//		return 0, err
//	}
//
//	claims, ok := token.Claims.(*jwt.StandardClaims)
//	if !ok {
//		return 0, errors.New("token claims are not type of  jwt.StandardClaims")
//	}
//	return strconv.Atoi(claims.Subject)
//}
