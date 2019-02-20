package engine

type Request struct {
	Url string
	ParserFunc func(string) RequestResult
}

type RequestResult struct {
	R []Request
	Items interface{}
}

func ParserNil(contents string) RequestResult {
	return RequestResult{}
}