package main

import "fmt"
import "strings"

type slopeSettings struct {
	x int
	y int
}

func main() {
	input := ".....#............#....#####.##\n.#.#....#......#....##.........\n......#.#.#.....###.#.#........\n......#...#.....#####....#..##.\n...#............##...###.##....\n#.....#...#....#......##....##.\n#...#.#....#..#..##.##...#.....\n.......#..........#..#..#.#....\n.#.....#.#.......#..#...#....#.\n#..#.##.#..................###.\n...#.#.##...##.###.....#..#...#\n..#.#...#............#.......#.\n#..#.#..#.#....#...#.#.....#..#\n#......##....#..#.#.#........#.\n....#..#.#.#.##............#..#\n....#..#..#...#.#.##......#...#\n##...#...........#.....###.#...\n..#...#.#...#.#.....#....##.##.\n....##...##.#....#.....#.##....\n#........##......#......#.#.#.#\n....#.#.#.........##......#....\n.#......#...#.....##..#....#..#\n....#..#.#.....#..........#..#.\n..##...#..##................#.#\n.....#....#.#..#......#........\n........#..#.#......#.#........\n.....#.#....##.###....#...#....\n...##.#.......#....###..#......\n............##.#..#...#........\n#..###..#.....#.####...........\n.......##.....#......#......#..\n#........##..#.....##.......#.#\n#.##...#...#...#......##..#.#.#\n......#....##.#.#...#...##....#\n#..#....##.#......#.......##...\n.#..........#..........#....#.#\n#.....##......##....#..........\n..#.#.....#.#...#........#.....\n...#........#..#..#.##..##.....\n......###.....#..#...#.###...##\n.##.##.......#.......###...#...\n#.#..#.#.#....#.....###..#...##\n......#.##..........#.......##.\n#..#.#.........#.....##...##...\n..#...#....#....###.#......#...\n.....#..#.######.....#..#.#....\n..#.#.....#.....##.#....##.#.##\n...#.#.#....#....##..#..#.#.##.\n...........#.#...#..#..####....\n.........#####.#.#.#...#.##.#..\n.......#...#......#.##.#.##....\n....#.....#.....###..........#.\n.#.###....##.#..#..........#...\n#...#.........##.....####....#.\n##....##...#..........#........\n...#.#.#.#....#..........#.....\n.......#....#......##.......#..\n.#.#..#.........#.#.##....#....\n..#.............#..##.#.##..###\n.#.##..............#..#..##..#.\n..##.#..#......#...##..#.##...#\n......#..#....#....#....##..#..\n...#...##.............#..###...\n...##....#.#.##........#.....##\n....#.#.......#..###..#....####\n#...#...##..#.####..#...##....#\n.......#..#.##..#...#.#........\n###.#......#..##..#......#.##..\n#....#............#.....#......\n..##...#..##......#..#....#....\n.#..##...#....#.#...#...#..#..#\n........#....###...#..##..###.#\n.........#....#....#..#.#.#...#\n.#....###.##...#.#...........##\n..#..#.#..#.#.##..#...##.......\n##..#.#.#....#...#..#..........\n#..#.......#....#..##...####...\n............#.#..........##.##.\n#...#..#.#....#..#.#....##.....\n......#...#...#.##............#\n#.....##..###..#.#..#.#.##..#.#\n#..#.#..#......#.......##.#....\n##..#.#..#...#......#.##...###.\n.#....#..............#....#.#..\n..#.#..##....#....#..##........\n.#.#...#..#.....#.#..##........\n.....#..#.#......#....#.#..#.#.\n....#.###...###.#.#.....#......\n...........#.#....##....##.....\n..#..#.##..........#...#...#..#\n.....#.###.#..........#........\n....#....##........###...#.....\n.#.....##.......#....#..##..###\n#.....#...............##......#\n#..#.#..#.#.#.....#.#...#......\n.##.###...#....#..........##...\n.#.......#.....................\n.#.#....#...##..#...#...##.....\n.#.#...#.......#.......#...#...\n....#.#..#.#..#...#....##......\n....##.....#.##....#.##..##..##\n..#............#...###.##..#...\n.#..#.........#.##....#....#..#\n.#..##..#..#........#.#.##.#.##\n.###.#...#...............#...#.\n...#.##.##.#......#...#....##.#\n#......##.......##...###....#.#\n#..##.....##......#.#.##....#.#\n...#.#....#.#.#...........##..#\n#.....##......##.#..........##.\n###....#.#...#.#..####.........\n.##.#.#...##..#.....#..#...#...\n#.....#.#......#..........#.#..\n..###.##.#...................#.\n#.............#..#........#.##.\n#.#.#.#..#.....##..##.#....#...\n...#...#..#...#..##..##........\n...##...##..#...##...........#.\n.####..#.#.#.##.#.......#......\n...#....#.......#......#.......\n.....#.#...#...#..##..#..#.....\n......#.....###.#..#..#.#..###.\n.#....#....#..#..##.....##...#.\n.#.............##.###.#...#.#..\n#..#..#......#.###............#\n##.#..##....#..........#.#.#...\n......#........#...#.......##..\n....#.#..#..........#.....#.#..\n...#..#...#.#...#........#.....\n.....##...#....#.........##.##.\n....#...#...#.##.##...#....#...\n.#..#.....##......#..#.#..#....\n........##...##.##......#.#.#.#\n.................#..#.....##.#.\n...#.....#...#.........#..#.#.#\n....##.#.....#........#...#..#.\n#...............#..#.....#...#.\n.....#..#....#...#.####.#.#....\n####.#..#.##...#....#...##.....\n#...##..#...####..#....#.#...#.\n..#.......#.##..##...#.#.......\n...........##.......#....#..#..\n#.##....#...#.....#....##......\n....##.#.......#..#...##.......\n...#.........##.#..#......#.###\n.#..#..#....#.#.##....###..###.\n....#.#........##........##....\n....########....#.#.#.###.#...#\n...#.###.###.##......##.......#\n.#...#.###.......#..........#..\n..#..##.........#............#.\n.......##.#...#...##...#...#..#\n#.##....#.#...#.....#..#.#.....\n..#........#..#.#.#.#....#.##..\n...#...#.#.........#...#.#..##.\n#....#......#.#...........#..##\n...#.#.#..#...##...#...#...#...\n###..........#.#..........#....\n..#....#.#.#.#............#.#..\n....#...#..###...#.#....#......\n#...........####......##.#.....\n..#..##.#...#.....#..#.......##\n#.....#..###.....#...##..##....\n##..###..##...........#.#...#..\n.....#......#..............#...\n#..#.##.###.......#.......#...#\n#........#....##......#.#......\n.#.#.#...#.......#........#.##.\n#..#..##.....#...#.#.#.#..###..\n.#.#....#..#..#.#....##.#.#....\n..#.#.........####.#...#.#.###.\n....##........##....#........#.\n................#..........#...\n..#...................###.##..#\n.........#..#..#.#...#....#.#.#\n......#.....###.....#.#..#...#.\n.#.#.....#..##............##...\n...##......##.#....#...........\n...##..##..###.#...##..........\n....###...#..#.#......#......#.\n....##..............#..#..#.#..\n####.......#...##.##..#.#......\n.#......#.....#....###..#....#.\n.#.......#...##...#..##.#......\n#.......#.......#.#....#.#.#..#\n........#..#..#............##.#\n#.#...#.#..##..#.......##..#...\n...#....#...#..........##..#...\n#.#...#.##....###......##....#.\n#..#...###........#..#....#..#.\n#....#....###....#..#.......#..\n....#.#........#.............#.\n.#.##........##...#...#...#...#\n#.....##.....#.......#.#.#.....\n.#.##..........##..#....#......\n.#..##.##.#...##....#.#....##..\n........#.#.##.#....#.#..#....#\n..#...........................#\n.#...........#....#....#.#..#..\n........##...........#...#...#.\n..#.....#..#......#..##.......#\n..#....###..###...#.#.#..#....#\n#..#.#...#......##......#......\n...........#...##..##....##....\n#.#......###..#.....#.......#.#\n#.....#....#....#.#...#...#....\n....#...#.......#....##.#..#...\n.####..##......##.#........#..#\n..###..#.#.....#...........##..\n..##.#.#..#....#..#..#.........\n..........#.#.#####...#........\n.###......##.#....#........#...\n.....#..#..#.#..#.........#....\n..#....#...#...#...##..........\n....#..##.#.........##.#..##...\n##.####..#...#.#...#.....#..###\n..#..#...#...#.....##....#..#.#\n#..##..#.....#....#.#.....##..#\n...#...........##.....#......#.\n......#...#.....#.#..###.......\n.........#.....###.##..#...#...\n.#...#.##...#..........#.#..##.\n......#.......##.....#.....##..\n........###..........#...#.....\n##.......###..###...##...#.....\n#.#.............#..#..#.#......\n..##........#.###.....#....##..\n......#...#......#....##......#\n..#.....#...##...#.......#..#..\n..#.###..##.##...#....#...##.#.\n........##...#..#.#..##.....#.#\n.......................#......#\n..##.###......#.#.............#\n....#...........###............\n##...##.....#.......##.......#.\n...#..##..##..#.#.###..#......#\n........#........#.#..#..#.....\n.#......#....##..........#...#.\n.##...........##....#..........\n.#..#....###.......#....#..##..\n.....###..........#....#.#.#...\n...#....###.#.#......#......#..\n#.#.##.#.....#..#........#...#.\n...#.##.........#..#.....#.....\n.##...##......##...###...#.....\n...#.....#.##..#...#..#........\n........#............#.#.#..##.\n###...#.....#...#..#........##.\n##...#..#.....#.#....#.#.#.....\n#..##.......#...#.#...##..#....\n#...#.##.....#.#..#.##......#.#\n..#......#.#.#.##.##..........#\n..#.##......#.#.#..##..........\n....#..#....#..#..............#\n..........###.....##..#........\n...#.....##.....#..#.#..#...##.\n.#..##.#..#....#.#......#.##...\n...#.....#..#.#...#..#.....#.#.\n#...#.#......##...#..#...#....#\n..#.......##...#..#.......#...#\n#.....#...........##.#.........\n.#......##.....####...#.......#\n........#..#.....#.......#..#..\n....#.#...##..##...#..#....#...\n#.#......#...#.#.###.....#.....\n..##...#.#........#.##....#.#.#\n.#....#......#.#...###.#.......\n.......#.#...##....#.#....#....\n.....##..##...#..#.#.....##..#.\n.##..#.#.#....##.#...#.....#...\n.#..#..##....#.##.......#...#..\n....#.##...#..##......#.....#..\n.#..#....##....#...............\n..##...#.....###...............\n..............#.#.##........#.#\n.#.#....#....#...#.#........#..\n.##...#...#.#....#....#.#.....#\n#..............#......#.####.#.\n......#...........#..#.....##..\n#.#..##.##.....#......#..#.#..#\n##.##..#.##.#.............#...#\n...#..#......#....#............\n........###.#.#..#...#.....#.##\n..#.......#.##.........#..#....\n...##.#........##...#.#.##..#..\n...#..#......#...#....#........\n...........#..#..#...##...#....\n...#.....#....#..####..##.....#\n.......#..#..#......#.........#\n#......#........###.....##....#\n..#..#..#.#.#....##...##......#\n#.#..#..###.#..#.....####......\n.#................#####....#...\n.#.........#...#.......#......#\n..#.......#######........#.....\n..#........#.....#..#...#..#..#\n.#..#.#..#....#.#..##...#..#.#.\n..#...........#.#...#.#.##.....\n...#.#.#....##.###....#...####.\n.....#..#.....#..#.#.........#.\n......##...#...###............#\n..#.#......###..####..#......#.\n###.##.#..#......##.#..##.....#\n....###...##............#.#....\n..#.....##...#...##....#...#...\n#.....#.....#...#...#.#..#.....\n####..........##.#.#..#.....##.\n...#..........#...#...##..##.#.\n..........#.........#.#..#..#..\n#....###.....#.#...#.......##.#\n#..#.##.....#..........#...#...\n...#.#.###.......##..#.....#...\n#...#.#..#.............#..#.#..\n#........#.................#..#\n..#.#....#.#..##.#...#..#....#.\n#...#..........#...###....#...#\n......#.............#....#....#\n#.#.......##.......#.#....##..#\n##...#....#.............#..#...\n........#...###.##.#..###.#...#\n...##...#..#..#...##..##......#\n..#.......##....#.#.##....#....\n.....#....#..#.#...##.#.#.....#"
	forestLine := strings.Split(input, "\n")
	slope1 := slopeSettings{
		x: 1,
		y: 1,
	}
	res1 := computeTreeCount(forestLine, slope1)
	fmt.Println(slope1, res1)
	slope2 := slopeSettings{
		x: 3,
		y: 1,
	}
	res2 := computeTreeCount(forestLine, slope2)
	fmt.Println(slope2, res2)
	slope3 := slopeSettings{
		x: 5,
		y: 1,
	}
	res3 := computeTreeCount(forestLine, slope3)
	fmt.Println(slope3, res3)
	slope4 := slopeSettings{
		x: 7,
		y: 1,
	}
	res4 := computeTreeCount(forestLine, slope4)
	fmt.Println(slope4, res4)
	slope5 := slopeSettings{
		x: 1,
		y: 2,
	}
	res5 := computeTreeCount(forestLine, slope5)
	fmt.Println(slope5, res5)

	fmt.Println(res1 * res2 * res3 * res4 * res5)

}

func computeTreeCount(forestLine []string, settings slopeSettings) int {
	rightIndex := 0
	treeCount := 0

	for i := 0; i < len(forestLine); i = i + settings.y {
		line := forestLine[i]
		if line[rightIndex] == "#"[0] {
			treeCount++
		}
		if rightIndex+settings.x >= len(line) {
			rightIndex = (rightIndex + settings.x) % len(line)
		} else {
			rightIndex += settings.x
		}
	}
	return treeCount
}
