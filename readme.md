Returns an array of time ranges spanning from rangeStart to rangeEnd.  

Each element is a partial range, the duration of it is defined by the step parameter.  

Each element is a struct with two elements, element `S` is the partial range start Time, element `E` is the partial range end Time  

The end of one range is the start of the next one, this could cause overlapping if your usage is `>=` start and `<=` end  
(you should use `>=` start and `<` end)  

If that is not possible you can subtract the smaller granularity of time you need from the end date while looping the array  

