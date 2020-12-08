package main

import "fmt"
import "strings"

func main() {
	input := ".....#............#....#####.##\n.#.#....#......#....##.........\n......#.#.#.....###.#.#........\n......#...#.....#####....#..##.\n...#............##...###.##....\n#.....#...#....#......##....##.\n#...#.#....#..#..##.##...#.....\n.......#..........#..#..#.#....\n.#.....#.#.......#..#...#....#.\n#..#.##.#..................###.\n...#.#.##...##.###.....#..#...#\n..#.#...#............#.......#.\n#..#.#..#.#....#...#.#.....#..#\n#......##....#..#.#.#........#.\n....#..#.#.#.##............#..#\n....#..#..#...#.#.##......#...#\n##...#...........#.....###.#...\n..#...#.#...#.#.....#....##.##.\n....##...##.#....#.....#.##....\n#........##......#......#.#.#.#\n....#.#.#.........##......#....\n.#......#...#.....##..#....#..#\n....#..#.#.....#..........#..#.\n..##...#..##................#.#\n.....#....#.#..#......#........\n........#..#.#......#.#........\n.....#.#....##.###....#...#....\n...##.#.......#....###..#......\n............##.#..#...#........\n#..###..#.....#.####...........\n.......##.....#......#......#..\n#........##..#.....##.......#.#\n#.##...#...#...#......##..#.#.#\n......#....##.#.#...#...##....#\n#..#....##.#......#.......##...\n.#..........#..........#....#.#\n#.....##......##....#..........\n..#.#.....#.#...#........#.....\n...#........#..#..#.##..##.....\n......###.....#..#...#.###...##\n.##.##.......#.......###...#...\n#.#..#.#.#....#.....###..#...##\n......#.##..........#.......##.\n#..#.#.........#.....##...##...\n..#...#....#....###.#......#...\n.....#..#.######.....#..#.#....\n..#.#.....#.....##.#....##.#.##\n...#.#.#....#....##..#..#.#.##.\n...........#.#...#..#..####....\n.........#####.#.#.#...#.##.#..\n.......#...#......#.##.#.##....\n....#.....#.....###..........#.\n.#.###....##.#..#..........#...\n#...#.........##.....####....#.\n##....##...#..........#........\n...#.#.#.#....#..........#.....\n.......#....#......##.......#..\n.#.#..#.........#.#.##....#....\n..#.............#..##.#.##..###\n.#.##..............#..#..##..#.\n..##.#..#......#...##..#.##...#\n......#..#....#....#....##..#..\n...#...##.............#..###...\n...##....#.#.##........#.....##\n....#.#.......#..###..#....####\n#...#...##..#.####..#...##....#\n.......#..#.##..#...#.#........\n###.#......#..##..#......#.##..\n#....#............#.....#......\n..##...#..##......#..#....#....\n.#..##...#....#.#...#...#..#..#\n........#....###...#..##..###.#\n.........#....#....#..#.#.#...#\n.#....###.##...#.#...........##\n..#..#.#..#.#.##..#...##.......\n##..#.#.#....#...#..#..........\n#..#.......#....#..##...####...\n............#.#..........##.##.\n#...#..#.#....#..#.#....##.....\n......#...#...#.##............#\n#.....##..###..#.#..#.#.##..#.#\n#..#.#..#......#.......##.#....\n##..#.#..#...#......#.##...###.\n.#....#..............#....#.#..\n..#.#..##....#....#..##........\n.#.#...#..#.....#.#..##........\n.....#..#.#......#....#.#..#.#.\n....#.###...###.#.#.....#......\n...........#.#....##....##.....\n..#..#.##..........#...#...#..#\n.....#.###.#..........#........\n....#....##........###...#.....\n.#.....##.......#....#..##..###\n#.....#...............##......#\n#..#.#..#.#.#.....#.#...#......\n.##.###...#....#..........##...\n.#.......#.....................\n.#.#....#...##..#...#...##.....\n.#.#...#.......#.......#...#...\n....#.#..#.#..#...#....##......\n....##.....#.##....#.##..##..##\n..#............#...###.##..#...\n.#..#.........#.##....#....#..#\n.#..##..#..#........#.#.##.#.##\n.###.#...#...............#...#.\n...#.##.##.#......#...#....##.#\n#......##.......##...###....#.#\n#..##.....##......#.#.##....#.#\n...#.#....#.#.#...........##..#\n#.....##......##.#..........##.\n###....#.#...#.#..####.........\n.##.#.#...##..#.....#..#...#...\n#.....#.#......#..........#.#..\n..###.##.#...................#.\n#.............#..#........#.##.\n#.#.#.#..#.....##..##.#....#...\n...#...#..#...#..##..##........\n...##...##..#...##...........#.\n.####..#.#.#.##.#.......#......\n...#....#.......#......#.......\n.....#.#...#...#..##..#..#.....\n......#.....###.#..#..#.#..###.\n.#....#....#..#..##.....##...#.\n.#.............##.###.#...#.#..\n#..#..#......#.###............#\n##.#..##....#..........#.#.#...\n......#........#...#.......##..\n....#.#..#..........#.....#.#..\n...#..#...#.#...#........#.....\n.....##...#....#.........##.##.\n....#...#...#.##.##...#....#...\n.#..#.....##......#..#.#..#....\n........##...##.##......#.#.#.#\n.................#..#.....##.#.\n...#.....#...#.........#..#.#.#\n....##.#.....#........#...#..#.\n#...............#..#.....#...#.\n.....#..#....#...#.####.#.#....\n####.#..#.##...#....#...##.....\n#...##..#...####..#....#.#...#.\n..#.......#.##..##...#.#.......\n...........##.......#....#..#..\n#.##....#...#.....#....##......\n....##.#.......#..#...##.......\n...#.........##.#..#......#.###\n.#..#..#....#.#.##....###..###.\n....#.#........##........##....\n....########....#.#.#.###.#...#\n...#.###.###.##......##.......#\n.#...#.###.......#..........#..\n..#..##.........#............#.\n.......##.#...#...##...#...#..#\n#.##....#.#...#.....#..#.#.....\n..#........#..#.#.#.#....#.##..\n...#...#.#.........#...#.#..##.\n#....#......#.#...........#..##\n...#.#.#..#...##...#...#...#...\n###..........#.#..........#....\n..#....#.#.#.#............#.#..\n....#...#..###...#.#....#......\n#...........####......##.#.....\n..#..##.#...#.....#..#.......##\n#.....#..###.....#...##..##....\n##..###..##...........#.#...#..\n.....#......#..............#...\n#..#.##.###.......#.......#...#\n#........#....##......#.#......\n.#.#.#...#.......#........#.##.\n#..#..##.....#...#.#.#.#..###..\n.#.#....#..#..#.#....##.#.#....\n..#.#.........####.#...#.#.###.\n....##........##....#........#.\n................#..........#...\n..#...................###.##..#\n.........#..#..#.#...#....#.#.#\n......#.....###.....#.#..#...#.\n.#.#.....#..##............##...\n...##......##.#....#...........\n...##..##..###.#...##..........\n....###...#..#.#......#......#.\n....##..............#..#..#.#..\n####.......#...##.##..#.#......\n.#......#.....#....###..#....#.\n.#.......#...##...#..##.#......\n#.......#.......#.#....#.#.#..#\n........#..#..#............##.#\n#.#...#.#..##..#.......##..#...\n...#....#...#..........##..#...\n#.#...#.##....###......##....#.\n#..#...###........#..#....#..#.\n#....#....###....#..#.......#..\n....#.#........#.............#.\n.#.##........##...#...#...#...#\n#.....##.....#.......#.#.#.....\n.#.##..........##..#....#......\n.#..##.##.#...##....#.#....##..\n........#.#.##.#....#.#..#....#\n..#...........................#\n.#...........#....#....#.#..#..\n........##...........#...#...#.\n..#.....#..#......#..##.......#\n..#....###..###...#.#.#..#....#\n#..#.#...#......##......#......\n...........#...##..##....##....\n#.#......###..#.....#.......#.#\n#.....#....#....#.#...#...#....\n....#...#.......#....##.#..#...\n.####..##......##.#........#..#\n..###..#.#.....#...........##..\n..##.#.#..#....#..#..#.........\n..........#.#.#####...#........\n.###......##.#....#........#...\n.....#..#..#.#..#.........#....\n..#....#...#...#...##..........\n....#..##.#.........##.#..##...\n##.####..#...#.#...#.....#..###\n..#..#...#...#.....##....#..#.#\n#..##..#.....#....#.#.....##..#\n...#...........##.....#......#.\n......#...#.....#.#..###.......\n.........#.....###.##..#...#...\n.#...#.##...#..........#.#..##.\n......#.......##.....#.....##..\n........###..........#...#.....\n##.......###..###...##...#.....\n#.#.............#..#..#.#......\n..##........#.###.....#....##..\n......#...#......#....##......#\n..#.....#...##...#.......#..#..\n..#.###..##.##...#....#...##.#.\n........##...#..#.#..##.....#.#\n.......................#......#\n..##.###......#.#.............#\n....#...........###............\n##...##.....#.......##.......#.\n...#..##..##..#.#.###..#......#\n........#........#.#..#..#.....\n.#......#....##..........#...#.\n.##...........##....#..........\n.#..#....###.......#....#..##..\n.....###..........#....#.#.#...\n...#....###.#.#......#......#..\n#.#.##.#.....#..#........#...#.\n...#.##.........#..#.....#.....\n.##...##......##...###...#.....\n...#.....#.##..#...#..#........\n........#............#.#.#..##.\n###...#.....#...#..#........##.\n##...#..#.....#.#....#.#.#.....\n#..##.......#...#.#...##..#....\n#...#.##.....#.#..#.##......#.#\n..#......#.#.#.##.##..........#\n..#.##......#.#.#..##..........\n....#..#....#..#..............#\n..........###.....##..#........\n...#.....##.....#..#.#..#...##.\n.#..##.#..#....#.#......#.##...\n...#.....#..#.#...#..#.....#.#.\n#...#.#......##...#..#...#....#\n..#.......##...#..#.......#...#\n#.....#...........##.#.........\n.#......##.....####...#.......#\n........#..#.....#.......#..#..\n....#.#...##..##...#..#....#...\n#.#......#...#.#.###.....#.....\n..##...#.#........#.##....#.#.#\n.#....#......#.#...###.#.......\n.......#.#...##....#.#....#....\n.....##..##...#..#.#.....##..#.\n.##..#.#.#....##.#...#.....#...\n.#..#..##....#.##.......#...#..\n....#.##...#..##......#.....#..\n.#..#....##....#...............\n..##...#.....###...............\n..............#.#.##........#.#\n.#.#....#....#...#.#........#..\n.##...#...#.#....#....#.#.....#\n#..............#......#.####.#.\n......#...........#..#.....##..\n#.#..##.##.....#......#..#.#..#\n##.##..#.##.#.............#...#\n...#..#......#....#............\n........###.#.#..#...#.....#.##\n..#.......#.##.........#..#....\n...##.#........##...#.#.##..#..\n...#..#......#...#....#........\n...........#..#..#...##...#....\n...#.....#....#..####..##.....#\n.......#..#..#......#.........#\n#......#........###.....##....#\n..#..#..#.#.#....##...##......#\n#.#..#..###.#..#.....####......\n.#................#####....#...\n.#.........#...#.......#......#\n..#.......#######........#.....\n..#........#.....#..#...#..#..#\n.#..#.#..#....#.#..##...#..#.#.\n..#...........#.#...#.#.##.....\n...#.#.#....##.###....#...####.\n.....#..#.....#..#.#.........#.\n......##...#...###............#\n..#.#......###..####..#......#.\n###.##.#..#......##.#..##.....#\n....###...##............#.#....\n..#.....##...#...##....#...#...\n#.....#.....#...#...#.#..#.....\n####..........##.#.#..#.....##.\n...#..........#...#...##..##.#.\n..........#.........#.#..#..#..\n#....###.....#.#...#.......##.#\n#..#.##.....#..........#...#...\n...#.#.###.......##..#.....#...\n#...#.#..#.............#..#.#..\n#........#.................#..#\n..#.#....#.#..##.#...#..#....#.\n#...#..........#...###....#...#\n......#.............#....#....#\n#.#.......##.......#.#....##..#\n##...#....#.............#..#...\n........#...###.##.#..###.#...#\n...##...#..#..#...##..##......#\n..#.......##....#.#.##....#....\n.....#....#..#.#...##.#.#.....#"
	forestLine := strings.Split(input, "\n")
	rightIndex := 0
	treeCount := 0
	for _, line := range forestLine {
		if line[rightIndex] == "#"[0] {
			treeCount++
		}
		if rightIndex+3 >= len(line) {
			rightIndex = (rightIndex + 3) % len(line)
		} else {
			rightIndex += 3
		}
	}
	fmt.Println(treeCount)
}
