package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)
type xy_tuple struct {
    x int32
    y int32
    dir direction
}

type direction int
const (
    UPDOWN direction = 0
    LEFTRIGHT direction = 1
    INVALID  direction = 2
)

// Complete the minimumMoves function below.
func minimumMoves(grid []string, startX int32, startY int32, goalX int32, goalY int32) int32 {
    bytegrid := make([][]byte, 0)
    for _, str := range grid {
        bytegrid = append(bytegrid, []byte(str))
    }

    queue := make([]xy_tuple, 0)
    current := xy_tuple{startX, startY, INVALID}
    goal := xy_tuple{goalX, goalY, INVALID}
    fmt.Println("start", current, "goal", goal)

    atGoal := false
    moves := 0

    // while not at goal
    for !atGoal {
        // if at goal, break. we're done
        fmt.Println("current:", current)
        if current.x == goal.x && current.y == goal.y {
            atGoal = true
            break
        }

        // otherwise, add all directions (except last traveling direction and its inverse) to the queue
        // BTW x is up/down
        going := current.dir
        if going != UPDOWN {
            // Try down
            lowestValid := current.x
            for i := current.x + 1; int(i) < len(grid); i++ {
                if bytegrid[i][current.y] == '.' {
                    lowestValid++
                } else {
                    break
                }
            }
            for i := lowestValid; i > current.x; i-- {
                var next xy_tuple
                next.x = i
                next.y = current.y
                next.dir = UPDOWN
                queue = enqueue(queue, next)
            }
            // Try up
            highestValid := current.x
            for i := current.x - 1; i >= 0; i-- {
                if bytegrid[i][current.y] == '.' {
                    highestValid--
                } else {
                    break
                }
            }
            for i := highestValid; i < current.x; i++ {
                var next xy_tuple
                next.x = i
                next.y = current.y
                next.dir = UPDOWN
                queue = enqueue(queue, next)
            }
        }
        if going != LEFTRIGHT {
            // Try left
            leftmost := current.y
            for i := current.y - 1; i >= 0; i-- {
                if bytegrid[current.x][i] == '.' {
                    leftmost--
                } else {
                    break
                }
            }
            for i := leftmost; i < current.y; i++ {
                var next xy_tuple
                next.x = current.x
                next.y = i
                next.dir = LEFTRIGHT
                queue = enqueue(queue, next)
            }
            // Try right
            rightmost := current.y
            for i := current.y + 1; int(i) < len(grid[0]); i++ {
                if bytegrid[current.x][i] == '.' {
                    rightmost++
                } else {
                    break
                }
            }
            for i := rightmost; i > current.y; i-- {
                var next xy_tuple
                next.x = current.x
                next.y = i
                next.dir = LEFTRIGHT
                queue = enqueue(queue, next)
            }
        }
        // Assuming queue is never empty
        var nextPosition xy_tuple
        fmt.Println(queue)
        queue, nextPosition = dequeue(queue)
        if nextPosition.dir != current.dir {
            moves++
        }
        current = nextPosition
    }

    return int32(moves)
}

func enqueue(queue []xy_tuple, pos xy_tuple) []xy_tuple{
    return append(queue, pos)
}
func dequeue(queue []xy_tuple) ([]xy_tuple, xy_tuple) {
    ret := queue[0]
    queue = queue[1:]
    return queue, ret
}
func isEmpty(queue []xy_tuple) bool {
    return len(queue) == 0
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    var grid []string

    for i := 0; i < int(n); i++ {
        gridItem := readLine(reader)
        grid = append(grid, gridItem)
    }

    startXStartY := strings.Split(readLine(reader), " ")

    startXTemp, err := strconv.ParseInt(startXStartY[0], 10, 64)
    checkError(err)
    startX := int32(startXTemp)

    startYTemp, err := strconv.ParseInt(startXStartY[1], 10, 64)
    checkError(err)
    startY := int32(startYTemp)

    goalXTemp, err := strconv.ParseInt(startXStartY[2], 10, 64)
    checkError(err)
    goalX := int32(goalXTemp)

    goalYTemp, err := strconv.ParseInt(startXStartY[3], 10, 64)
    checkError(err)
    goalY := int32(goalYTemp)

    result := minimumMoves(grid, startX, startY, goalX, goalY)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
