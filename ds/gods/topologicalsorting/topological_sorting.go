package topologicalsorting

/*
210. Course Schedule II
There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take course ai.

For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
Return the ordering of courses you should take to finish all courses. If there are many valid answers, return any of them. If it is impossible to finish all courses, return an empty array.

Example 1:

Input: numCourses = 2, prerequisites = [[1,0]]
Output: [0,1]
Explanation: There are a total of 2 courses to take. To take course 1 you should have finished course 0. So the correct course order is [0,1].
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		indegree[v[0]]++
	}
	queue := make([]int, 0, numCourses)
	for i, v := range indegree {
		if v == 0 {
			queue = append(queue, i)
		}
	}
	res := make([]int, 0, numCourses)
	for len(queue) > 0 {
		h := queue[0]
		queue = queue[1:]
		res = append(res, h)
		for _, v := range prerequisites {
			if v[1] == h {
				indegree[v[0]]--
				if indegree[v[0]] == 0 {
					queue = append(queue, v[0])
				}
			}
		}
	}

	if len(res) == numCourses {
		return res
	}

	return []int{}
}
