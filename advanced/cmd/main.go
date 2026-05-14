package main

import (
	"beginner/advanced/app/campaign/controllers"
	"beginner/advanced/pkg/web"
)

func main() {

	app := web.App{}

	controller := controllers.CampaignController{}

	app.AddResource("campaign", controller)

}
