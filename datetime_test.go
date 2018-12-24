package letsmock

import (
	"fmt"
	"testing"
)

var dateMock = DateMock{}

func TestMockHour(t *testing.T) {
	fmt.Println("mock hour: ", dateMock.MockHour())
}

func TestMockMinutes(t *testing.T) {
	fmt.Println("mock minutes: ", dateMock.MockMinutes())
}

func TestMockYear(t *testing.T)  {
	fmt.Println("mock year: ", dateMock.MockYear())
}

func TestMockDay(t *testing.T)  {
	fmt.Println("mock day: ", dateMock.MockDay())
}

func TestMockWeek(t *testing.T) {
	fmt.Println("mock week: ", dateMock.MockWeek())
}

func TestMockDate(t *testing.T)  {
	fmt.Println("mock date: ", dateMock.MockDate())
}

func TestMockDateTime(t *testing.T)  {
	fmt.Println("mock date time: ", dateMock.MockDateTime())
}

func TestMockHistoryYear(t *testing.T)  {
	fmt.Println("mock history year: ", dateMock.MockHistoryYear())
}

func TestMockHistoryDate(t *testing.T)  {
	fmt.Println("mock history date: ", dateMock.MockHistoryDate())
}

func TestMockHistoryDateTime(t *testing.T)  {
	fmt.Println("mock history date time: ", dateMock.MockHistoryDateTime())
}

func TestMockTime(t *testing.T) {
	fmt.Println("mock time: ", dateMock.MockTime())
}

