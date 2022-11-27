package resource

type Employee struct {
	Id    int
	Email string
	Address
}

type Address struct {
	HouseNo int
	Society string
	Pin     int
}

type employeeBuilder struct {
	employee *Employee
}

type employeeAddressBuilder struct {
	employeeBuilder
}

type employeeOfficialBuilder struct {
	employeeBuilder
}

func NewEmployeeBuilder() *employeeBuilder {
	return &employeeBuilder{employee: &Employee{}}
}

func (eb *employeeBuilder) PersonalInfo() *employeeAddressBuilder {
	return &employeeAddressBuilder{*eb}
}

func (eb *employeeAddressBuilder) AtHouseNo(house int) *employeeAddressBuilder {
	eb.employee.HouseNo = house
	return eb
}
func (eb *employeeAddressBuilder) InSociety(society string) *employeeAddressBuilder {
	eb.employee.Society = society
	return eb
}
func (eb *employeeAddressBuilder) HavingPin(pin int) *employeeAddressBuilder {
	eb.employee.Pin = pin
	return eb
}

func (eb *employeeBuilder) WorkInfo() *employeeOfficialBuilder {
	return &employeeOfficialBuilder{*eb}
}
func (eb *employeeOfficialBuilder) WithID(id int) *employeeOfficialBuilder {
	eb.employee.Id = id
	return eb
}
func (eb *employeeOfficialBuilder) WithEmail(email string) *employeeOfficialBuilder {
	eb.employee.Email = email
	return eb
}

func (eb *employeeBuilder) Build() *Employee {
	return eb.employee
}
