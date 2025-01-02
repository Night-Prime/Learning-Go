// Structs is a collection type that contains other different types of variable
// Example: 
type robot struct{
	Model string
	Make string
	Firepower Feature
	Weight int
}

// Nested Struct
type Feature struct {
	Power int
	Radius int
}

// Anonymous struct
myCar := struct {
	Make string
	ModelString
} {
	Make: "tesla",
	Model: "model 3"
}

// Nested-Anonymous Struct
type RobotX struct{
	Model string
	Make string
	Feature struct {
		Power int 
		Radius int
	}
	Weight int
}

// Embedded Structs
type truct struct{
	car
	bedSize int
}

// Working with Methods
type authenticationInfo struct {
	username string
	password string
}

func (aInfo authenticationInfo) getBasicAuth() string  {
	return fmt.Sprintf("Authorization: Basic %s:%s", aInfo.username, aInfo.password,)
}