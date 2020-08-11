package dinamic_programing

import (
	"sort"
)

// You're given a non-empty array of arrays where each subarray holds three integers and represents a disk.
// These integers denote each disk's width, depth, and height, respectively. Your goal is to stack up the disks
// and to maximize the total height of the stack. A disk must have a strictly smaller width, depth,
// and height than any other disk below it. Write a function that returns an array of the disks in the final stack,
// starting with the top disk and ending with the bottom disk. Note that you can't rotate disks; in other words,
// the integers in each subarray must represent [width, depth, height] at all times.
// You can assume that there will only be one stack with the greatest total height.
// Sample Input disks=[
				//    [2, 1, 2],
				//    [3, 2, 3],
				//    [2, 2, 8],
				//    [2, 3, 4],
				//    [1, 3, 1],
				//    [4, 4, 5]
				//  ],
// Sample Output result=[
				//    [2, 1, 2],
				//    [3, 2, 3],
				//    [4, 4, 5]
				//  ] //10 (2 +3+ 5) // stacking disks is the tallest height we can get by following the rules laid out above.

type Disk []int
type Disks []Disk

func(disks Disks) Len()int {return len(disks)}
func(disks Disks) Swap(i, j int) {disks[i], disks[j] = disks[j], disks[i]}
func(disks Disks) Less(i,j int)bool {return disks[i][2] < disks[j][2]}

func DiskStacking(input [][]int) [][]int {
	disks := make(Disks, len(input))
	for i, disk := range input {
		disks[i] = disk
	}
	sort.Sort(disks)
	heights := make([]int, len(disks))
	sequences := make([]int, len(disks))
	for i := range disks {
		heights[i] = disks[i][2]
		sequences[i] = -1
	}

	for i:=1; i<len(disks); i++ {
		disk := disks[i]
		for j:=0; j<i; j++ {
			otherDisk := disks[j]
			if areValidDimensions(otherDisk, disk) {
				if heights[i] <= disk[2]+heights[j] {
					heights[i] = disk[2]+heights[j]
					sequences[i] = j
				}
			}
		}
	}
	maxIndex := 0
	for i, height := range heights {
		if heights[maxIndex] <= height {
			maxIndex = i
		}
	}
	sequence := buildSequenceDiskStacking(disks, sequences, maxIndex)
	return sequence
}


func areValidDimensions(diskOne []int, diskTwo []int) bool {
	return diskOne[0] < diskTwo[0] && diskOne[1] < diskTwo[1] && diskOne[2] < diskTwo[2]
}

func buildSequenceDiskStacking(arr []Disk, sequences []int, index int) [][]int {
	result := [][]int{}
	for index != -1 {
		result = append([][]int{arr[index]}, result...)
		index = sequences[index]
	}
	return result
}
