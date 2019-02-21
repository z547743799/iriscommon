package utils

type E3Result struct {
	Status int64       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func Ok(data interface{}) *E3Result {
	e3 := new(E3Result)
	e3.Status = 200
	e3.Msg = "OK"
	e3.Data = data
	return e3

}

func Build(state int64, msg string) *E3Result {
	e3 := new(E3Result)
	e3.Status = state
	e3.Msg = msg
	e3.Data = nil
	return e3

}
