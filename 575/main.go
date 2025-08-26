package main

import "fmt"


func main() {
    candyType := []int{1,1,2,2,3,3}
    numberOfCandies := distributeCandies(candyType)
    candyType2 := []int{1,1,2,3}
    numberOfCandies2 := distributeCandies(candyType2)
    candyType3 := []int{6,6,6,6}
    numberOfCandies3 := distributeCandies(candyType3)
    fmt.Print(numberOfCandies)
    fmt.Print(numberOfCandies2)
    fmt.Print(numberOfCandies3)
}


func distributeCandies(candyType []int) int {
    maximumCandies := len(candyType) / 2
    
    if len(ReduceUniques(candyType)) > maximumCandies {
        return maximumCandies
    }

    return len(ReduceUniques(candyType))
}

func ReduceUniques(arr []int) []int {
    var uniques []int
    uniques = append(uniques, arr[0])
    for _, value := range arr {
        isUnique := true
        for _, unique := range uniques {
            if value == unique {
               isUnique = false 
            }
        }
        if isUnique {
            uniques = append(uniques, value)
        }
    }
    return uniques
}
