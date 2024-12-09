package solution

import (
	"fmt"
	"strconv"
)

type disk struct {
	OriginalSectors []int64
	FileSizes       map[int64]int64
	MaxFileId       int64
}

func parseDisk(data []byte) disk {
	res := disk{
		OriginalSectors: []int64{},
		FileSizes:       map[int64]int64{},
	}
	fileId := int64(0)
	file := true
	for _, rn := range string(data) {
		size, _ := strconv.ParseInt(string(rn), 10, 64)
		for range size {
			if file {
				res.OriginalSectors = append(res.OriginalSectors, fileId)
			} else {
				res.OriginalSectors = append(res.OriginalSectors, -1)
			}
		}
		if file {
			res.MaxFileId = fileId
			res.FileSizes[fileId] = size
			fileId += 1
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

func generateChecksum(data []int64) int64 {
	fmt.Println(data)
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
	panic("unimplemented")
}
