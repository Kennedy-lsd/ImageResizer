package handlers

import (
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"

	"github.com/Kennedy-lsd/ImageResizer/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/image/draw"
)

func ResizeHandler(c echo.Context) error {
	maxWidth, err := strconv.Atoi(c.QueryParam("width"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid width parameter"})
	}
	maxHeight, err := strconv.Atoi(c.QueryParam("height"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid height parameter"})
	}

	if validationErr := utils.SizeChecker(maxWidth, maxHeight); validationErr != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": validationErr.Error()})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "File upload error"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unable to open uploaded file"})
	}
	defer src.Close()

	ext, err := utils.ValidateImageExtension(src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	src.Seek(0, 0)

	img, _, err := image.Decode(src)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to decode image"})
	}

	originalWidth := img.Bounds().Dx()
	originalHeight := img.Bounds().Dy()
	aspectRatio := float64(originalWidth) / float64(originalHeight)

	var newWidth, newHeight int
	if aspectRatio > 1 {
		newWidth = maxWidth
		newHeight = int(float64(maxWidth) / aspectRatio)
	} else { //Portrait
		newHeight = maxHeight
		newWidth = int(float64(maxHeight) * aspectRatio)
	}

	// blank image with the new size
	resizedImg := image.NewRGBA(image.Rect(0, 0, newWidth, newHeight))

	draw.BiLinear.Scale(resizedImg, resizedImg.Rect, img, img.Bounds(), draw.Over, nil)

	switch ext {
	case ".jpeg", ".jpg":
		c.Response().Header().Set("Content-Type", "image/jpeg")
		if err := jpeg.Encode(c.Response().Writer, resizedImg, nil); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to encode JPEG image"})
		}
	case ".png":
		c.Response().Header().Set("Content-Type", "image/png")
		if err := png.Encode(c.Response().Writer, resizedImg); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to encode PNG image"})
		}
	default:
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Unsupported output file format"})

	}

	return nil
}
