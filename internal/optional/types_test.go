package optional

import (
	"encoding/json"
	"testing"

	"github.com/eduardolat/openroutergo/internal/assert"
)

func TestOptionalGenericType(t *testing.T) {
	t.Run("MarshalJSON", func(t *testing.T) {
		// Test unset value
		unset := &Optional[string]{IsSet: false}
		data, err := unset.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(data))

		// Test set value
		set := &Optional[string]{IsSet: true, Value: "hello"}
		data, err = set.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, `"hello"`, string(data))
	})

	t.Run("UnmarshalJSON", func(t *testing.T) {
		// Test null value
		var opt Optional[int]
		err := opt.UnmarshalJSON([]byte("null"))
		assert.NoError(t, err)
		assert.False(t, opt.IsSet)

		// Test valid value
		err = opt.UnmarshalJSON([]byte("42"))
		assert.NoError(t, err)
		assert.True(t, opt.IsSet)
		assert.Equal(t, 42, opt.Value)
	})
}

func TestDerivedTypes(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		// Marshal then unmarshal to verify full cycle
		original := String{IsSet: true, Value: "test"}
		data, err := json.Marshal(&original)
		assert.NoError(t, err)
		var result String
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)
		assert.True(t, result.IsSet)
		assert.Equal(t, "test", result.Value)
	})

	t.Run("Int", func(t *testing.T) {
		// Test int with null value
		var num Int
		err := json.Unmarshal([]byte("null"), &num)
		assert.NoError(t, err)
		assert.False(t, num.IsSet)

		// Test with value
		err = json.Unmarshal([]byte("99"), &num)
		assert.NoError(t, err)
		assert.True(t, num.IsSet)
		assert.Equal(t, 99, num.Value)
	})

	t.Run("Bool", func(t *testing.T) {
		// Testing marshal/unmarshal false value
		original := Bool{IsSet: true, Value: false}
		data, err := json.Marshal(original)
		assert.NoError(t, err)

		var result Bool
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)
		assert.True(t, result.IsSet)
		assert.False(t, result.Value)
	})
}

func TestEmptyStringBehavior(t *testing.T) {
	// This test is critical because empty string isn't valid JSON
	t.Run("Direct vs json.Marshal", func(t *testing.T) {
		opt := &Optional[string]{IsSet: false}

		// Direct call returns empty string
		direct, err := opt.MarshalJSON()
		assert.NoError(t, err)
		assert.Equal(t, "null", string(direct))

		// json.Marshal might handle it differently
		_, err = json.Marshal(opt)
		assert.NoError(t, err)
	})
}

func TestRealWorldUsage(t *testing.T) {
	// Test with a realistic struct
	type User struct {
		Name     string `json:"name"`
		Age      Int    `json:"age"`
		Email    String `json:"email,omitempty"`
		Verified Bool   `json:"verified,omitempty"`
	}

	t.Run("Complete serialization cycle", func(t *testing.T) {
		original := User{
			Name:     "Jane",
			Age:      Int{IsSet: true, Value: 30},
			Email:    String{IsSet: true, Value: "jane@example.com"},
			Verified: Bool{IsSet: false},
		}

		data, err := json.Marshal(original)
		assert.NoError(t, err)

		var result User
		err = json.Unmarshal(data, &result)
		assert.NoError(t, err)

		assert.Equal(t, original.Name, result.Name)
		assert.Equal(t, original.Age.Value, result.Age.Value)
		assert.True(t, result.Age.IsSet)
		assert.Equal(t, original.Email.Value, result.Email.Value)
		assert.True(t, result.Email.IsSet)
		assert.False(t, result.Verified.IsSet)
	})
}
