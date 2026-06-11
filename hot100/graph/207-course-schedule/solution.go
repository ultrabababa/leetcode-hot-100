package course_schedule

func canFinish(numCourses int, prerequisites [][]int) bool {
	// 1. 初始化
	// inDegree 记录每个节点的入度 (有几个前置节点)
	inDegree := make([]int, numCourses)

	// adj 记录每个节点的邻接表 (当前节点指向了哪些节点)
	adj := make([][]int, numCourses)

	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}

	// 2. 构建图和入度表
	for _, p := range prerequisites {
		// p[0] 是想要学习的课程，p[1] 是必须先修的课程
		// 所以箭头方向是：p[1] -> p[0]
		cur := p[0]
		pre := p[1]

		adj[pre] = append(adj[pre], cur)
		inDegree[cur]++
	}

	// 3. 将所有入度为 0 的节点 (可以直接上的课) 放入队列
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	// count 记录成功学完的课程数量
	count := 0

	// 4. 开始 BFS
	for len(queue) > 0 {
		// 出队：当前上完的课
		node := queue[0]
		queue = queue[1:]
		count++

		// 遍历这门课的所有后续课程
		for _, nextCourse := range adj[node] {
			// 因为前置课上完了，后续课程的限制少了一个
			inDegree[nextCourse]--

			// 如果后续课程的所有限制都解除了 (入度为 0)，可以去上了
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	// 5. 检查是否上完了所有课
	return count == numCourses
}
