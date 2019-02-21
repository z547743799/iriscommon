package utils

// developers can use any library to add a custom cookie encoder/decoder.
// At this example we use the gorilla's securecookie package:
// $ go get github.com/gorilla/securecookie
// $ go run main.go

import (
	"gitlab.com/z547743799/iriscontent/redisinit"
	"time"

	"github.com/kataras/iris/sessions"

	"github.com/gorilla/securecookie"
)

var Manager *sessions.Sessions


func init() {

	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte("the-big-and-secret-fash-key-here")
	blockKey := []byte("lot-secret-of-characters-big-too")
	secureCookie := securecookie.New(hashKey, blockKey)

	Manager = sessions.New(sessions.Config{
		Cookie:       "token",
		Encode:       secureCookie.Encode,
		Decode:       secureCookie.Decode,
		Expires:      45 * time.Minute, // <=0 means unlimited life. Defaults to 0.
		AllowReclaim: true,
	})
	Manager.UseDatabase(redisinit.IRRe)
}
