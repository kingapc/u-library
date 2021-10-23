package main

import (
	"log"

	"github.com/gin-gonic/gin"
	handler "github.com/rpinedafocus/u-library/pkg/handler"
	security "github.com/rpinedafocus/u-library/pkg/middleware"
)

/**********	MAIN	************/
func main() {

	var router = gin.Default()
	security.Init()
	router.POST("/login", handler.LoginController)
	router.POST("/logout", security.TokenAuthMiddleware(), handler.LogoutController)

	router.GET("/books", security.TokenAuthMiddleware(), handler.FetchAllBooksController)
	router.POST("/books/create", handler.CreateBookController)
	router.GET("/books/:id", handler.FetchBookByIdController)

	router.POST("/bookings/create", handler.CreateBookingController)
	router.POST("/bookings/release/:id", handler.ReleaseBookingController)

	router.POST("/users/create", handler.CreateUserController)
	router.POST("/roles/create", handler.CreateRoleController)
	router.POST("/genres/create", handler.CreateGenreController)
	router.POST("/authors/create", handler.CreateAuthorController)

	router.POST("/rent/create", handler.CreateRentController)
	router.POST("/rent/return/:id", handler.ReturnRentedBookController)

	router.GET("/common/mybooks", security.TokenAuthMiddleware(), handler.FetchBookingRentBooks) //User inquiry his books
	router.GET("/common/mybooks/:id", handler.FetchBookingRentBooks)                             //admin inquiry the user's books

	//router.POST("/logins", Login)
	// router.GET("/authors", getAuthors)
	// router.GET("/books", getBooks)
	// router.GET("/books/:id", getBook)
	// router.GET("/books/mybooks", getReserveRent)
	// router.POST("/login", Login)
	// router.POST("/manage/rent", setRent)
	// router.POST("/manage/reserve", setReserve)
	// router.POST("/manage/return/:id", setReturn)
	// router.POST("/create/user", createUser)
	// router.POST("/create/book", createBook)
	// router.POST("/todo", CreateTodo)
	// router.POST("/logout", Logout)

	// router.POST("/login", Login)
	// router.POST("/todo", TokenAuthMiddleware(), CreateTodo)
	// router.POST("/logout", TokenAuthMiddleware(), Logout)
	// router.POST("/refresh", Refresh)

	// router.POST("/login", Login)
	// router.POST("/todo", CreateTodo)
	// router.POST("/logout", Logout)

	log.Fatal(router.Run("localhost:8080"))
}

// func EndPointFactory(c *gin.Context) {

// 	type All struct {
// 		a []*model.FetchBook
// 		b *model.UserEntity
// 	}
// 	var ent All
// 	var response string
// 	var test All
// 	endpoint := c.Request.URL.Path

// 	fmt.Print(endpoint)
// 	switch endpoint {
// 	case "/books":
// 		ent.a, response = handler.FetchAllBooksController(c)
// 		test.a = ent.a
// 		// case "/users/create":
// 		// 	ent.b, response = handler.CreateUserController(c)
// 	}

// 	if response == "" {
// 		c.IndentedJSON(http.StatusOK, test)
// 	} else {
// 		c.IndentedJSON(http.StatusBadRequest, response)
// 	}
// }

// func CreateUser(c *gin.Context) {
// 	ent, response := handler.CreateUserController(c)

// 	if response == "" {
// 		c.IndentedJSON(http.StatusOK, ent)
// 	} else {
// 		c.IndentedJSON(http.StatusBadRequest, response)
// 	}
// }

// func CreateRol(c *gin.Context) {
// 	ent, response := handler.CreateRolController(c)

// 	if response == "" {
// 		c.IndentedJSON(http.StatusOK, ent)
// 	} else {
// 		c.IndentedJSON(http.StatusBadRequest, response)
// 	}
// }

/*********** END POINTS	***********/
// func getAuthors(c *gin.Context) {

// 	authors := entities.GetRows()
// 	fmt.Print(authors)
// 	c.IndentedJSON(http.StatusOK, authors)
// }

// //Get all books
// func getBooks(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {

// 			books := entities.GetBooks()
// 			c.IndentedJSON(http.StatusOK, books)
// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// //Rent the book
// func setRent(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {
// 			var nr manage.Status

// 			if err := c.BindJSON(&nr); err != nil {
// 				return
// 			}

// 			result := manage.GetRent(nr)

// 			if result {
// 				c.IndentedJSON(http.StatusCreated, gin.H{"message": "Rent successfully created"})
// 			} else {
// 				c.IndentedJSON(http.StatusCreated, gin.H{"message": "Unable to rent the book"})
// 			}
// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// //Reserve the book
// func setReserve(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {

// 			books := entities.GetBooks()
// 			c.IndentedJSON(http.StatusOK, books)
// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// func getBook(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, "/books/:id")

// 		if isAccess {
// 			var b entities.Books

// 			id, fail := strconv.ParseUint(c.Param("id"), 10, 32)

// 			if fail != nil {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to parse the id"})
// 			}

// 			b, err := entities.GetBook(int(id))

// 			if !err {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to get the detail"})
// 			} else {
// 				c.IndentedJSON(http.StatusOK, b)
// 			}
// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// func getReserveRent(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {
// 			fmt.Print(isLogin1.USER_NAME)
// 			b, err := entities.GetMyReservesRents(isLogin1.USER_NAME)

// 			if !err {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to get the reserves and rents"})
// 			} else {
// 				c.IndentedJSON(http.StatusOK, b)
// 			}
// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// //Reserve the book
// func setReturn(c *gin.Context) {

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, "/manage/return/:id")

// 		if isAccess {

// 			id, fail := strconv.ParseUint(c.Param("id"), 10, 32)

// 			if fail != nil {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to parse the id"})
// 			}

// 			result := manage.ReturBook(int(id))

// 			if result {
// 				c.IndentedJSON(http.StatusOK, gin.H{"message": "Return succesfully"})
// 			} else {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to return the book"})
// 			}

// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// //Login
// func Login(c *gin.Context) {
// 	var credentials Credentials

// 	if err := c.BindJSON(&credentials); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, "Invalid json provide")
// 		return
// 	}

// 	isLogin, errc := security.Login(credentials.USER, credentials.PASSWORD)

// 	if errc {
// 		c.JSON(http.StatusUnprocessableEntity, "Error")
// 		return
// 	}

// 	isLogin1.FULL_NAME = isLogin.FULL_NAME
// 	isLogin1.USER_NAME = isLogin.USER_NAME
// 	isLogin1.ROLE_ID = isLogin.ROLE_ID

// 	if credentials.USER != isLogin1.USER_NAME || credentials.PASSWORD != "root" {
// 		c.JSON(http.StatusUnauthorized, "Please prove valid login details")
// 		return
// 	}

// 	ts, err := CreateToken(isLogin1.USER_NAME)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}

// 	tokenD = ts
// 	saveErr := CreateAuth(isLogin1.USER_NAME, ts)
// 	if saveErr != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 	}

// 	tokens := map[string]string{
// 		"access_token":  ts.AccessToken,
// 		"refresh_token": ts.RefreshToken,
// 	}

// 	c.JSON(http.StatusOK, tokens)
// }

// func createUser(c *gin.Context) {

// 	var nu security.Users

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {

// 			if err := c.BindJSON(&nu); err != nil {
// 				panic(err)
// 				return
// 			}

// 			nur, err := security.CreateUser(nu)

// 			if err {
// 				c.IndentedJSON(http.StatusOK, nur)
// 			} else {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to create the user"})
// 			}

// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// func createBook(c *gin.Context) {

// 	var nb entities.Books

// 	if isLogin1.ROLE_ID != 0 {
// 		isAccess := security.GetAccess(isLogin1.ROLE_ID, c.Request.URL.Path)

// 		if isAccess {

// 			if err := c.BindJSON(&nb); err != nil {
// 				panic(err)
// 				return
// 			}

// 			nur, err := entities.CreateBook(nb)

// 			if err {
// 				c.IndentedJSON(http.StatusOK, nur)
// 			} else {
// 				c.IndentedJSON(http.StatusForbidden, gin.H{"message": "Unable to create the book"})
// 			}

// 		} else {
// 			c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You don't have access to this option"})
// 		}
// 	} else {
// 		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "You must log in before to access."})
// 	}
// }

// func CreateToken(user string) (*TokenDetails, error) {
// 	td := &TokenDetails{}
// 	td.AtExpires = time.Now().Add(time.Minute * 1).Unix()
// 	td.AccessUuid = uuid.NewV4().String()

// 	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
// 	td.RefreshUuid = uuid.NewV4().String()

// 	var err error
// 	os.Setenv("ACCESS_SECRET", GoDotEnvVariable("ACCESS_SECRET"))
// 	atClaims := jwt.MapClaims{}
// 	atClaims["authorized"] = true
// 	atClaims["access_uuid"] = td.AccessUuid
// 	atClaims["user_id"] = user
// 	atClaims["exp"] = td.AtExpires

// 	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

// 	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

// 	if err != nil {
// 		return nil, err
// 	}

// 	//Refreh Token
// 	os.Setenv("REFRESH_SECRET", GoDotEnvVariable("REFRESH_SECRET"))
// 	rtClaims := jwt.MapClaims{}
// 	rtClaims["refresh_uuid"] = td.RefreshUuid
// 	rtClaims["user_id"] = user
// 	rtClaims["exp"] = td.RtExpires
// 	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
// 	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))

// 	if err != nil {
// 		return nil, err
// 	}

// 	return td, nil
// }

// func init() {
// 	dsn := os.Getenv("REDIS_DSN")
// 	if len(dsn) == 0 {
// 		dsn = "localhost:6379"
// 	}

// 	client = redis.NewClient(&redis.Options{
// 		Addr: dsn,
// 	})

// 	_, err := client.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func CreateAuth(user string, td *TokenDetails) error {

// 	at := time.Unix(td.AtExpires, 0)
// 	rt := time.Unix(td.RtExpires, 0)
// 	now := time.Now()

// 	errAccess := client.Set(td.AccessUuid, user, at.Sub(now)).Err()
// 	if errAccess != nil {
// 		return errAccess
// 	}

// 	errRefresh := client.Set(td.RefreshUuid, user, rt.Sub(now)).Err()
// 	if errRefresh != nil {
// 		return errRefresh
// 	}

// 	return nil
// }

// func ExtractToken(r *http.Request) string {
// 	bearToken := r.Header.Get("Authorization")
// 	//normally Authorization the_token_xxx
// 	strArr := strings.Split(bearToken, " ")
// 	if len(strArr) == 2 {
// 		return strArr[1]
// 	}
// 	return ""
// }

// func VerifyToken(r *http.Request) (*jwt.Token, error) {
// 	tokenString := ExtractToken(r)
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("ACCESS_SECRET")), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func TokenValid(r *http.Request) (*jwt.Token, error) {
// 	token, err := VerifyToken(r)

// 	if err != nil {
// 		return nil, err
// 	}

// 	_, ok := token.Claims.(jwt.MapClaims)
// 	if ok && !token.Valid {
// 		return nil, err
// 	}

// 	return token, nil
// }

// func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {

// 	token, err := VerifyToken(r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if ok && token.Valid {
// 		accessUuid, ok := claims["access_uuid"].(string)
// 		if !ok {
// 			return nil, err
// 		}

// 		userId := claims["user_id"].(string)

// 		return &AccessDetails{
// 			AccessUuid: accessUuid,
// 			UserId:     userId,
// 		}, nil
// 	}

// 	return nil, err
// }

// func FetchAuth(authD *AccessDetails) (string, error) {

// 	userID, err := client.Get(authD.AccessUuid).Result()

// 	if err != nil {
// 		return "", err
// 	}

// 	return userID, nil
// }

// func CreateTodo(c *gin.Context) {
// 	var td Todo
// 	if err := c.ShouldBindJSON(&td); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, "invalid json")
// 		return
// 	}

// 	//Extract the access token metadata
// 	metadata, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	userid, err := FetchAuth(metadata)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "aqui la caga")
// 		return
// 	}
// 	td.UserID = userid
// 	td.Uu = metadata.AccessUuid
// 	td.UserFromMeta = metadata.UserId

// 	//you can proceed to save the Todo to a database
// 	//but we will just return it to the caller:

// 	c.JSON(http.StatusCreated, td)
// }

// func DeleteAuth(givenUuid string) (int64, error) {
// 	deleted, err := client.Del(givenUuid).Result()
// 	if err != nil {
// 		return 0, err
// 	}
// 	return deleted, nil
// }

// func Logout(c *gin.Context) {

// 	metadata, err := ExtractTokenMetadata(c.Request)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	if metadata != nil {
// 		delErr := DeleteTokens()
// 		if delErr != nil {
// 			c.JSON(http.StatusUnauthorized, delErr.Error())
// 			return
// 		}
// 		c.JSON(http.StatusOK, "Successfully logged out")
// 	}
// }

// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		t, err := TokenValid(c.Request)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, err.Error())
// 			c.Abort()
// 			return
// 		}

// 		claims, ok := t.Claims.(jwt.MapClaims)
// 		if ok && t.Valid {
// 			uuidValue, ok := claims["access_uuid"].(string)
// 			if !ok {
// 				c.JSON(http.StatusUnprocessableEntity, err)
// 				c.Abort()
// 				return
// 			}

// 			err := ValidateSession(uuidValue)
// 			if err != nil {
// 				c.JSON(http.StatusUnauthorized, "Invalid Session Test")
// 				c.Abort()
// 				return
// 			}
// 		}
// 		c.Next()
// 	}
// }

// func Refresh(c *gin.Context) {

// 	mapToken := map[string]string{}
// 	if err := c.ShouldBindJSON(&mapToken); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}

// 	refreshToken := mapToken["refresh_token"]

// 	//verify the token
// 	os.Setenv("REFRESH_SECRET", GoDotEnvVariable("REFRESH_SECRET")) //this should be in an env file
// 	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
// 		//Make sure that the token method conform to "SigningMethodHMAC"
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("REFRESH_SECRET")), nil
// 	})

// 	//if there is an error, the token must have expired
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Refresh token expired")
// 		return
// 	}

// 	//is token valid?
// 	_, ok := token.Claims.(jwt.MapClaims)
// 	if !ok && !token.Valid {
// 		c.JSON(http.StatusUnauthorized, err)
// 		return
// 	}

// 	//Since token is valid, get the uuid:
// 	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
// 	if ok && token.Valid {
// 		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
// 		if !ok {
// 			c.JSON(http.StatusUnprocessableEntity, err)
// 			return
// 		}
// 		userId := claims["user_id"].(string)

// 		//Delete the previous Refresh Token
// 		deleted, delErr := DeleteAuth(refreshUuid)
// 		if delErr != nil || deleted == 0 { //if any goes wrong
// 			c.JSON(http.StatusUnauthorized, "unauthorized")
// 			return
// 		}

// 		//Create new pairs of refresh and access tokens
// 		ts, createErr := CreateToken(userId)
// 		if createErr != nil {
// 			c.JSON(http.StatusForbidden, createErr.Error())
// 			return
// 		}

// 		//save the tokens metadata to redis
// 		saveErr := CreateAuth(userId, ts)
// 		if saveErr != nil {
// 			c.JSON(http.StatusForbidden, saveErr.Error())
// 			return
// 		}

// 		tokens := map[string]string{
// 			"access_token":  ts.AccessToken,
// 			"refresh_token": ts.RefreshToken,
// 		}
// 		c.JSON(http.StatusCreated, tokens)
// 	} else {
// 		c.JSON(http.StatusUnauthorized, "refresh expired")
// 	}
// }

// func DeleteTokens() error {

// 	//delete access token
// 	deletedAt, err := client.Del(tokenD.AccessUuid).Result()
// 	if err != nil {
// 		return err
// 	}

// 	//delete refresh token
// 	deletedRt, err := client.Del(tokenD.RefreshUuid).Result()
// 	if err != nil {
// 		return err
// 	}

// 	//When the record is deleted, the return value is 1
// 	if deletedAt != 1 || deletedRt != 1 {
// 		return errors.New("something went wrong")
// 	}

// 	return nil
// }

// func ValidateSession(uuid string) error {

// 	_, err := client.Get(uuid).Result()

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
