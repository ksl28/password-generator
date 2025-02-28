package main

import (
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	letterBytes  = "qJzXrMVoBpWYtKaUscFCdEIkHgNlLZShAOyGDmTxbvfPnRJwQu"
	specialBytes = "`|)~!_#?*$]@{+([:<^&\\=}/;,.'%>-\""
	numberBytes  = "4798563021"
)

func genPassword(gc *gin.Context) {
	var (
		includeLetter, includeSpecial, includeNumbers bool
		length                                        int
		err                                           error
	)

	strincludeLetter := gc.Query("includeLetter")
	if strincludeLetter == "" {
		includeLetter = true
	} else {
		includeLetter, err = strconv.ParseBool(strincludeLetter)
		if err != nil {
			log.Printf("Failed to convert the includeLetter - %v \n", err)
			includeLetter = true
		}
	}

	strincludeSpecial := gc.Query("includeSpecial")
	if strincludeSpecial == "" {
		includeSpecial = true
	} else {
		includeSpecial, err = strconv.ParseBool(strincludeSpecial)
		if err != nil {
			log.Printf("Failed to convert the includeSpecial - %v \n", err)
			includeSpecial = true
		}
	}

	strincludeNumbers := gc.Query("includeNumbers")
	if strincludeNumbers == "" {
		includeNumbers = true
	} else {
		includeNumbers, err = strconv.ParseBool(strincludeNumbers)
		if err != nil {
			log.Printf("Failed to convert the includeNumbers - %v", err)
			includeNumbers = true
		}
	}

	strlength := gc.Query("length")
	if strlength == "" {
		length = 16
	} else {
		length, err = strconv.Atoi(strlength)
		if err != nil {
			log.Printf("Failed to convert the length to an integer - %v", err)
			length = 16
		} else if length < 8 {
			length = 8
			log.Printf("The length was set to %v - thats to short", length)
		}
	}

	pwd := make([]byte, length)
	var characterSet string

	if includeLetter {
		characterSet += letterBytes
	}

	if includeSpecial {
		characterSet += specialBytes
	}
	if includeNumbers {
		characterSet += numberBytes
	}

	pwdValid := false
	var clearPwd string
	for !pwdValid {
		for x := range length {
			pwd[x] = characterSet[rand.IntN(len(characterSet))]
		}
		clearPwd = string(pwd)
		pwdValid = validatePassword(clearPwd, includeLetter, includeSpecial, includeNumbers)
	}

	gc.JSON(http.StatusOK, clearPwd)

}

func validatePassword(pwd string, includeLetter bool, includeSpecial bool, includeNumbers bool) bool {
	var letterValid, specialValid, numberValid bool
	for i := range pwd {
		if includeLetter && strings.Contains(letterBytes, string(pwd[i])) {
			letterValid = true
		}
		if includeSpecial && strings.Contains(specialBytes, string(pwd[i])) {
			specialValid = true
		}
		if includeNumbers && strings.Contains(numberBytes, string(pwd[i])) {
			numberValid = true
		}
	}

	if (!letterValid && includeLetter) || (!specialValid && includeSpecial) || (!numberValid && includeNumbers) {
		return false
	} else {
		return true
	}
}

func main() {

	apiPort := os.Getenv("apiport")
	apilistener := fmt.Sprintf("0.0.0.0:%v", apiPort)
	router := gin.Default()
	router.GET("/api/v1/genpwd", genPassword)
	router.Run(apilistener)
}
