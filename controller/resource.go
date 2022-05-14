package controller

import (
	"fmt"
	"log"
	"net/http"
	"time"

	config "restapi-oauth2-go/config"
	model "restapi-oauth2-go/model"

	"github.com/labstack/echo/v4"
)

// @Summary Get data all cake
// @description API get all data cake
// @Tags Cake
// @version 1.0
// @accept application/x-json-stream
// @Success 200 {object} []model.Cake
// @failure 400
// @router /cakes [get]
func (co *Controller) GetCake(c echo.Context) error {
	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	defer db.Close()

	results, err := db.Query("SELECT * FROM cake_table")
	if err != nil {
		panic(err.Error())
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
// @accept application/x-json-stream
// @Param id path string true "id"
// @Success 200 {object} []model.Cake
// @failure 400
// @router /cakes/{id} [get]
func (co *Controller) GetCakebyId(c echo.Context) error {
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
// @accept application/x-json-stream
// @Param cakebody body model.CreateCake true "Cake data"
// @Success 200 {object} []model.CreateCake
// @failure 400
// @router /cakes [post]
func (co *Controller) CreatePostCake(c echo.Context) error {
	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)

	var postbody model.Cake

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	defer db.Close()

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

	return errd
}

// @Summary Delete data cake by id
// @description API delete data cake by id
// @Tags Cake
// @version 1.0
// @accept application/x-json-stream
// @Param id path string true "id"
// @Success 200 {object} []model.MessageData
// @failure 400
// @router /cakes/{id} [delete]
func (co *Controller) DeleteCake(c echo.Context) error {
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

	return errdel
}

// @Summary Patch data cake by id
// @description API update data cake by id
// @Tags Cake
// @version 1.0
// @accept application/x-json-stream
// @Param id path string true "id"
// @Param cakebody body model.UpdateCake true "Cake data"
// @Success 200 {object} []model.UpdateCake
// @failure 400
// @router /cakes/{id} [patch]
func (co *Controller) UpdatePatchCake(c echo.Context) error {
	conn := co.NewConnection()
	db := config.ConnectDB(conn.Server, conn.Port, conn.User, conn.Password, conn.Database)
	defer db.Close()

	id := c.Param("id")

	var postbody model.UpdateCake

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

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

	return errd
}
