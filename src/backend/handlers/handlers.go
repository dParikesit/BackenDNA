package handlers

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kristabdi/Tubes3_13520058/controllers"
	"github.com/kristabdi/Tubes3_13520058/models"
	"github.com/kristabdi/Tubes3_13520058/utils"
	"gorm.io/gorm"
)

func DiseaseInsert(c *fiber.Ctx) error {
	var err error
	var query = new(models.Disease)
	if err = c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Parsing error")
	}

	if !utils.IsValidDiseaseSearchInput(query.Name) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Name")
	}

	valid := utils.IsValidDNA(query.Sequence)
	if !valid {
		return c.Status(fiber.StatusBadRequest).SendString("Sequence invalid")
	}

	if _, err = controllers.DiseaseGetOne(query.Name); errors.Is(err, gorm.ErrRecordNotFound) {
		err = controllers.DiseaseInsertOne(query)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Disease Insert Error")
		}

		return c.Status(fiber.StatusOK).SendString("Disease Data Added")
	} else {
		return c.Status(fiber.StatusBadRequest).SendString("Disease already in DB")
	}
}

func DiseaseMatch(c *fiber.Ctx) error {
	var err error
	query := new(models.QueryMatch)
	log.Println(query)
	if err = c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Parsing error")
	}

	valid := utils.IsValidDNA(query.Sequence)
	if !valid {
		return c.Status(fiber.StatusBadRequest).SendString("Sequence invalid")
	}

	if !utils.IsValidDiseaseSearchInput(query.Penyakit) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Penyakit Name")
	}

	var disease models.Disease
	if disease, err = controllers.DiseaseGetOne(query.Penyakit); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(fiber.StatusNotFound).SendString("Disease not found")
		} else {
			return c.Status(fiber.StatusInternalServerError).SendString("Disease Get Error")
		}
	}

	var similar bool
	if c.Params("algo") == "bm" {
		similar = utils.BMMatch(query.Sequence, disease.Sequence)
	} else if c.Params("algo") == "kmp" {
		similar = utils.KMPMatch(query.Sequence, disease.Sequence)
	}

	var similarity float32
	if !similar {
		// String similarity antara query.Sequence dengan disease.Sequence
		similarity = utils.CalculateSimiliarity(query.Sequence, disease.Sequence)
		if similarity >= 0.8 {
			similar = true
		} else {
			similar = false
		}
	} else {
		similarity = 1.0
	}

	history := models.History{
		Name:       query.Name,
		Penyakit:   disease.Name,
		Similarity: similarity,
		IsTrue:     similar,
		CreatedAt:  time.Now(),
	}

	if err = controllers.HistoryInsertOne(&history); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("History Insert Error")
	}

	var dateString string
	if dateString, err = utils.ConvertTime(history.CreatedAt); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Date string Error")
	}

	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("%s %s %s %.2f %t", dateString, history.Name, history.Penyakit, history.Similarity, history.IsTrue))
}

func HistoryQuery(c *fiber.Ctx) error {
	var err error
	var history []models.History

	query := new(models.QueryHistory)
	if err = c.BodyParser(query); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Parsing error")
	}

	// get input text
	var text string = query.Text
	var arr []string = utils.SplitText(text, " ")
	var date string
	var name string

	if !utils.IsValidInputSearch(text) {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid Search Query")
	} else {
		if len(arr) >= 3 {
			date = arr[0] + " " + arr[1] + " " + arr[2]
			if utils.IsValidDate(date) {
				name = utils.JoinArray(arr, 3)
				if !utils.IsValidDiseaseSearchInput(name) {
					return c.Status(fiber.StatusBadRequest).SendString("Invalid Name")
				}
			} else {
				// date = ""
				// name = utils.JoinArray(arr, 0)
				return c.Status(fiber.StatusBadRequest).SendString("Invalid Date")
			}
		} else {
			name = utils.JoinArray(arr, 0)
			if !utils.IsValidDiseaseSearchInput(name) {
				return c.Status(fiber.StatusBadRequest).SendString("Invalid Name")
			}
		}
	}

	if date == "" {
		if history, err = controllers.HistoryGetByName(name); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).SendString("History not found")
			} else {
				return c.Status(fiber.StatusInternalServerError).SendString("History Get Error")
			}
		}
	} else if name == "" {
		if history, err = controllers.HistoryGetByDate(date); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).SendString("History not found")
			} else if err.Error() == "Invalid Date" {
				return c.Status(fiber.StatusInternalServerError).SendString("Invalid Date")
			} else {
				return c.Status(fiber.StatusInternalServerError).SendString("History Get Error")
			}
		}
	} else {
		if history, err = controllers.HistoryGetByAll(name, date); err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.Status(fiber.StatusNotFound).SendString("History not found")
			} else if err.Error() == "Invalid Date" {
				return c.Status(fiber.StatusInternalServerError).SendString("Invalid Date")
			} else {
				return c.Status(fiber.StatusInternalServerError).SendString("History Get Error")
			}
		}
	}

	if len(history) <= 0 {
		return c.Status(fiber.StatusNotFound).SendString("History not found")
	}

	var output []string
	var dateString string

	for _, hist := range history {
		if dateString, err = utils.ConvertTime(hist.CreatedAt); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Date string Error")
		}

		output = append(output, fmt.Sprintf("%s %s %s %.2f %t", dateString, hist.Name, hist.Penyakit, hist.Similarity, hist.IsTrue))
	}

	return c.Status(fiber.StatusOK).JSON(output)
}
