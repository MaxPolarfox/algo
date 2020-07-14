package arrays

// Imagine that you're a teacher who's just graded the final exam in a class.
// You have a list of student scores on the final exam in a particular order (not necessarily sorted), and you want to reward your students.
// You decide to do so fairly by giving them arbitrary rewards following two rules:
//		1) All students must receive at least one reward.
//		2) Any given student must receive strictly more rewards than an adjacent student (a student immediately to the left or to the right)
//		with a lower score and must receive strictly fewer rewards than an adjacent student with a higher score.
//	Write a function that takes in a list of scores and returns the minimum number of rewards that you must give out to students to satisfy the two rules.
//	You can assume that all students have different scores; in other words, the scores are all unique.

//	Sample Input scores = [8, 4, 2, 1, 3, 6, 8, 9, 5]
//	Sample Output 25 // you would give out the following rewards:  [4, 3, 2, 1, 2, 3, 4, 5, 1]

// Time: O(n) Space: O(n)
func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	fill(rewards, 1)

	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}

	for i := len(scores) - 2; i >= 0; i-- {
		if scores[i] > scores[i+1] {
			rewards[i] = max(rewards[i+1]+1, rewards[i])
		}
	}

	return sum(rewards)
}

func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func fill(arr []int, val int) {
	for i := range arr {
		arr[i] = val
	}
}
