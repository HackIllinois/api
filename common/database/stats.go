package database

import (
	"errors"
	"fmt"
	"reflect"
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
func AddEntryToStats(stats map[string]interface{}, entry map[string]interface{}) error {
	for key, v := range entry {
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

			AddEntryToStats(mapped_stats, value)
		default:
			reflect_type := reflect.TypeOf(value)
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
