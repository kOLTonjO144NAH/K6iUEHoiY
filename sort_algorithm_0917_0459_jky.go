// 代码生成时间: 2025-09-17 04:59:51
package main

import (
    "fmt"
    "github.com/astaxie/beego"
)

// Sorter is an interface that groups the Sort method.
// Implementers can sort a slice of any type.
type Sorter interface {
    Sort([]int)
}

// BubbleSorter sorts a slice of integers using the bubble sort algorithm.
type BubbleSorter struct{}

// Sort sorts a slice of integers using the bubble sort algorithm.
func (s BubbleSorter) Sort(arr []int) {
    n := len(arr)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j] // swap
            }
        }
    }
}

// QuickSorter sorts a slice of integers using the quick sort algorithm.
type QuickSorter struct{}

// Sort sorts a slice of integers using the quick sort algorithm.
func (s QuickSorter) Sort(arr []int) {
    if len(arr) < 2 {
        return // Do nothing if the array has 0 or 1 elements
    }
    quickSort(arr, 0, len(arr)-1)
}

// quickSort is a recursive function that implements the quick sort algorithm.
func quickSort(arr []int, low int, high int) {
    if low < high {
        pi := partition(arr, low, high)
        quickSort(arr, low, pi-1) // Before pi
        quickSort(arr, pi+1, high) // After pi
    }
}

// partition is a helper function that rearranges the elements in arr[l..r]
// such that all elements less than the partition element are on the left of pi,
// and all elements greater than the partition element are on the right of pi.
func partition(arr []int, low int, high int) int {
    i := low - 1 // Index of smaller element
    pivot := arr[high] // Pivot
    for j := low; j <= high-1; j++ {
        // If current element is smaller than or equal to pivot
        if arr[j] <= pivot {
            i++ // increment index of smaller element
            arr[i], arr[j] = arr[j], arr[i] // swap
        }
    }
    arr[i+1], arr[high] = arr[high], arr[i+1] // swap pivot element with element at index i+1
    return i + 1
}

// SortAlgorithmController is a Beego controller that exposes sorting functionality via HTTP endpoints.
type SortAlgorithmController struct {
    beego.Controller
}

// Get handles GET requests and returns a simple welcome message.
func (c *SortAlgorithmController) Get() {
    c.Data["json"] = map[string]string{"message": "Welcome to the Sort Algorithm Service!"}
    c.ServeJSON()
}

// Post handles POST requests and sorts an array of integers provided in the request body.
func (c *SortAlgorithmController) Post() {
    var arr []int
    if err := json.Unmarshal(c.Ctx.Input.RequestBody, &arr); err != nil {
        c.Data["json"] = map[string]string{"error": "Invalid input. Please provide an array of integers."}
        c.ServeJSON()
        c.Ctx.Output.SetStatus(400)
        return
    }

    sorter := BubbleSorter{} // or QuickSorter{} to use quick sort
    sorter.Sort(arr)

    c.Data["json"] = map[string][]int{"sortedArray": arr}
    c.ServeJSON()
}

func main() {
    beego.Router("/", &SortAlgorithmController{})
    beego.Router("/sort", &SortAlgorithmController{}, "post:Post")
    beego.Run()
}