package letsmock

type NumberMock struct {
}

//  数字MOCK
func (m *NumberMock) MockPositive() int  {
	return mockInt()
}

func (m *NumberMock) MockNegative() int  {
	return mockNegative()
}

func (m *NumberMock) MockIntn(n int) int  {
	return mockIntn(n)
}

func (m *NumberMock) MockIntmn(m1, n int) int  {
	return mockIntmn(m1, n)
}

func (m *NumberMock) MockNumber() int  {
	return (2*mockIntn(2)-1) * mockInt()
}

func (m *NumberMock) MockFloat() float32  {
	return float32(m.MockNumber()) * mockFloat32()
}

func (m *NumberMock) MockFloatmn(m1, n int) float32  {
	if m1 < n {
		return float32(m1) + float32(m.MockIntn(n-m1))*mockFloat32()
	}
	
	return float32(m1)
}

func (m *NumberMock) MockNorm() float64  {
	return mockNormStandard()
}

func (m *NumberMock) MockExp() float64  {
	return mockExp()
}
