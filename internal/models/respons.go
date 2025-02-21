package response

const (
	StatusSuccess = "Success operaton"
	StatusError   = "Error"
)

type Response struct {
	Status string
	Error  string
}

func Success() Response {
	return Response{
		Status: StatusSuccess,
	}
}

func Error(msg string) Response {
	return Response{
		Status: StatusError,
		Error:  msg,
	}
}
