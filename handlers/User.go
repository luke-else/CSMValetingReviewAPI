package handlers

import (
	"CSMValetingReviewAPI/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func Upload(c echo.Context) error {

	//Personal details -- Potentially nil
	firstName := c.FormValue("FirstName")
	lastName := c.FormValue("LastName")
	contactNumber := c.FormValue("ContactNumber")

	//Quality Statistics
	service := c.FormValue("Service")
	quality, err := strconv.ParseInt(c.FormValue("Quality"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
	}

	recommendation, err := strconv.ParseInt(c.FormValue("Recommendation"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Error{Error: err.Error()})
	}

	notes := c.FormValue("Notes")
	//images, _ := c.FormFile("image")

	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, "Failed")
	// }

	Review := models.UserReview{
		FirstName:     firstName,
		LastName:      lastName,
		ContactNumber: contactNumber,

		Service:        service,
		Quality:        quality,
		Recommendation: recommendation,
		Notes:          notes,
		//Image:          images,
	}

	fmt.Println(Review)

	//Upload into MongoDB

	return c.JSON(http.StatusOK, models.Valid{Status: "Valid"})
}
