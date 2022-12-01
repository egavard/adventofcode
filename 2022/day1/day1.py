caloriesByElve = list()
caloriesSumByElve = 0
with open("day1.txt") as file_content:
    for line in file_content:
        if len(line.strip()) == 0:
            caloriesByElve.append(caloriesSumByElve)
            caloriesSumByElve = 0
        else:
            caloriesSumByElve += int(line)

print(max(caloriesByElve))
caloriesPerElveSortedDesc = sorted(caloriesByElve, reverse=True)
print(caloriesPerElveSortedDesc[0] + caloriesPerElveSortedDesc[1] + caloriesPerElveSortedDesc[2])
