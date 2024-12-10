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
	src := make([]int64, len(dsk.OriginalSectors))
	for i := range dfrgd {
		dfrgd[i] = dsk.OriginalSectors[i]
		src[i] = dsk.OriginalSectors[i]
	}
	// fuck it, we go dumb.
	backint := 1
	fileCollection := []int64{}
	var nxt int64
	for backint < len(dfrgd) && len(src) > 1 {
		nxt, src = src[len(src)-backint], src[len(src)-backint:]

		if len(fileCollection) > 0 && (nxt == -1 || nxt != fileCollection[0]) {
			// end of file
			// write buffered file to first available "slot"
			ln := 0
			idx := 0
			for i, c := range dfrgd {
				if c == -1 {
					ln += 1
					if ln == len(fileCollection) {
						for iii, c := range dfrgd {
							if c == fileCollection[0] {
								dfrgd[iii] = -1
							}
						}
						for ii := range len(fileCollection) {
							dfrgd[ii+idx] = fileCollection[0]
						}
					}
				} else {
					idx = i
					ln = 0
				}
			}
			fileCollection = []int64{}
		} else {
			if nxt != -1 {
				fileCollection = append(fileCollection, nxt)
			}
		}
		backint -= 1
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
