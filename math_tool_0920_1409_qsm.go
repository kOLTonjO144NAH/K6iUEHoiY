// 代码生成时间: 2025-09-20 14:09:19
package main

import (
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
    "time"

    "github.com/astaxie/beego"
)

// MathService provides the math operations
type MathService struct {
    beego.Controller
}

// Add handles the addition operation
func (m *MathService) Add() {
    // Get the input parameters
    a, errA := strconv.ParseFloat(m.GetString("a"), 64)
    b, errB := strconv.ParseFloat(m.GetString("b"), 64)

    // Error handling
    if errA != nil || errB != nil {
        m.Data["json"] = map[string]string{
            "error": "Invalid input. Please provide valid numbers.",
        }
        m.ServeJSON()
        return
    }

    // Perform the addition
    result := a + b

    // Return the result
    m.Data["json"] = map[string]float64{
        "result": result,
    }
    m.ServeJSON()
}

// Subtract handles the subtraction operation
func (m *MathService) Subtract() {
    // Get the input parameters
    a, errA := strconv.ParseFloat(m.GetString("a"), 64)
    b, errB := strconv.ParseFloat(m.GetString("b"), 64)

    // Error handling
    if errA != nil || errB != nil {
        m.Data["json"] = map[string]string{
            "error": "Invalid input. Please provide valid numbers.",
        }
        m.ServeJSON()
        return
    }

    // Perform the subtraction
    result := a - b

    // Return the result
    m.Data["json"] = map[string]float64{
        "result": result,
    }
    m.ServeJSON()
}

// Multiply handles the multiplication operation
func (m *MathService) Multiply() {
    // Get the input parameters
    a, errA := strconv.ParseFloat(m.GetString("a"), 64)
    b, errB := strconv.ParseFloat(m.GetString("b"), 64)

    // Error handling
    if errA != nil || errB != nil {
        m.Data["json"] = map[string]string{
            "error": "Invalid input. Please provide valid numbers.",
        }
        m.ServeJSON()
        return
    }

    // Perform the multiplication
    result := a * b

    // Return the result
    m.Data["json"] = map[string]float64{
        "result": result,
    }
    m.ServeJSON()
}

// Divide handles the division operation
func (m *MathService) Divide() {
    // Get the input parameters
    a, errA := strconv.ParseFloat(m.GetString("a"), 64)
    b, errB := strconv.ParseFloat(m.GetString("b\), 64)

    // Error handling
    if errA != nil || errB != nil || b == 0 {
        m.Data["json"] = map[string]string{
            "error": "Invalid input. Please provide valid numbers and a non-zero divisor.",
        }
        m.ServeJSON()
        return
    }

    // Perform the division
    result := a / b

    // Return the result
    m.Data["json"] = map[string]float64{
        "result": result,
    }
    m.ServeJSON()
}

func main() {
    // Start the Beego application
    beego.Router("/math/add", &MathService{}, "post:Add")
    beego.Router("/math/subtract", &MathService{}, "post:Subtract")
    beego.Router("/math/multiply", &MathService{}, "post:Multiply")
    beego.Router("/math/divide", &MathService{}, "post:Divide")

    // Start the server
    fmt.Println("Math Tool server is running at 8080 port...
")
    beego.Run()
}
