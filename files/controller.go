package files

import "github.com/gofiber/fiber/v2"

func upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}
	f, err := file.Open()
	if err != nil {
		return err
	}
	defer f.Close()

	buffer := make([]byte, file.Size)
	if _, err := f.Read(buffer); err != nil {
		return err
	}
	return nil
}
