package services

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type FilterOptions struct {
	DefaultLimit int
	AllowSortBy  []string          // e.g., []string{"createdAt", "name"}
	ExtraFilters map[string]string // key: query param, value: db field
}

type FilterResult struct {
	Pagination struct {
		Page  int
		Limit int
		Skip  int
	}
	Filter   bson.M
	FindOpts *options.FindOptions
}

func FilterBuilder(c *fiber.Ctx, otps FilterOptions) (FilterResult, error) {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit

	filter := bson.M{}
	nameFilter := strings.TrimSpace(c.Query("name"))
	if nameFilter != "" {
		filter["name"] = bson.M{"$regex": nameFilter, "$options": "i"}
	}

	for queryKey, dbField := range otps.ExtraFilters {
		val := strings.TrimSpace(c.Query(queryKey))
		if val != "" {
			filter[dbField] = bson.M{"$regex": val, "$options": "i"}
		}
	}

	// Sort logic
	sortField := c.Query("sortBy", "createdAt")
	sortOrderStr := strings.ToLower(c.Query("sortOrder", "desc")) // default is "desc"
	sortOrder := -1
	switch sortOrderStr {
	case "asc":
		sortOrder = 1
	case "desc":
		sortOrder = -1
	default:
		sortOrder = -1 // fallback if value is invalid
	}

	sort := bson.D{}
	for _, allowed := range otps.AllowSortBy {
		if sortField == allowed {
			sort = append(sort, bson.E{Key: sortField, Value: sortOrder})
			break
		}
	}

	// Find options
	findOpts := options.Find()
	findOpts.SetLimit(int64(limit))
	findOpts.SetSkip(int64(skip))
	if len(sort) > 0 {
		findOpts.SetSort(sort)
	}

	return FilterResult{
		Pagination: struct {
			Page  int
			Limit int
			Skip  int
		}{page, limit, skip},
		Filter:   filter,
		FindOpts: findOpts,
	}, nil
}
