package main

import (
	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
}

func main() {
	dsn := "host=localhost user=taskuser password=taskpass dbname=postgres port=5432 sslmode=disable"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&Task{})

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, config)

	r := gin.Default()

	r.POST("/tasks", func(c *gin.Context) {
		var task Task
		c.BindJSON(&task)
		db.Create(&task)

		msg := &sarama.ProducerMessage{
			Topic: "todo-events",
			Value: sarama.StringEncoder("Task created: " + task.Title),
		}
		producer.SendMessage(msg)

		c.JSON(200, task)
	})

	r.Run(":8080")
}
