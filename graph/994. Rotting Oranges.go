package main
import "fmt"

func orangesRotting(grid [][]int) int {
    rows := len(grid)
    cols := len(grid[0])

	queue := [][]int{}
	fresh := 0

	for r:=0; r < rows; r++ {
		for c:=0; c < cols; c++ {
			if grid[r][c] == 2 {
				queue = append(queue, []int{r, c})
			}else if grid[r][c] == 1 {
				fresh++
			}
		}
	}

	if fresh == 0 {
		return 0
	}
	minute := 0
	directions := [][]int {
		{-1, 0},//top
		{1, 0},//button
		{0, -1},//left
		{0, 1},//right
	}
	for len(queue) > 0 {
		rottedMenute := false
		size := len(queue)
		for i:=0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]
			r := curr[0]
			c := curr[1]
			for _, dir := range directions {
				nr := r + dir[0]
				nc := c + dir[1]

				if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
					continue
				}
				if grid[nr][nc] != 1 {
					continue
				}

				grid[nr][nc] = 2
				fresh--

				queue = append(queue, []int{nr ,nc})
				rottedMenute = true
			}
		}
		if rottedMenute == true {
			minute++
		}
	} 
	return minute
}


func main() {
	grid := [][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	}

	fmt.Println(orangesRotting(grid))
}