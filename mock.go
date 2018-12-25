package letsmock

type mock struct {
	Number NumberMock
	Date DateMock
	Name NameMock
	Password PasswordMock
	Identity IdentityMock
	Car CarMock
	Inet InternetMock
}

var Mocker mock

func init() {
	Mocker.Date = DateMock{}
	Mocker.Number = NumberMock{}
	Mocker.Name = NameMock{}
	Mocker.Password = PasswordMock{}
	Mocker.Identity = IdentityMock{}
	Mocker.Car = CarMock{}
	Mocker.Inet = InternetMock{}
}

