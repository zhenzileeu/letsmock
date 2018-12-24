package letsmock

type mock struct {
	Number NumberMock
	Date DateMock
	Name NameMock
}

var Mocker mock

func init() {
	Mocker.Date = DateMock{}
	Mocker.Number = NumberMock{}
	Mocker.Name = NameMock{}
}

