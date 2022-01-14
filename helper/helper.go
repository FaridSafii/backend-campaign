package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	//meta datanya pasti
	//sedangkan data karna fleksibel dijadikan interface kosong
	Meta Meta        `json:"meta"` //Meta saat diparse json dirubah meta
	Data interface{} `json:"data"` //Data saat diparse json dirubah data
}

type Meta struct {
	Message string `json:"message"` //Message saat diparse json dirbah menjadi message
	Code    int    `json:"code"`    //Code saat diparse json dirubah menjadi code
	Status  string `json:"status"`  //Status saat diparse json dirubah menjadi status
}

//Setelah parameter ada nilai balik yaitu Response ->
func APIResponse(message string, code int, status string, data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
