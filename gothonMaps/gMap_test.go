package gothonMaps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*---------------------------------------------------------------------------------------------------------
------------------------------------------SETDEFAULT METHOD TESTS------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestSetDefault(t *testing.T) {
	t.Run("SetDefault string to map[string]string when key does not exist", func(t *testing.T) {

		m := map[string]string{
			"brand": "Ford",
			"year":  "1964",
		}
		expected := map[string]string{
			"brand": "Ford",
			"model": "Bronco",
			"year":  "1964",
		}
		value := SetDefault(m, "model", "Bronco")
		assert.Equal(t, "Bronco", value)
		assert.Equal(t, m, expected)
	})

	t.Run("SetDefault string to map[string]string when key does exist", func(t *testing.T) {

		m := map[string]string{
			"brand": "Ford",
			"model": "Mustang",
			"year":  "1964",
		}
		expected := map[string]string{
			"brand": "Ford",
			"model": "Mustang",
			"year":  "1964",
		}
		value := SetDefault(m, "model", "Bronco")
		assert.Equal(t, "Mustang", value)
		assert.Equal(t, m, expected)
	})
}

/*---------------------------------------------------------------------------------------------------------
------------------------------------------ITEMS METHOD TESTS-----------------------------------------------
-----------------------------------------------------------------------------------------------------------
*/

func TestItems(t *testing.T) {
	t.Run("SetDefault string to map[string]string when key does exist", func(t *testing.T) {

		m := map[string]string{
			"brand": "Ford",
			"model": "Mustang",
			"year":  "1964",
		}
		expected := []Pair[string, string]{
			{"brand", "Ford"},
			{"model", "Mustang"},
			{"year", "1964"},
		}
		items := Items(m)
		assert.Equal(t, expected, items)
	})
}
