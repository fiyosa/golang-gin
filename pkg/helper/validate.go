package helper

import (
	"encoding/json"
	"go-gin/lang"
	"go-gin/pkg/secret"
	"go-gin/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate(c *gin.Context, input interface{}) bool {
	if err := c.ShouldBindJSON(input); err != nil {
		newErrors := map[string][]string{}
		msg := ""

		switch v := err.(type) {
		case *json.UnmarshalTypeError:
			newMsg := "Json binding error: " + v.Field + " type error"
			newErrors[v.Field] = append(newErrors[v.Field], newMsg)
			msg = newMsg

		case validator.ValidationErrors:
			for _, e := range v {
				newMsg := e.Translate(validation.Translator)
				newErrors[e.Field()] = append(newErrors[e.Field()], newMsg)
				if msg == "" {
					msg = newMsg
				}
			}

		default:
			if secret.APP_ENV == "development" {
				msg = v.Error()
			} else {
				msg = lang.L(lang.SetL().SOMETHING_WENT_WRONG, nil)
			}

		}

		SendErrors(c, msg, newErrors)

		return true
	}

	return false
}
