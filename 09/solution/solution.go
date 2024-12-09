package solution

import (
	"fmt"
	"strconv"
)

type Block struct {
	StartIndex int64
	Length     int64
	Id         int64
}

type disk struct {
	OriginalSectors []int64
	FileSizes       map[int64]int64
	MaxFileId       int64

	FileBlocks  []Block
	SpaceBlocks []Block
}

func parseDisk(data []byte) disk {
	res := disk{
		OriginalSectors: []int64{},
		FileSizes:       map[int64]int64{},
		FileBlocks:      []Block{},
		SpaceBlocks:     []Block{},
	}
	fileId := int64(0)
	file := true
	index := int64(0)
	for _, rn := range string(data) {
		size, _ := strconv.ParseInt(string(rn), 10, 64)
		sidx := index
		for range size {
			if file {
				res.OriginalSectors = append(res.OriginalSectors, fileId)
			} else {
				res.OriginalSectors = append(res.OriginalSectors, -1)
			}
			index += 1
		}
		if file {
			res.MaxFileId = fileId
			res.FileSizes[fileId] = size
			res.FileBlocks = append(res.FileBlocks, Block{
				StartIndex: sidx,
				Id:         fileId,
				Length:     size,
			})
			fileId += 1
		} else {
			res.SpaceBlocks = append(res.SpaceBlocks, Block{
				StartIndex: sidx,
				Id:         -1,
				Length:     size,
			})
		}
		file = !file
	}
	return res
}

func defragmentDisk(dsk disk) []int64 {
	dfrgd := []int64{}
	movingfile := dsk.MaxFileId
	movingfileremainingsectors := dsk.FileSizes[movingfile]
	seenfiles := int64(0)
	freespace := false
	for _, sctr := range dsk.OriginalSectors {
		if freespace {
			// gotcha on there being data left
			if movingfileremainingsectors > 0 {
				movingfileremainingsectors -= 1
				dfrgd = append(dfrgd, movingfile)
			}
			continue
		}
		if seenfiles >= movingfile && sctr != -1 {
			// we have hit ourselves
			continue
		}
		if seenfiles >= movingfile && sctr == -1 {
			freespace = true
			continue
		}

		if sctr == -1 {
			movingfileremainingsectors -= 1
			dfrgd = append(dfrgd, movingfile)
			if movingfileremainingsectors == 0 {
				movingfile -= 1
				movingfileremainingsectors = dsk.FileSizes[movingfile]
			}
		} else {
			dfrgd = append(dfrgd, sctr)
			if movingfile == sctr {
				movingfileremainingsectors -= 1
			}
			seenfiles = sctr
		}
	}
	return dfrgd
}

func defragmentFiles(dsk disk) []int64 {
	dfrgd := make([]int64, len(dsk.OriginalSectors))
	for i := range dfrgd {
		dfrgd[i] = -1
	}
	for {
		fmt.Println(dsk.FileBlocks)
		fmt.Println(dsk.SpaceBlocks)
		if len(dsk.FileBlocks) == 1 {
			// last one, just put in place
			for i := range dsk.FileBlocks[0].Length {
				dfrgd[dsk.FileBlocks[0].StartIndex+i] = dsk.FileBlocks[0].Id
			}
			break
		}
		var fblk Block
		if len(dsk.FileBlocks) == 2 {
			dsk.FileBlocks, fblk = []Block{dsk.FileBlocks[0]}, dsk.FileBlocks[1]
		} else {
			dsk.FileBlocks, fblk = dsk.FileBlocks[:len(dsk.FileBlocks)-1], dsk.FileBlocks[len(dsk.FileBlocks)-1]
		}
		moved := false
		for i, spc := range dsk.SpaceBlocks {
			if spc.Length == fblk.Length {
				// insert in place and remove space

				if i == 0 && len(dsk.SpaceBlocks) > 1 {
					dsk.SpaceBlocks = dsk.SpaceBlocks[i+1:]
				} else if i == 0 && len(dsk.SpaceBlocks) == 1 {
					dsk.SpaceBlocks = []Block{}
				} else {
					dsk.SpaceBlocks = append(dsk.SpaceBlocks[:i-1], dsk.SpaceBlocks[i+1:]...)
				}
				for ii := range fblk.Length {
					dfrgd[spc.StartIndex+ii] = fblk.Id
				}
				moved = true

			}
			if spc.Length > fblk.Length {
				// insert and then size down/"move" space
				for ii := range fblk.Length {
					dfrgd[spc.StartIndex+ii] = fblk.Id
				}
				dsk.SpaceBlocks[i].Length = dsk.SpaceBlocks[i].Length - fblk.Length
				dsk.SpaceBlocks[i].StartIndex = dsk.SpaceBlocks[i].StartIndex + fblk.Length

				moved = true

			}
			if moved {
				// insert a free block at the location then break
				dsk.SpaceBlocks = append(dsk.SpaceBlocks, Block{
					Id:         -1,
					StartIndex: fblk.StartIndex,
					Length:     fblk.Length,
				})
				break
			}
		}
		if !moved {
			// insert in original position
			for i := range fblk.Length {
				dfrgd[fblk.StartIndex+i] = fblk.Id
			}
		}
	}
	return dfrgd
}

func generateChecksum(data []int64) int64 {
	acc := int64(0)
	for i, sctr := range data {
		if sctr != -1 {
			acc += (sctr * int64(i))
		}
	}
	return acc
}

func ComputeSolutionOne(data []byte) int64 {
	dsk := parseDisk(data)
	defragged := defragmentDisk(dsk)
	res := generateChecksum(defragged)
	return res
}

func ComputeSolutionTwo(data []byte) int64 {
	dsk := parseDisk(data)
	fmt.Println(dsk.OriginalSectors)
	defragged := defragmentFiles(dsk)
	fmt.Println(defragged)
	res := generateChecksum(defragged)
	return res
}
