package helpers

func HandleError(data interface{}, err interface{}) (interface{}, interface{}) {
	if err != nil {
		return data, err
	}
	return data, nil
}
