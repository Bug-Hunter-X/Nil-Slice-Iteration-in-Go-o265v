func myFunc(a []int) {
    if a == nil {
        return // Handle the nil case appropriately
    }
    for i := range a {
        a[i]++
    }
} 

//Alternative solution using a loop:
func myFunc(a []int) {
    if a == nil {
        return
    }
    for i := 0; i < len(a); i++ {
        a[i]++
    }
}