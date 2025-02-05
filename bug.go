func myFunc(a []int) {
    for i := range a {
        a[i]++ // This line might cause unexpected behavior if the slice is nil
    }
}