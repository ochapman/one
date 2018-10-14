/*
* File Name:	onestrokes.go
* Description:
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2018-10-13
 */

package main

import (
	"container/list"
	"fmt"
	"os"

	"github.com/ochapman/one"
)

/*
	one.GBlocks = one.Blocks{
		one.Pos{X: 0, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 0,
			},
			Exist: false,
		},
		one.Pos{X: 1, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 0,
			},
			Exist: true,
		},
		one.Pos{X: 2, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 0,
			},
			Exist: true,
		},
		one.Pos{X: 3, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 0,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 0,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 0}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 0,
			},
			Exist: true,
		},
		one.Pos{X: 0, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 1,
			},
			Exist: true,
		},
		one.Pos{X: 1, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 1,
			},
			Exist: true,
		},
		one.Pos{X: 2, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 1,
			},
			Exist: false,
		},
		one.Pos{X: 3, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 1,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 1,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 1}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 1,
			},
			Exist: true,
		},
		one.Pos{X: 0, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 2,
			},
			Exist: true,
		},
		one.Pos{X: 1, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 2,
			},
			Exist: false,
		},
		one.Pos{X: 2, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 2,
			},
			Exist: true,
		},
		one.Pos{X: 3, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 2,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 2,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 2}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 2,
			},
			Exist: true,
		},
		one.Pos{X: 0, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 3,
			},
			Exist: true,
		},
		one.Pos{X: 1, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 3,
			},
			Exist: false,
		},
		one.Pos{X: 2, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 3,
			},
			Exist: true,
		},
		one.Pos{X: 3, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 3,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 3,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 3}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 3,
			},
			Exist: false,
		},
		one.Pos{X: 0, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 4,
			},
			Exist: true,
		},
		one.Pos{X: 1, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 4,
			},
			Exist: true,
		},
		one.Pos{X: 2, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 4,
			},
			Exist: true,
		},
		one.Pos{X: 3, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 4,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 4,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 4}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 4,
			},
			Exist: false,
		},
		one.Pos{X: 0, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 0,
				Y: 5,
			},
			Exist: true,
		},
		one.Pos{X: 1, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 1,
				Y: 5,
			},
			Exist: true,
		},
		one.Pos{X: 2, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 2,
				Y: 5,
			},
			Exist: true,
		},
		one.Pos{X: 3, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 3,
				Y: 5,
			},
			Exist: true,
		},
		one.Pos{X: 4, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 4,
				Y: 5,
			},
			Exist: true,
		},
		one.Pos{X: 5, Y: 5}: &one.Block{
			Pos: one.Pos{
				X: 5,
				Y: 5,
			},
			Exist: true,
		},
	}
*/
func main() {
	var Start one.Pos
	var End one.Pos

	one.GBlocks = make(one.Blocks)
	one.GMeta = one.Meta{Width: 6, Height: 8}

	one.GPaths = list.New()
	var i, j int32
	var v int32

	f, err := os.Open("./data.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Open err: %v", err)
		os.Exit(-1)
	}

	var val int32
	var exist bool
	for i = one.GMeta.Height - 1; i >= 0; i-- {
		for j = 0; j < one.GMeta.Width; j++ {
			pos := one.Pos{X: j, Y: i}
			val = 0
			fmt.Fscanf(f, "%d", &val)
			if val == 0 {
				exist = false
			} else {
				exist = true
			}
			block := &one.Block{
				Pos:   pos,
				Exist: exist,
			}

			one.GBlocks[pos] = block

			fmt.Printf("i: %d, j: %d, val: %d, GBlocks: %v, pos: %v\n", i, j, val, one.GBlocks[pos], pos)
		}
	}
	fmt.Println("Please input Start:")
	fmt.Fscanf(f, "%d %d\n", &Start.X, &Start.Y)

	fmt.Println("Please input End:")
	fmt.Fscanf(f, "%d %d\n", &End.X, &End.Y)

	fmt.Printf("start: %v, end: %v\n", Start, End)
	for i = one.GMeta.Height - 1; i >= 0; i-- {
		for j = 0; j < one.GMeta.Width; j++ {
			pos := one.Pos{X: j, Y: i}
			//fmt.Printf("pos: %vGBlocks: %v\n", pos, one.GBlocks)
			if one.GBlocks[pos].Exist {
				v = 1
			} else {
				v = 0
			}
			fmt.Printf("%v", v)
		}
		fmt.Printf("\n")
	}
	one.GBlocks[Start].Visited = true
	one.GPaths.PushBack(Start)
	one.DFS(Start, End)

	newGB := one.PrintPath(one.GPaths, one.GBlocks)
	for i = one.GMeta.Height*2 - 1; i >= 0; i-- {
		for j = 0; j < one.GMeta.Width*2; j++ {
			pos := one.Pos{X: j, Y: i}
			//fmt.Printf("pos: %v\n", pos)
			if _, ok := newGB[pos]; ok {
				if newGB[pos].Exist {
					v = 1
				} else {
					v = 0
				}
				if newGB[pos].Val > 0 {
					fmt.Printf("%c", newGB[pos].Val)
				} else {
					fmt.Printf("%d", v)
				}
			} else {
				fmt.Printf("%v", " ")
			}
		}
		fmt.Printf("\n")
	}
	//Print Path

}
