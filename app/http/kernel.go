package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/memory"
)

var memoryStore = memory.New()
var store = session.New(session.Config{
	Storage: memoryStore,
})

func Kernel(app *fiber.App) {
	app.Use(recover.New())
}

func GetSession(c *fiber.Ctx) *session.Session {
	sess, err := store.Get(c)
	if err != nil {
		panic(err)
	}
	return sess
}

func GetStore() *session.Store {
	return store
}
