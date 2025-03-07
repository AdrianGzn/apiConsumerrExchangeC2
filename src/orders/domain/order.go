package domain
import "time"

type Order struct {
	Id int32 			`json:"id"`
	Name  string  		`json:"name"`
	Description string 	`json:"description"`
	Price int32 		`json:"price"`
	UserName string		`json:"userName"`
	UserCellphone string`json:"cellPhone"`
	Status string		`json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

func NewOrder(name string, description string, price int32, userName string, userCellphon string, createdAt time.Time) *Order {
	return &Order{Name: name, Description: description, Price: price, UserName: userName, UserCellphone: userCellphon, CreatedAt: createdAt}
}