package pushapi

func setDefault(input_map map[string]interface{}, key string, defaultValue interface{}) {
	if _, ok := input_map[key]; !ok {
		input_map[key] = defaultValue
	}
}

func ValidateRequestBody (requestBody map[string]interface{}) (map[string]interface{}, bool) {
      v := new(Validator)
      v.MustContainKey(requestBody, "client_id")
      v.MustContainKey(requestBody, "message")
      v.MustContainKey(requestBody,"topic")
      if v.IsValid() {
      	return requestBody, true
      }
      return nil, false

}
