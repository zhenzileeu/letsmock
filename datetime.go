package letsmock

import "time"

type DateMock struct {
}

func (m *DateMock) MockHour() int  {
	return mockIntn(23)
}

func (m *DateMock) MockMinutes() int  {
	return mockIntn(59)
}

func (m *DateMock) MockSeconds() int  {
	return mockIntn(59)
}

func (m *DateMock) MockWeek() int  {
	return mockIntmn(1, 7)
}

func (m *DateMock) MockMonth() int  {
	return mockIntmn(1, 12)
}

func (m *DateMock) MockDay() int  {
	return mockIntmn(1, 31)
}

func (m *DateMock) MockYear() int  {
	return 1970+mockIntn(100)
}

func (m *DateMock) MockHistoryYear() int  {
	return 1970+mockIntn(time.Now().Year()-1970)
}

func (m *DateMock) MockDate() string  {
	return time.Date(
		m.MockYear(),
		time.Month(m.MockMonth()),
		m.MockDay(), 0, 0, 0, 0,
		time.Local).Format("2006-01-02")
}

func (m *DateMock) MockTime() string {
	return time.Date(m.MockYear(), time.Month(m.MockMonth()), m.MockDay(), m.MockHour(), m.MockMinutes(), m.MockSeconds(), 0,
			time.Now().Location()).Format("15:04:05")
}

func (m *DateMock) MockDateTime() string  {
	return time.Date(m.MockYear(), time.Month(m.MockMonth()), m.MockDay(), m.MockHour(), m.MockMinutes(), m.MockSeconds(), 0,
		time.Now().Location()).Format("2006-01-02 15:04:05")
}

func (m *DateMock) MockHistoryDate() string  {
	return time.Date(
		m.MockHistoryYear(),
		time.Month(m.MockMonth()),
		m.MockDay(), 0, 0, 0, 0,
		time.Now().Location()).Format("2006-01-02")
}

func (m *DateMock) MockHistoryDateTime() string  {
	return time.Date(m.MockHistoryYear(), time.Month(m.MockMonth()), m.MockDay(), m.MockHour(), m.MockMinutes(), m.MockSeconds(), 0,
		time.Now().Location()).Format("2006-01-02 15:04:05")
}
