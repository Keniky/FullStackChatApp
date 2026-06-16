package services

import (
	"chatApp/repository"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func SignInNewUser(name string, file *multipart.File) (res bool) {

	newFile, err := os.CreateTemp("./static/app/", "image*.png")
	if err != nil {
		fmt.Println("failed to create the image")
		return false
	}
	defer newFile.Close()

	//copy the image inside this one close file return true
	io.Copy(newFile, *file)

	//send username and file name
	//get file name and add to it /public/
	// repository.SaveUserInDB(name, "/public/"+filepath.Base(newFile.Name()))
	repository.SaveUserInDB(name, "/static/app/"+filepath.Base(newFile.Name()))

	res = true
	return res
}

func VerifyIfUserIsSignedUp(name string) *repository.UserNameAndPfp {
	res := repository.VerifyIfUserIsSignedUp(name)

	return res
}
