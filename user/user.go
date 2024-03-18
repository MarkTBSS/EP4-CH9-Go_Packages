package user

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Step 3.2 : Err struct
type Err struct {
	Message string `json:"message"`
}
