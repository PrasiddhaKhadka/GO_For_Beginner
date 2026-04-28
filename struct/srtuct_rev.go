package struct

import "time"

type Order struct{
	id  int
	amount float
	status string
	created time.time
}

//  SIMPLE CONSTRUCTOR CONCEPT 
func newOrder(id int, amount float, status string, created time.time) *Order{
	order := Order{
		id= id
		amount= amount
		status= status
		created= time.Now()
	}
	return &order
}


// 🙀🙀🙀🙀🙀🙀 WONT WORK !!! :-> RATHE THAN CHANGING VALUE A NEW STATUS values is set in the memory 
// funct (o Order) changeStatus(status string){
// 	o.status = status
// }



// 🔥🔥🔥🔥🔥🔥 TO MAKE IT WORK!! 🔥🔥🔥🔥🔥🔥
func (o *Order) changeStatus(status string){
	o.status = status
}


func main(){
	orderObject1 := Order{
		id: 1,
		amount: 60.0,
		status:packing,
		created: time.Now()
	}

	//  STATUS CHANGED LATER

	orderObject1.changeStatus("Delivered")

	fmt.Println("Object ", orderObject1)
}