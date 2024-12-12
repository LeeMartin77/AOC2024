package solution

import (
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
	reversedfiles := make([]Block, len(dsk.FileBlocks))
	ln := len(dsk.FileBlocks)
	i := 0
	for ln > 0 {
		ln -= 1
		reversedfiles[i] = dsk.FileBlocks[ln]
		i += 1
	}
	for _, fl := range reversedfiles {
		placed := false
		for i, empty := range dsk.SpaceBlocks {
			if fl.StartIndex < empty.StartIndex {
				// jump out because we've passed
				break
			}
			if empty.Length >= fl.Length {
				dsk.SpaceBlocks[i] = Block{
					Id:         -1,
					Length:     empty.Length - fl.Length,
					StartIndex: empty.StartIndex + fl.Length,
				}
				for i := range fl.Length {
					dfrgd[i+empty.StartIndex] = fl.Id
				}
				placed = true
				break
			}
		}
		if !placed {
			for i := range fl.Length {
				dfrgd[i+fl.StartIndex] = fl.Id
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
	defragged := defragmentFiles(dsk)
	res := generateChecksum(defragged)
	return res
}
