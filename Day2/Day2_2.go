package main
 
import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)
 
 
func main() {
    file, err := os.Open("in2.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()
 
    scanner := bufio.NewScanner(file)
    scanner.Scan()
    s := strings.Split(scanner.Text(), ",")
    var nums [276]int
    for a,b := range s{
        nums[a], _ = strconv.Atoi(b)
    }
    out:
    for noun := 0; noun < 100; noun++ {
        for verb := 0; verb < 100; verb++{
            var nums2 [len(nums)]int
            for i := range nums {
              nums2[i] = nums[i]
            }
            i := 0
            nums2[1] = noun
            nums2[2] = verb
            for{
                if nums2[i] == 1{
                    nums2[ nums2[i+3]] = nums2[ nums2[i+1]] + nums2[nums2[i+2]]
                } else if nums2[i] == 2{
                    nums2[ nums2[i+3]] = nums2[nums2[i+1]] * nums2[nums2[i+2]]
                } else if nums2[i] == 99{
                    break
                }
                i += 4
            }
            if nums2[0] == 19690720{
                fmt.Println(100 * noun + verb)
                break out
            }
        }
    }
 
}
