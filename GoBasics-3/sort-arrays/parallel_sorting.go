package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
  "math"
)

// A fuction to swap two elements of a slice
func Swap(arr []int, i int) {
  arr[i], arr[i+1] = arr[i+1], arr[i]
}

// Used for sorting each of four parts of the initial array
func BubbleSortPart(arr []int, c chan bool, ind int) {
  
  if ind != -1 {
    fmt.Println("Goroutine ", ind, " is sorting subarray: ", arr)
  }

  for i := 0; i < len(arr)-1; i++ {
    ok := true
    for j := 0; j < len(arr)-1-i; j++ {
      if arr[j] > arr[j+1] {
        Swap(arr, j)
        ok = false
      }
    }
    if ok {
      break
    }
  }

  if c != nil {
    c <- true
  }
}

// Merge two sorted slices into one
func MergeSorted(arr1, arr2 []int) []int {

  merged := make([]int, len(arr1) + len(arr2))
  mergeInd := 0
  arr1Ind := 0
  arr2Ind := 0

  for i := 0; i < len(merged); i++ {
    // if first slice is over
    if arr1Ind >= len(arr1) {
      for arr2Ind < len(arr2) {
        merged[mergeInd] = arr2[arr2Ind]
        mergeInd += 1
        arr2Ind += 1
      }
      break
    }
    // if first slice is over
    if arr2Ind >= len(arr2) {
      for arr1Ind < len(arr1) {
        merged[mergeInd] = arr1[arr1Ind]
        mergeInd += 1
        arr1Ind += 1
      }
      break
    }
    // else
    if arr1[arr1Ind] < arr2[arr2Ind] {
      merged[mergeInd] = arr1[arr1Ind]
      arr1Ind += 1
    } else {
      merged[mergeInd] = arr2[arr2Ind]
      arr2Ind += 1
    }

    mergeInd += 1
  }

  return merged
}


func SplitSort(arr []int) []int {

  arrCopy := make([]int, len(arr))
  copy(arrCopy, arr)  

  // If length of the array is less then 4, each of four parts is already sorted.
  if len(arr) <= 4 {
    BubbleSortPart(arrCopy, nil, -1)
    return arrCopy
  } 

  // Break to four parts
  step := int(math.Round(float64(len(arr))/4))
  var parts [4][]int
  parts[0] = arrCopy[: step]
  parts[1] = arrCopy[step : 2*step]
  parts[2] = arrCopy[2*step : 3*step]
  parts[3] = arrCopy[3*step :]

  // Start parallel sorting
  c := make(chan bool, 4)
  for i := 0; i < len(parts); i++ {
    go BubbleSortPart(parts[i], c, i)
  }
  // Wait for all sorts to finish
  for i := 0; i < len(parts); i++ {
    <- c
  }

  res := parts[0]
  for i:= 1; i < len(parts); i++ {
    res = MergeSorted(res, parts[i])
  }
  
  return res
}


func main() {
  // Reading input
  scanner := bufio.NewScanner(os.Stdin)
  fmt.Println("Enter a list of numbers separated by spaces:")
  
  scanner.Scan()
  listString := strings.Fields(scanner.Text())
  
  //Casting input to int
  var listInt = make([]int, 0, 10)
  for _, val := range listString {
    converted, err := strconv.Atoi(val)
    if err != nil {
      panic(err)
    }
    listInt = append(listInt, converted)
  }
  
  fmt.Println("")
  fmt.Println("You entered: ", listInt)

  sorted := SplitSort(listInt)
  
  fmt.Println("")
  fmt.Println("Sorted: ", sorted)
  fmt.Println("")
}

