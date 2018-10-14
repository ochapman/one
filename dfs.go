/*
* File Name:	dfs.go
* Description:	Depth first search
* Author:	Chapman Ou <ochapman.cn@gmail.com>
* Created:	2018-10-13
 */

package one

import (
	"container/list"
)

type Pos struct {
	X int32
	Y int32
}

type Meta struct {
	Width  int32
	Height int32
}

type Block struct {
	Pos
	Visited bool
	Exist   bool
	Val     byte
}

type Blocks map[Pos]*Block

var GBlocks Blocks
var GMeta Meta
var GPaths *list.List

type Paths []Pos

func FindNextPos(pos Pos) []Pos {
	nextSteps := []Pos{
		{1, 0},  //Right
		{0, -1}, //Down
		{-1, 0}, //Left
		{0, 1},  //Up
	}
	nextPos := make([]Pos, 4)
	for i, ns := range nextSteps {
		nextPos[i].X, nextPos[i].Y = pos.X+ns.X, pos.Y+ns.Y
	}
	return nextPos
}

func DFS(curPos Pos, dstPos Pos) int {

	if curPos.X == dstPos.X && curPos.Y == dstPos.Y {
		isFinished := true
		for p, _ := range GBlocks {
			if GBlocks[p].Exist {
				if !GBlocks[p].Visited {
					isFinished = false
				}
			}
		}

		if isFinished {
			//	for e := GPaths.Front(); e != nil; e = e.Next() {
			//		fmt.Printf("%v ", e.Value)
			//	}
			return 1
		}
	}

	nextPos := FindNextPos(curPos)

	for _, np := range nextPos {
		//check touch edge
		if np.X < 0 || np.X >= GMeta.Width || np.Y < 0 || np.Y >= GMeta.Height {
			continue
		}
		//fmt.Printf("np: %v\n", np)
		if GBlocks[np].Exist && !GBlocks[np].Visited {
			GBlocks[np].Visited = true
			GPaths.PushBack(np)
			if DFS(np, dstPos) == 1 {
				return 1
			}
			GBlocks[np].Visited = false
			e := GPaths.Back()
			GPaths.Remove(e)
			//fmt.Printf("Remove: %v, Len of list: %d\n", e, GPaths.Len())
		}
	}
	return 0
}

func PrintPath(gp *list.List, gb Blocks) Blocks {
	newGB := make(Blocks)
	for e := gp.Front(); e != nil; e = e.Next() {
		var newE Pos
		if e.Next() == nil {
			continue
		}
		nextE := e.Next().Value.(Pos)
		curE := e.Value.(Pos)
		var arrow byte

		if nextE.X == curE.X {
			newE.X = curE.X * 2
			if nextE.Y > curE.Y {
				newE.Y = curE.Y*2 + 1
				arrow = '^'
			} else {
				newE.Y = curE.Y*2 - 1
				arrow = 'v'
			}

		} else {
			newE.Y = curE.Y * 2
			if nextE.X > curE.X {
				newE.X = curE.X*2 + 1
				arrow = '>'

			} else {
				newE.X = curE.X*2 - 1
				arrow = '<'
			}
		}
		newGB[newE] = &Block{
			Pos: Pos{
				X: newE.X,
				Y: newE.Y,
			},
			Val: arrow,
		}
	}
	for p, _ := range gb {
		newP := Pos{X: p.X * 2, Y: p.Y * 2}
		newGB[newP] = gb[p]
	}
	return newGB
}
