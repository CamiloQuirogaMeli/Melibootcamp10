package main

func InsertionSort(list *[]int) {
   
   for i := 0; i < len(*list); i++ {
	   for j := i; j > 0; j-- {
		   item1 := (*list)[j-1]
		   item2 := (*list)[j]
		   if item2 < item1 {
			   (*list)[j-1], (*list)[j] = item2, item1
		   } else {
			   break
		   }
	   }
   }
} 

func SelectionSort(list *[]int) {
   
   for i := 0; i < len(*list); i++ {
	   minValue := (*list)[i]
	   minIndex := i
	   for j := i; j < len(*list); j++ {
		   value := (*list)[j]
		   if value < minValue {
			   minValue = value
			   minIndex = j
		   }
	   }
	   (*list)[minIndex], (*list)[i] = (*list)[i], (*list)[minIndex]
   }
} 

func BubbleSort(list *[]int) {

   changed := true
   for changed {
	   changed = false
	   for i := 1; i < len(*list); i++ {
		   item1 := (*list)[i-1]
		   item2 := (*list)[i]
		   if item1 > item2 {
			   (*list)[i-1], (*list)[i] = item2, item1
			   changed = true
		   }
	   }
   }
}