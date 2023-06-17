package main

import "DevelopHub/AuthServer/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {

}
