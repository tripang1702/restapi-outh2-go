package controller

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	config "restapi-oauth2-go/config"
	model "restapi-oauth2-go/model"

	echoserver "github.com/dasjott/oauth2-echo-server"
	"github.com/labstack/echo"
)

// @Summary Get Access Token
// @description API to generate access token
// @Tags Authorization
// @version 1.0
// @accept application/json
// @Param grant_type query string true "Grant type (example : password, refresh_token)"
// @Param client_id query string true "Client id"
// @Param client_secret query string true "Client secret"
// @Param username query string false "used when grant_type=password"
// @Param password query string false "used when grant_type=password"
// @Param scope query string true "scope (example : read)"
// @Param refresh_token query string false "used when grant_type=refresh_token"
// @Success 200 {object} model.Tokensuccess
// @failure 500 {object} model.ErrorToken
// @router /oauth2/token [get]
func (co *Controller) HandleTokenRequest(c echo.Context) error {
	return echoserver.HandleTokenRequest(c)
}

// @Summary Get data all cake
// @description API get all data cake
// @Tags Cake
// @version 1.0
// @accept application/json
// @Param Authorization header string true "Bearer Authorization"
// @Success 200 {object} []model.Cake
// @failure 400
// @router /api/cakes [get]
func (co *Controller) GetCake(c echo.Context) error {
	userID := CheckuserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorize",
		})

	}
	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	defer db.Close()

	results, err := db.Query("SELECT * FROM cake_table")
	if err != nil {
		log.Print(err.Error())
	}
	defer results.Close()

	var final_result []model.Cake
	var cake model.Cake
	for results.Next() {

		err = results.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.Created_at, &cake.Updated_at)
		if err != nil {
			panic(err.Error())
		} else {
			final_result = append(final_result, model.Cake{Id: cake.Id, Title: cake.Title, Description: cake.Description, Rating: cake.Rating, Image: cake.Image, Created_at: cake.Created_at, Updated_at: cake.Updated_at})
		}
	}

	return c.JSON(http.StatusOK, final_result)
}

// @Summary Get data cake by id
// @description API get data cake by id
// @Tags Cake
// @version 1.0
// @accept application/json
// @Param Authorization header string true "Bearer Authorization"
// @Param id path string true "id"
// @Success 200 {object} []model.Cake
// @failure 400
// @router /api/cakes/{id} [get]
func (co *Controller) GetCakebyId(c echo.Context) error {
	userID := CheckuserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorize",
		})

	}
	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	id := c.Param("id")

	defer db.Close()

	results, err := db.Query("SELECT * FROM cake_table where id = ?", id)
	if err != nil {
		log.Print(err.Error())
	}
	defer results.Close()

	var final_result []model.Cake
	var cake model.Cake
	for results.Next() {

		err = results.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.Created_at, &cake.Updated_at)
		if err != nil {
			log.Print(err.Error())
		} else {
			final_result = append(final_result, model.Cake{Id: cake.Id, Title: cake.Title, Description: cake.Description, Rating: cake.Rating, Image: cake.Image, Created_at: cake.Created_at, Updated_at: cake.Updated_at})
		}
	}
	if len(final_result) == 0 {
		var null_result = model.MessageData{
			Status:  false,
			Message: fmt.Sprintf("No data for cakes id = %s", id),
		}
		return c.JSON(http.StatusOK, null_result)

	}
	return c.JSON(http.StatusOK, final_result)
}

// @Summary Post data cake by id
// @description API create data cake
// @Tags Cake
// @version 1.0
// @accept application/json
// @Param Authorization header string true "Bearer Authorization"
// @Param cakebody body model.CreateCake true "Cake data"
// @Success 200 {object} []model.CreateCake
// @failure 400
// @router /api/cakes [post]
func (co *Controller) CreatePostCake(c echo.Context) error {
	userID := CheckuserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorize",
		})

	} else {
		if userID != "1" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User token cannot access this resource",
			})
		}
	}

	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)
	defer db.Close()

	var postbody model.Cake

	loc, _ := time.LoadLocation("Asia/Bangkok")
	t := time.Now().In(loc)

	dt := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	errbind := c.Bind(&postbody)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": false, "Message error": errbind.Error()})
	}

	insert, errd := db.Exec("INSERT INTO cake_table (title,description,rating,image,created_at) values (?,?,?,?,?)", postbody.Title, postbody.Description, postbody.Rating, postbody.Image, dt)

	if insert != nil {
		user_id, errx := insert.LastInsertId()

		if errx == nil {
			var messageResponse = model.MessageData{
				Status:  true,
				Message: fmt.Sprintf("Add new data successfully with cakes id = %d", user_id),
			}
			return c.JSON(http.StatusCreated, messageResponse)

		} else {
			var messageResponse = model.MessageData{
				Status:  false,
				Message: errx.Error(),
			}
			return c.JSON(http.StatusNotImplemented, messageResponse)
		}
	}

	return c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"Status":  false,
		"Message": errd.Error(),
	})
}

// @Summary Delete data cake by id
// @description API delete data cake by id
// @Tags Cake
// @version 1.0
// @accept application/json
// @Param Authorization header string true "Bearer Authorization"
// @Param id path string true "id"
// @Success 200 {object} []model.MessageData
// @failure 400
// @router /api/cakes/{id} [delete]
func (co *Controller) DeleteCake(c echo.Context) error {
	userID := CheckuserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorize",
		})

	} else {
		if userID != "1" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User token cannot access this resource",
			})
		}
	}

	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	id := c.Param("id")

	defer db.Close()

	delete, errdel := db.Exec("DELETE FROM cake_table where id = ?", id)

	if errdel != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": false, "Message error": errdel.Error()})
	} else {
		count, err2 := delete.RowsAffected()

		if err2 != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"status": false, "Message error": err2.Error()})
			log.Print(err2.Error())
		} else {
			if count != 0 {
				var messageResponse = model.MessageData{
					Status:  true,
					Message: fmt.Sprintf("delete data successfully with affected row = %d", count),
				}
				return c.JSON(http.StatusOK, messageResponse)

			} else {
				var messageResponse = model.MessageData{
					Status:  false,
					Message: fmt.Sprintf("Delete Fail, cake id %s not found", id),
				}
				return c.JSON(http.StatusNotImplemented, messageResponse)

			}
		}
	}

	return c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"Status":  false,
		"Message": errdel.Error(),
	})
}

// @Summary Patch data cake by id
// @description API update data cake by id
// @Tags Cake
// @version 1.0
// @accept application/json
// @Param Authorization header string true "Bearer Authorization"
// @Param id path string true "id"
// @Param cakebody body model.UpdateCake true "Cake data"
// @Success 200 {object} []model.UpdateCake
// @failure 400
// @router /api/cakes/{id} [patch]
func (co *Controller) UpdatePatchCake(c echo.Context) error {
	userID := CheckuserID(c)
	if userID == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "Unauthorize",
		})

	} else {
		if userID != "1" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "User token cannot access this resource",
			})
		}
	}

	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)
	defer db.Close()

	id := c.Param("id")

	var postbody model.UpdateCake

	loc, _ := time.LoadLocation("Asia/Bangkok")
	t := time.Now().In(loc)

	dt := fmt.Sprintf("%d-%d-%d %d:%d:%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())

	errbind := c.Bind(&postbody)
	if errbind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"status": false, "Message error": errbind.Error()})
	}

	update, errd := db.Exec("UPDATE cake_table SET title=?,description=?,rating=?,image=?,updated_at=? WHERE id =?", postbody.Title, postbody.Description, postbody.Rating, postbody.Image, dt, id)

	if update != nil {
		count, errx := update.RowsAffected()

		if errx != nil {
			c.JSON(http.StatusBadRequest, map[string]interface{}{"status": false, "Message error": errx.Error()})
			log.Print(errx.Error())
		} else {
			if count != 0 {
				var messageResponse = model.MessageData{
					Status:  true,
					Message: fmt.Sprintf("update data successfully with affected row = %d", count),
				}
				return c.JSON(http.StatusOK, messageResponse)

			} else {
				var messageResponse = model.MessageData{
					Status:  false,
					Message: fmt.Sprintf("Update Fail, cake id %s not found", id),
				}
				return c.JSON(http.StatusNotImplemented, messageResponse)
			}
		}
	}

	return c.JSON(http.StatusNotImplemented, map[string]interface{}{
		"Status":  false,
		"Message": errd.Error(),
	})
}

func CheckuserID(c echo.Context) string {
	ti := c.Get(echoserver.DefaultConfig.TokenKey)
	if ti != "" {
		r := reflect.ValueOf(ti)
		f := reflect.Indirect(r).FieldByName("UserID")
		return f.Interface().(string)
	}
	return ""
}
