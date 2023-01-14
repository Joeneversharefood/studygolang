package main

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	_ "github.com/gomodule/redigo/redis"
	_ "github.com/jmoiron/sqlx"
	_ "log"
	"net/http"
	_ "net/http"
	"time"
)

var jwtkey = []byte("my_secret_key")

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func signUp(w http.ResponseWriter, r *http.Request) {

	var creds Credential

	err := json.NewDecoder(r.Body).Decode(&creds)

	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	conn, err := redis.Dial("tcp", "localhost:6379")

	defer conn.Close()

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	result, err := conn.Do("get", creds.Username)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if nil != result {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("register failed, %s already been used", creds.Username)))
		return
	}

	_, err = conn.Do("set", creds.Username, creds.Password)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write([]byte("register success,now you can log in"))

}

func signIn(w http.ResponseWriter, r *http.Request) {
	var creds Credential

	err := json.NewDecoder(r.Body).Decode(&creds)

	if nil != err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	conn, err := redis.Dial("tcp", "localhost:6379")

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("err = %v\n", err)
		return
	}

	result, err := redis.String(conn.Do("get", creds.Username))

	if creds.Password != result {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("user not exit"))
		return
	}

	expiretime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiretime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expiretime,
	})

	w.Write([]byte("signin success,you can go nuts now"))

	defer conn.Close()

}

func index(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("token")

	if nil != err {
		if http.ErrNoCookie == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("your cookie has already expire, please sign in again"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenstr := ck.Value

	claims := Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstr, &claims, func(token *jwt.Token) (interface{}, error) { return jwtkey, nil })

	if nil != err {
		if jwt.ErrSignatureInvalid == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("your token is illegal, please sign up for a new account"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("welcome ! %s", claims.Username)))

}

func refresh(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("token")

	if nil != err {
		if http.ErrNoCookie == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("your cookie has already expire, please sign in again"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenstr := ck.Value

	claims := Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstr, &claims, func(token *jwt.Token) (interface{}, error) { return jwtkey, nil })

	if nil != err {
		if jwt.ErrSignatureInvalid == err {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expireTime := time.Now().Add(5 * time.Minute)

	claims.ExpiresAt = expireTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(jwtkey)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenstring,
		Expires: expireTime,
	})

	w.Write([]byte(fmt.Sprintf("your cookie's expire time has been extend to : \n%v", expireTime)))

}

func deregister(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("token")

	if nil != err {
		if http.ErrNoCookie == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("please sign in first"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jwttoken := ck.Value

	var creds Credential

	err = json.NewDecoder(r.Body).Decode(&creds)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("decode json body failed\n")
		return
	}

	claim := Claims{}

	token, err := jwt.ParseWithClaims(jwttoken, &claim, func(token *jwt.Token) (interface{}, error) { return jwtkey, nil })

	if nil != err {
		if jwt.ErrSignatureInvalid == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("jwt verify failed,please sign in"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("unknown parse error\n")
		return
	}

	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("jwt verify failed,please sign in"))
	}

	conn, err := redis.Dial("tcp", "localhost:6379")

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		retstr := fmt.Sprintf("conn to redis failed\n")
		fmt.Printf(retstr)
		return
	}

	result, err := conn.Do("get", creds.Username)

	defer conn.Close()

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("query redis failed\n")
		return
	}

	if nil == result {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("user not exits"))
		return
	}

	ret, err := redis.Int(conn.Do("del", creds.Username))

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Printf("del user data from redis failed\n")
		return
	}

	if 1 != ret {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("deregister failed"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deregister success~"))

}

func testcookie(w http.ResponseWriter, r *http.Request) {
	ck, err := r.Cookie("token")

	if nil != err {
		if http.ErrNoCookie == err {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("your cookie has already expire, please sign in again"))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenstr := ck.Value

	claims := Claims{}

	tkn, err := jwt.ParseWithClaims(tokenstr, &claims, func(token *jwt.Token) (interface{}, error) { return jwtkey, nil })

	if nil != err {
		if jwt.ErrSignatureInvalid == err {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expireTime := time.Now().Add(5 * time.Minute)

	claims.ExpiresAt = expireTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenstring, err := token.SignedString(jwtkey)

	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token3",
		Value:   tokenstring,
		Expires: expireTime,
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "token1",
		Value:   tokenstring,
		Expires: expireTime,
	})

	w.Write([]byte(fmt.Sprintf("your cookie's expire time has been extend to : \n%v", expireTime)))

}

func main() {

	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/signin", signIn)
	http.HandleFunc("/refresh", refresh)
	http.HandleFunc("/index", index)
	http.HandleFunc("/deregister", deregister)
	http.HandleFunc("/testcookie", testcookie)

	http.ListenAndServe("localhost:8000", nil)
}
