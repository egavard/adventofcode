package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var Input = "1472\n1757\n1404\n1663\n1365\n1974\n1649\n1489\n1795\n1821\n1858\n1941\n1943\n1634\n1485\n1838\n817\n1815\n1442\n639\n1182\n1632\n1587\n1918\n1040\n1441\n1784\n1725\n1951\n1285\n285\n1224\n1755\n1748\n1488\n1374\n1946\n1771\n1809\n1929\n1621\n1462\n2001\n1588\n1888\n1959\n1787\n1690\n1363\n1567\n1853\n1990\n1819\n1904\n1458\n1882\n1348\n1957\n1454\n1557\n1471\n332\n1805\n1826\n1745\n1154\n1423\n1852\n1751\n1194\n1430\n1849\n1962\n1577\n1470\n1509\n1673\n1883\n1479\n1487\n2007\n1555\n1504\n1570\n2004\n978\n1681\n1631\n1791\n1267\n1245\n1383\n1482\n1355\n1792\n1806\n1376\n1199\n1391\n1759\n1474\n1268\n1942\n1936\n1766\n1233\n1876\n1674\n1761\n1542\n1468\n1543\n1986\n2005\n1689\n1606\n1865\n1783\n1807\n1779\n1860\n1408\n1505\n1435\n1205\n1952\n1201\n1714\n1743\n1872\n1897\n1978\n1683\n1846\n858\n1528\n1629\n1510\n1446\n1869\n1347\n685\n1478\n1387\n687\n1964\n1968\n1429\n1460\n1777\n1417\n1768\n1672\n1767\n1400\n1914\n1715\n1425\n1700\n1756\n1835\n1926\n1889\n1568\n1393\n1960\n1540\n1810\n1401\n1685\n830\n1789\n1652\n1899\n796\n1483\n1261\n1398\n1727\n1566\n1812\n1937\n1993\n1286\n1992\n1611\n1825\n1868\n1870\n1746\n1361\n1418\n1820\n1598\n1911\n1428\n1734\n1833\n1436\n1560\n"
	var splitted = strings.Split(Input, "\n")
	var size = len(splitted)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				var operand1, err = strconv.Atoi(splitted[i])
				var operand2, err_ = strconv.Atoi(splitted[j])
				var operand3, err__ = strconv.Atoi(splitted[k])
				if err != nil || err_ != nil || err__ != nil {
					break
				}
				//fmt.Println("i", i, "=",operand1)
				//fmt.Println("j",j,"=",operand2)
				//fmt.Println(operand1+operand2)
				if operand1+operand2+operand3 == 2020 {
					println(fmt.Sprintf("%i + %i + %i = 2020", operand1, operand2, operand3))
					fmt.Println(fmt.Println(operand1 * operand2 * operand3))
				}

			}
		}
	}
}
