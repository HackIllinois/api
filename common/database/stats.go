package database

import (
	"errors"
	"fmt"
	"github.com/HackIllinois/api/common/utils"
	"reflect"
	"strings"
)

var ErrTypeMismatch = errors.New("Error: TYPE_MISMATCH")

/*
	Returns a maps of default stats
*/
func GetDefaultStats() map[string]interface{} {
	stats := make(map[string]interface{})
	return stats
}

/*
	Updates the stats with the given entry
*/
func AddEntryToStats(stats map[string]interface{}, entry map[string]interface{}, fields []string) error {
	for key, v := range entry {
		top_level_fields := ExtractTopLevel(fields)
		if !utils.ContainsString(top_level_fields, key) {
			continue
		}

		switch value := v.(type) {
		case map[string]interface{}:
			_, exists := stats[key]

			if !exists {
				stats[key] = make(map[string]interface{})
			}

			mapped_stats, ok := stats[key].(map[string]interface{})

			if !ok {
				return ErrTypeMismatch
			}

			stripped_fields := RemoveTopLevel(fields)

			AddEntryToStats(mapped_stats, value, stripped_fields)
		default:
			reflect_type := reflect.TypeOf(value)
			if reflect_type == nil {
				return nil
			}
			switch reflect_type.Kind() {
			case reflect.Array:
				fallthrough
			case reflect.Slice:
				slice_value := reflect.ValueOf(value)
				for i := 0; i < slice_value.Len(); i++ {
					element := slice_value.Index(i)
					err := UpdateStatsField(stats, key, element)

					if err != nil {
						return err
					}
				}
			default:
				err := UpdateStatsField(stats, key, value)

				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func UpdateStatsField(stats map[string]interface{}, key string, value interface{}) error {
	_, exists := stats[key]

	if !exists {
		stats[key] = make(map[string]int)
	}

	mapped_stats, ok := stats[key].(map[string]int)

	if !ok {
		return ErrTypeMismatch
	}

	value_key := fmt.Sprintf("%v", value)
	mapped_stats[value_key] = mapped_stats[value_key] + 1

	return nil
}

/*
	Remove everything in each field after, and including, the first '.'
*/
func ExtractTopLevel(fields []string) []string {
	top_level_fields := []string{}

	for _, field := range fields {
		split_field := strings.SplitN(field, ".", 2)
		if len(split_field) > 0 {
			top_level_fields = append(top_level_fields, split_field[0])
		}
	}

	return top_level_fields
}

/*
	Remove everything in each field before, and including, the first '.'
*/
func RemoveTopLevel(fields []string) []string {
	stripped_fields := []string{}

	for _, field := range fields {
		split_field := strings.SplitN(field, ".", 2)
		if len(split_field) > 1 {
			stripped_fields = append(stripped_fields, split_field[1])
		}
	}

	return stripped_fields
}
