package webFunctions

//Create Structs
// type newUser struct {
// 	Name        string  `json:"name" binding:"required"`
// 	Price       float64 `json"price" binding:"required,gt=0"`
// 	Description string  `json:"description" binding:"omitempty,max=250"`
// }

// type User struct {
// 	GUID        string  `json:"guid"`
// 	Name        string  `json:"name"`
// 	Price       float64 `json"price"`
// 	Description string  `json:"description"`
// 	CreatedAt   string  `json:"createdAt"`
// }

type UpdatedUser struct {
	Name          string `json:"name"`          //binding:"required_without_all=Price Description"`
	Read_Setting  bool   `json:"read_setting"`  //binding:"omitempty,gt=0"`
	Write_Setting bool   `json:"write_setting"` //binding:"omitempty,max=250`
}

// type newNote struct {
// 	Name        string  `json:"name"` //binding:"required"`
// 	Text            string `json:"text"` //binding:"omitempty,max=250"`
// 	Completion_Time string `json:"completion_date"` //binding:"omitempty,max=250"`
// 	Status          string `json:"status"` //binding:"required,gt=0"`
// 	Delegation      int    `json:"delegation"` //binding:"required"`
// 	Shared_Users    string `json:"shared_users"` //binding:"required"`
// }

// type User struct {
// 	GUID        string  `json:"guid"`
// 	Name        string  `json:"name"`
// 	Price       float64 `json"price"`
// 	Description string  `json:"description"`
// 	CreatedAt   string  `json:"createdAt"`
// }

type UpdatedNote struct {
	Name         string `json:"name"`         //binding:"required"`
	Text         string `json:"text"`         //binding:"omitempty,max=250"`
	Status       string `json:"status"`       //binding:"required,gt=0"`
	Delegation   int    `json:"delegation"`   //binding:"required"`
	Shared_Users string `json:"shared_users"` //binding:"required"`
}

type guidBinding struct {
	GUID string `url:"guid" binding:"required.uuid64"`
}
