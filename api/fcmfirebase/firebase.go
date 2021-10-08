package fcmfirebase

import (
	"log"
	"net/http"

	"github.com/appleboy/go-fcm"
	"github.com/labstack/echo/v4"
)

func GetNotifikasiFCM(c echo.Context)error{
// Create the message to be sent.
	msg := &fcm.Message{
		To: "dGLdDK73fam5tis2polq_Y:APA91bFp9PiTlpW3y2-oZmLVONpmbWYJSsnKcleAvOZRTD8x4clfKV8i_IAA77qjuQ_YwCBUjrcRZkQD6Lfh8SBirrZ1AqRiuIYKh5B5SUaZPu9sEjYUyY3crgrZ4iuB_fG0JfnkzLpw",
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "title",
			Body: "body",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient("	AAAAsGQmN-I:APA91bH9FEiIqJ5EsEbeGOkJ6qeChm4mrgbxaIqVJL6fvtF6sXgEJTVWyXRbVnu5fitwPPAUXLPeYz8JxT1LxU1BuSxsqhTeLqiGbq6fGA1em2AGGohvLBKdviizOEZ3-AE2E44CfnvK")
	if err != nil {
		log.Fatalln(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%#v\n", response)
	//send json response
	return c.JSON(http.StatusOK, response)
}