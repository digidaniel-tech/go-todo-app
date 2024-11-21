package drivers

import (
	"encoding/csv"
	"os"
	"reflect"
	"time"
	"todo/internal/models"
)

const filename = "data.csv"

func WriteCsv(todos []models.Todo) {
    file, err := os.Create(filename)
    if err != nil {
        panic(err)
    }
    
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    headers := readFields(todos[0]) 
    data := readData(todos, headers)

    writer.Write(headers)
    for _, row := range data {
        writer.Write(row)
    }
}

func readData(todos []models.Todo, headers []string) [][]string {
    data := make([][]string, len(todos))

    for i, todo := range todos {
        row := make([]string, len(headers))
        reflection := reflect.ValueOf(todo)

        for x, header := range headers {
            field := reflection.FieldByName(header)
            row[x] = field.String()
        }

        data[i] = row
    }

    return data
}

func readFields(todo models.Todo) []string {
    reflection := reflect.ValueOf(todo)
    types := reflection.Type()
    
    values := make([]string, reflection.NumField())

    for i := 0; i < reflection.NumField(); i++ {
        values[i] = types.Field(i).Name
    }

    return values
}

func timeToString(time time.Time) string {
    return time.String() 
}
