package model

type TestModel struct {
	TestA
	TestB
	TestExtendsAnimal
	TestExtendsVehicle
}

type TestExtendsAnimal struct {
	Name string
}

type TestExtendsVehicle struct {
	Price int
}

type TestA interface {
	TestAA() string
}

type TestB interface {
	TestBB() string
}

func (t TestModel) TestAA() string {
	return "TestModel TestAA"
}

func (t TestModel) TestBB() string {
	return "TestModel TestBB"
}

func (t TestModel) UseBothExtends() string {
	t.TestExtendsVehicle.Price = 125
	t.TestExtendsAnimal.Name = "Insane Tiger"
	return "Extends: TestExtendsVehicle and TestExtendsAnimal"
}
