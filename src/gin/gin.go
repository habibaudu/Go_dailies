// package main

// import (
// 	"net/http"
// 	"github.com/gin-gonic/gin"
// 	"errors"
// )

// type todo struct{
// 	ID     string  `json:"id"`
// 	Item   string   `json:"item"`
// 	Completed bool   `json:"completed"`
// }

// var todos =[]todo{
//  {ID:"1",Item:"Read A book",Completed :false},
//  {ID:"2",Item:"Watch a video",Completed:false},
//  {ID:"3",Item:"Record a video",Completed:false},
// }

// func getTodos(context *gin.Context){
// 	context.IndentedJSON(http.StatusOK,todos)
// }

// func addTodos(context *gin.Context){
// 	var newTodo todo

// 	if err := context.BindJSON(&newTodo); err != nil{
// 		return
// 	}

// 	todos = append(todos,newTodo)

// 	context.IndentedJSON(http.StatusCreated,newTodo)
// }

// func getTodoByID(id string)(*todo,error){

// 	for i,t := range todos{
// 		if t.ID == id{
// 			return &todos[i],nil
// 		}
// 	}

// 	return nil,errors.New("Todo not found")

// }

// func getTodo(context *gin.Context){
//   id := context.Param("id")
//   todo,err := getTodoByID(id)

//   if err != nil{
// 	context.IndentedJSON(http.StatusNotFound,gin.H{"Message":"Todo not found"})
// 	return
//   }

//   context.IndentedJSON(http.StatusOK,todo)
// }

// func toggleTodoStatus(context *gin.Context){
// 	id := context.Param("id")
// 	todo,err := getTodoByID(id)

// 	if err != nil{
// 		context.IndentedJSON(http.StatusNotFound,gin.H{"Message":"Todo not found"})
// 		return
// 	}
// 	todo.Completed = !todo.Completed
// 	context.IndentedJSON(http.StatusOK,todo)
// }

// func main(){
//     // our server
// 	router := gin.Default()
// 	router.GET("/todos",getTodos)
// 	router.POST("/todos",addTodos)
// 	router.PATCH("/todos/:id",toggleTodoStatus)
// 	router.GET("/todos/:id",getTodo)

// 	router.Run("localhost:9090")

// }



package main 

import ("net/http"
        "errors"
	    "github.com/gin-gonic/gin")

type todo struct {
	ID string `json:"id"`
	Item string `json:"item"`
	Completed bool `json:"completed"`
}

var todos = []todo{
{ID:"1",Item:"Watch video",Completed:false},
{ID :"2",Item:"Meditate on the word",Completed:false},
{ID:"3",Item:"Write some Code",Completed:false},
}

func getTodos(context *gin.Context){
	context.IndentedJSON(http.StatusOK,todos)
}

func addTodo(context *gin.Context){
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos,newTodo)

	context.IndentedJSON(http.StatusCreated,todos)
}

func getTodoById(id string)(*todo,error){
	for i,t :=range todos{
		if t.ID == id{
			return &todos[i],nil
		}
	}

	return nil,errors.New("Todo Not Found")
}

func getTodo(context *gin.Context){
	id := context.Param("id")

	todo,err := getTodoById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound,gin.H{"Message":"Todo not found"})
	}

	context.IndentedJSON(http.StatusOK,todo)
}

func toggleTodoStatus(context *gin.Context){
	id := context.Param("id")
	todo,err := getTodoById(id)

	if err !=nil{
		context.IndentedJSON(http.StatusNotFound,gin.H{"Message":"Todo not found"})
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK,todo)
}


func main (){
	router := gin.Default()
	router.GET("/todos",getTodos)
	router.POST("/todos",addTodo)
	router.GET("/todos/:id",getTodo)
	router.PATCH("/todos/:id",toggleTodoStatus)
	router.Run("localhost:4040")
}