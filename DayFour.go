package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type Passport struct {
	byr int
	iyr int
	eyr int
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func main() {
	input := "iyr:1928 cid:150 pid:476113241 eyr:2039 hcl:a5ac0f\necl:#25f8d2\nbyr:2027 hgt:190\n\nhgt:168cm eyr:2026 ecl:hzl hcl:#fffffd cid:169 pid:920076943\nbyr:1929 iyr:2013\n\nhgt:156cm ecl:brn eyr:2023\niyr:2011\nhcl:#6b5442 pid:328412891 byr:1948\n\nbyr:1950 iyr:2019 eyr:2020 ecl:amb cid:279 pid:674907993 hgt:189cm hcl:#602927\n\nbyr:1976\necl:hzl iyr:2015 hgt:178cm eyr:2022 hcl:#341e13\npid:473630095\n\niyr:2020 eyr:2023 ecl:blu byr:1984\nhgt:163cm hcl:#866857 pid:628113926\n\necl:amb\npid:312508073\nhgt:70in byr:1922 iyr:2019 eyr:2030 hcl:#866857\n\nhcl:#007d7c pid:195125455 cid:213 hgt:154cm eyr:2021 ecl:grn byr:1981\n\necl:oth hgt:185cm pid:958027833 hcl:#b6652a iyr:2028 cid:171\neyr:1994\n\necl:hzl byr:1982 hcl:#fffffd hgt:188cm iyr:2018 pid:039931872 cid:288 eyr:2025\n\ncid:71 iyr:2012 byr:1950 hcl:#7d3b0c pid:803324747 eyr:2023 hgt:151cm ecl:oth\n\niyr:2013\necl:grn eyr:2022\npid:053411982 byr:1946 cid:302 hcl:#60ca85\nhgt:160cm\n\nhgt:169cm eyr:2035 pid:023983645 iyr:2014 ecl:amb hcl:#c0946f byr:1975 cid:258\n\nbyr:1933 ecl:hzl\nhcl:#c0946f iyr:2013 pid:655452550\nhgt:170cm\neyr:2024\n\nhgt:156\necl:oth\ncid:235\npid:609823906 iyr:2016 eyr:2021 hcl:#6b5442\nbyr:1932\n\niyr:2006\nhgt:103 ecl:#2d77e5 cid:214 byr:2018 hcl:6c53a4 pid:190cm eyr:1940\n\necl:grn\npid:497830156 byr:2002 eyr:2023 hgt:169cm iyr:2016 hcl:#733820\n\necl:gmt hgt:75cm byr:2007 eyr:2037 iyr:2028 hcl:4591f6 cid:118\n\ncid:94\necl:hzl byr:1972 hcl:#7d3b0c iyr:2015 pid:219956257\neyr:2022 hgt:165cm\n\neyr:2022 hgt:180cm ecl:amb hcl:#c0946f\npid:543330083\niyr:2014\ncid:286 byr:1989\n\necl:hzl eyr:2027 iyr:2019 pid:125201586\nbyr:1947 cid:74 hcl:#341e13\n\niyr:2020 hgt:192cm ecl:oth\npid:651509606 byr:1965 eyr:2029\nhcl:#b6652a\n\nhgt:165cm eyr:2025 ecl:oth pid:844167324 byr:1950 iyr:2014 hcl:#a97842\n\nhgt:159cm\nbyr:1945 hcl:#6b5442 iyr:2027\neyr:2024\ncid:94 ecl:brn pid:476551927\n\npid:479260033 hcl:#efcc98 iyr:2018 ecl:grn\nbyr:1993 cid:92 hgt:165cm\neyr:2027\n\niyr:2015 pid:106083602\nhgt:168cm eyr:2025 ecl:gry byr:1996 cid:341\nhcl:#fffffd\n\niyr:2010 hgt:192cm\npid:247508683 ecl:#57a15d byr:1972\nhcl:#602927 eyr:2024\n\necl:blu byr:1934 hcl:#888785 iyr:2019 pid:905361316 eyr:2021 hgt:150cm\n\nhgt:184cm hcl:#cfa07d cid:335 iyr:2018 byr:1995\necl:grn eyr:2026 pid:435090537\n\npid:302395756\necl:grn hcl:z byr:2005 hgt:111 eyr:2031 cid:147\n\necl:gry pid:561021264 cid:178 byr:1980 iyr:2010\neyr:2028 hcl:#7d3b0c hgt:181cm\n\nhgt:172cm byr:1923 pid:741415636 ecl:grn eyr:2022 iyr:2013\n\npid:457776708\nbyr:1992\nhcl:#b6652a hgt:157cm eyr:2024 iyr:2011\n\npid:177860177\necl:blu\nhgt:154cm hcl:#cfa07d iyr:2015 eyr:2022\nbyr:1977\n\npid:992814815 eyr:2028 iyr:2017 hgt:181cm hcl:#cfa07d\nbyr:1961 ecl:hzl\n\neyr:2025 hcl:#a97842\nbyr:1930 pid:468404395\niyr:2013 ecl:oth cid:220 hgt:170cm\n\ncid:198\niyr:2018 hcl:#a97842 hgt:74in\npid:279483949 eyr:2029 ecl:gry byr:1931\n\nbyr:2004 iyr:2021 pid:165cm ecl:#7e7d04\nhcl:#18171d\neyr:2035 hgt:61\n\necl:#492a33\nhgt:168cm\niyr:2018\nbyr:2017 cid:293\npid:1764204298 hcl:#cfa07d eyr:2022\n\nhcl:#866857\neyr:2026\ncid:193 hgt:160cm byr:1970 iyr:2011 ecl:amb pid:895650554\n\neyr:2022\niyr:2018\nhcl:#efcc98 cid:181\nbyr:2029 ecl:utc hgt:188cm pid:332630362\n\nhcl:#ceb3a1\niyr:2013 pid:592603167\ncid:95 ecl:blu eyr:2022\n\nhcl:#efcc98\niyr:2011 pid:550968343\necl:hzl byr:1924 eyr:2022\nhgt:191cm cid:120\n\nhgt:150cm ecl:grn\nhcl:8f3824 pid:735766540 eyr:2029\nbyr:2000 iyr:2015\n\nhcl:z\necl:hzl byr:2003 hgt:118 eyr:2008 iyr:2022\npid:157cm\n\nbyr:1950 ecl:blu hgt:163cm\npid:455597862 cid:302 eyr:2027\nhcl:#341e13 iyr:2015\n\niyr:2015 ecl:oth eyr:2023 byr:1998 hcl:#ceb3a1 cid:136 pid:253146183\nhgt:179cm\n\niyr:2018 hcl:#cfa07d cid:80\npid:347839572 byr:1946 eyr:2023 ecl:blu\nhgt:163cm\n\niyr:1969 cid:324 eyr:1927 ecl:lzr\nhcl:z\nbyr:2030 hgt:172cm\npid:#997235\n\niyr:2017 ecl:brn\nhgt:165cm\npid:818623102 byr:1968 hcl:#fffffd eyr:2020\n\neyr:2023 byr:1966 ecl:blu\ncid:295 pid:347753668\nhcl:#c0946f\niyr:2010 hgt:163cm\n\nhcl:#ceb3a1 pid:395843182 hgt:168cm eyr:2025 iyr:2014 byr:1991 ecl:gry cid:283\n\niyr:2011 byr:1928 pid:438089427\nhgt:152cm\necl:hzl eyr:2022 cid:254 hcl:#866857\n\niyr:2015\nhcl:#ceb3a1\necl:lzr eyr:2022 hgt:173cm pid:1799325911 cid:210 byr:2018\n\niyr:2010\npid:121142355\neyr:2020\ncid:302\nhgt:158cm ecl:amb\nbyr:1978 hcl:#623a2f\n\npid:110863702\nhcl:#341e13 iyr:2017 byr:1942 hgt:175cm cid:277 eyr:2030\necl:amb\n\nhcl:#c0946f\npid:473360783 byr:1986\nhgt:159cm ecl:brn\niyr:2011 eyr:2023\n\niyr:2015 hcl:#733820 pid:245692263\necl:oth byr:1960 eyr:2022\n\nhcl:b9c0fd iyr:1996 hgt:83\nbyr:2029 pid:#449a30\necl:grt eyr:1925\n\nhgt:68cm\neyr:2039 hcl:#cfa07d\npid:193cm iyr:1984\necl:#b9ec76\n\neyr:2023 ecl:amb\nbyr:1942\niyr:2012 hcl:#b6652a hgt:156cm pid:398126488\n\necl:oth hgt:150cm byr:1937 pid:927692980 iyr:2013 eyr:2023 hcl:#623a2f\n\neyr:2026 byr:1921 pid:297672804 hgt:172cm iyr:2011 ecl:brn\n\neyr:2026 cid:241\nhcl:#341e13\npid:316611397 hgt:193cm\nbyr:1977\n\npid:509492550 hgt:64cm eyr:2030 hcl:#b6652a byr:1986 iyr:1922 ecl:gry\n\nhgt:165cm cid:248 hcl:#6b5442 eyr:2026\npid:703744314\nbyr:1921 iyr:2020\necl:blu\n\nbyr:2001 pid:332016728\niyr:2018 cid:89\neyr:2031 hgt:155cm ecl:zzz\nhcl:#866857\n\nbyr:2023\nhcl:z pid:3586415546 iyr:2022 cid:209 hgt:188in ecl:brn\n\necl:grn\nhgt:61in iyr:1925 byr:1984 hcl:#733820\npid:216995428 eyr:1944\n\nbyr:1969 hcl:#a97842 cid:226\niyr:2011 pid:621770561\neyr:2024 ecl:blu\n\nhcl:#efcc98 eyr:2024\niyr:2010 ecl:hzl\npid:153620883 byr:1957\n\niyr:2015\ncid:162 eyr:2020\npid:89806820 byr:1955\nhcl:b043dd ecl:brn\n\nhgt:162cm\nhcl:2ee8db\nbyr:2008 iyr:2003 pid:50279629 eyr:2030 ecl:grt\n\npid:939011546 byr:1945\nhgt:70in hcl:#cfa07d eyr:2027 ecl:grn iyr:2015\n\nhgt:83 ecl:hzl hcl:z eyr:2026 byr:2029\n\ncid:244 hcl:#623a2f iyr:2012 pid:527925497\nbyr:1957\neyr:2024 ecl:brn\n\nhgt:179cm\nbyr:1928\npid:933893768 hcl:#18171d ecl:gry iyr:2016 eyr:2027\n\nhgt:158cm iyr:2017 ecl:brn byr:1935 eyr:2020\npid:331047535 cid:345 hcl:#888785\n\nbyr:2009\necl:#893922\niyr:2020 hcl:a59633 hgt:170in eyr:1995\npid:28540793\n\nbyr:1955 hgt:68cm\nhcl:#67dac3 eyr:2031 pid:502641687 ecl:oth iyr:1922\n\npid:2523045951 cid:203 hgt:75cm eyr:2031 hcl:#888785\niyr:1937 byr:1988\n\npid:558076850 eyr:2030\nhgt:192cm ecl:brn\ncid:296 byr:1954\nhcl:#733820 iyr:2012\n\ncid:272 eyr:2030 pid:044961585\nhcl:#602927 byr:1990 hgt:173cm ecl:gry iyr:2018\n\nbyr:1958 iyr:2019 hgt:163cm eyr:2029\npid:384542472 hcl:819959\necl:#866be8\n\niyr:2027\npid:7267919678 byr:2013 hgt:161in hcl:z ecl:brn\n\npid:855195796 ecl:oth\neyr:2030 hgt:163cm hcl:#341e13 byr:1978\niyr:2011 cid:206\n\necl:brn eyr:2029 hcl:#fffffd iyr:2018 pid:065149883 byr:1938 hgt:178cm\n\neyr:2024\nbyr:1983\necl:gry\nhgt:154cm\niyr:2019\npid:#f331f5 hcl:#7d3b0c cid:315\n\necl:brn pid:131551626 iyr:2013 eyr:2022 byr:1949\nhgt:155cm hcl:#18171d\n\ncid:203 eyr:2028 iyr:2019\nbyr:1939\nhcl:#18171d pid:091534428 hgt:175cm\n\nbyr:1921 eyr:2025 iyr:2014 pid:719127279 ecl:brn hcl:#cfa07d cid:243 hgt:176cm\n\nbyr:1976 hgt:182cm\necl:gry pid:534666141\niyr:2019 eyr:2027 cid:197 hcl:#602927\n\nbyr:2015\npid:164cm hgt:90 eyr:2036 iyr:1947 hcl:b7b0e6 ecl:#fd96b3\n\neyr:2029 cid:264 pid:931433692\nbyr:1974 ecl:oth hcl:z hgt:67in iyr:2012\n\npid:179cm ecl:#00a56d\neyr:2025 hcl:eed83e iyr:1949 hgt:177in\n\nhgt:159cm ecl:blu\npid:5642951907 iyr:2029 byr:1952\nhcl:#6b5442\n\necl:amb hgt:163cm\npid:811866600 byr:1952\niyr:2019 hcl:#888785\ncid:250 eyr:2027\n\nbyr:1953 hgt:190cm\npid:156cm hcl:#7d3b0c eyr:2022 ecl:#1b0b35 iyr:2015\n\npid:709465009 byr:1971 iyr:2018 hcl:#602927 ecl:oth\ncid:222 eyr:2025\n\nhcl:#623a2f pid:583448566\nbyr:1999\neyr:2026 hgt:179cm\niyr:2015 ecl:gry cid:55\n\nhgt:179cm iyr:2013 ecl:amb hcl:#95766f pid:620956072\nbyr:1997 eyr:2026\n\necl:blu iyr:1924 pid:866797032 hgt:193cm cid:90 hcl:#fffffd eyr:1998 byr:1990\n\nhcl:#733820 ecl:brn byr:1950 eyr:2028\nhgt:155cm iyr:2017\npid:605542221\n\nhgt:171cm iyr:2019 byr:1930\necl:hzl\neyr:2026 hcl:#a6ef22 pid:294449839\n\npid:480248391\nhgt:150cm eyr:2027 cid:226 hcl:#cfa07d\nbyr:1940 ecl:brn\niyr:2018\n\nhcl:z ecl:#769ca0 pid:180cm\nbyr:1922 iyr:2026 eyr:2028\nhgt:180cm\n\necl:lzr byr:1967 pid:50313895 hcl:34441e hgt:158in iyr:2030 eyr:2039\n\niyr:2025\npid:2210753 byr:2010 hgt:173cm cid:208\neyr:2008 hcl:de66d6\necl:grt\n\niyr:2018 eyr:2026\ncid:289 byr:1992\nhgt:170cm pid:856187601 ecl:gry hcl:#efcc98\n\ncid:94 byr:1934 hgt:59in eyr:2022\nhcl:#623a2f pid:573884719\niyr:2016 ecl:oth\n\npid:206185815 ecl:grn hcl:#cfa07d eyr:2027\niyr:2018 byr:1989\nhgt:176cm\n\nhgt:175cm byr:1999\npid:409477026\nhcl:#cfa07d\necl:amb eyr:2021 iyr:2017 cid:75\n\nbyr:2018\ncid:150 eyr:2033 pid:043853978 iyr:2017 hgt:61cm hcl:z\necl:#f19d87\n\npid:549507973 hgt:178cm byr:1929 ecl:oth\niyr:2020 eyr:2025\nhcl:#7d3b0c\n\niyr:2014 hgt:171cm ecl:blu byr:1999\nhcl:#6b5442 pid:813505466\neyr:2029\n\necl:zzz eyr:2034\nbyr:2022\npid:52407584 iyr:2016 hcl:#888785\nhgt:176in\n\necl:oth\nbyr:1994 iyr:2018 hgt:64in pid:136896463\neyr:2022\nhcl:#a97842\n\necl:#535e3c hgt:84\neyr:1963 hcl:z\niyr:1986 pid:187cm byr:2028 cid:258\n\neyr:2029\ncid:257 hgt:175cm\necl:oth iyr:2016\nhcl:#602927 pid:506432649\n\niyr:2015 hgt:165cm\necl:gmt cid:116 hcl:z\nbyr:1998\neyr:2021\npid:170cm\n\niyr:2023 hgt:178cm cid:109 pid:#6eca6e hcl:#7d3b0c eyr:1961\necl:xry byr:2012\n\neyr:2025\necl:grn\npid:708755870 hgt:189cm hcl:#e23d5f\niyr:2017 byr:1982\n\nhcl:#866857 pid:85618849 ecl:brn byr:1958 eyr:2025\nhgt:111\ncid:190\n\nhgt:75cm byr:1983 iyr:2000\neyr:2007\ncid:307\npid:227345093 ecl:#080923 hcl:#ceb3a1\n\nhcl:#602927\necl:oth hgt:158cm byr:1992 iyr:2012 pid:708206240\neyr:2026 cid:125\n\neyr:1920 pid:873476029 hgt:192cm byr:1971 ecl:gry iyr:2020 hcl:#f463f6\n\npid:295847270 hcl:#7d3b0c ecl:oth iyr:2015\nbyr:2000 hgt:181cm eyr:2025\n\nhgt:189cm\nhcl:#18171d iyr:2013\npid:686835652 byr:1972\necl:grn eyr:2029\n\niyr:2010\necl:grn hgt:63cm eyr:2027 hcl:#602927 pid:240973955 byr:1984 cid:280\n\npid:883408516 eyr:2022\niyr:2010 hgt:182cm ecl:hzl byr:2000 cid:220\n\niyr:2018 pid:026680847 cid:117 hcl:#602927 hgt:67cm ecl:xry eyr:2030\nbyr:1989\n\nbyr:1933 ecl:hzl\nhgt:179cm\npid:500053352 eyr:2020 hcl:#fffffd\niyr:2014\n\nhgt:153cm\npid:523083973 ecl:brn\niyr:2011 byr:2000 hcl:#cfa07d\neyr:2020 cid:114\n\nhcl:#efcc98 ecl:blu\nbyr:1974 iyr:2019\nhgt:165cm\neyr:2020 pid:755433303\n\neyr:2022\necl:amb byr:1927 iyr:2012 pid:409960222 hcl:#733820 hgt:169cm cid:336\n\necl:#564a01\nhgt:136 iyr:1984\npid:#646419\neyr:2032\nhcl:z\n\nhgt:71in hcl:14d37b\nbyr:2017 cid:243 ecl:zzz pid:208245975\niyr:2029\n\nbyr:1974 hcl:#6b5442 pid:562222331 hgt:68in\ncid:319\necl:grn\niyr:2012 eyr:2028\n\niyr:2010 byr:1948 hgt:169cm eyr:2022 hcl:#623a2f\ncid:93 ecl:hzl\n\ncid:347\nbyr:1939 hgt:151cm eyr:2026\niyr:2010\nhcl:#fffffd ecl:gry\npid:562919031\n\nhgt:171cm\niyr:2010 pid:812511153 byr:1971 eyr:2026 ecl:hzl\nhcl:#6b5442\n\ncid:319 eyr:2026 iyr:2013\nhgt:155in\nhcl:z pid:185cm\n\nhgt:178cm ecl:gry cid:139 hcl:#341e13 pid:390510619 eyr:2026 iyr:2012\nbyr:1952\n\neyr:2025 pid:78761845\nhcl:#866857 iyr:2019\nhgt:173cm ecl:blu byr:1936\n\neyr:2028 hgt:192cm\nbyr:1946 pid:897533472 ecl:brn hcl:#efcc98\n\npid:467427172 hcl:#efcc98\neyr:2021 byr:1923\niyr:2012 cid:139 hgt:176cm\n\niyr:2015 eyr:2028\npid:069618718\nhgt:190cm ecl:grn hcl:#888785\nbyr:1956 cid:68\n\necl:brn hgt:173cm eyr:2022\niyr:2010 pid:525711593 byr:1990\n\ncid:292\necl:blu hcl:#602927 hgt:67in iyr:2011 byr:1990 eyr:2027 pid:298224903\n\nhgt:159cm eyr:2029 pid:854089988 iyr:2018 ecl:gry byr:1962 hcl:#efcc98\n\necl:grn byr:1964 eyr:2022\nhgt:61in pid:202756433 hcl:#cfa07d cid:241\niyr:2015\n\nhgt:68in byr:1973 hcl:#18171d ecl:hzl\npid:701847555 eyr:2030 iyr:2019\n\neyr:2022\necl:grn hgt:151cm iyr:2020 hcl:#83f878 byr:1982 pid:816902510\n\ncid:130 hgt:187in eyr:2040\necl:brn\niyr:2020\nhcl:z pid:7364218001\nbyr:1949\n\nhgt:183cm\neyr:2023 iyr:2019 byr:1946 pid:684966686\ncid:307 ecl:brn hcl:#cfa07d\n\nhcl:#6b5442 eyr:2024 pid:7727182081\niyr:2017\nhgt:110 ecl:dne\n\necl:blu byr:1987 cid:167 iyr:2015 hgt:189cm\npid:797675433 eyr:2024 hcl:#6b5442\n\niyr:2018 byr:1929 ecl:brn hgt:60in eyr:2024 pid:152cm hcl:#a97842\n\niyr:2020 eyr:2025 byr:1942 pid:007017276 ecl:oth hgt:170cm\nhcl:#ceb3a1 cid:104\n\niyr:2012 ecl:oth eyr:2020\nbyr:1965 hcl:#efcc98\nhgt:173cm\ncid:102 pid:302599543\n\nhgt:187cm pid:958933966\necl:hzl byr:1955\neyr:2027 hcl:#6b5442\n\necl:oth iyr:2013\neyr:2027 hgt:153cm cid:86 hcl:#602927\npid:568040159 byr:1926\n\nhgt:187cm iyr:2008 pid:151cm ecl:blu eyr:1954\n\nbyr:2014\npid:9029821667 hgt:59cm eyr:2035 hcl:e9c79a\niyr:2010\n\neyr:2027 pid:#d676d0\nhcl:d2fcfa hgt:154cm ecl:hzl byr:1938\n\necl:lzr hgt:61in eyr:2025\npid:556812665\nbyr:1923 iyr:2019\nhcl:e962ed\n\niyr:2019\neyr:2029\nhcl:#866857 byr:1977 pid:115229656 hgt:193cm\necl:brn cid:350\n\nhcl:z pid:#8d311d iyr:2023 hgt:71cm\nbyr:1923 ecl:zzz eyr:2039\n\ncid:66 hgt:165cm\neyr:2027 iyr:2012 hcl:#b6652a ecl:amb pid:946987379 byr:1999\n\nbyr:2028 iyr:2013 ecl:#376cda\neyr:1928 pid:#c135ce hcl:z hgt:185in\n\nhcl:100344 iyr:1933 eyr:2023 hgt:71cm byr:2010 ecl:#6a8007 pid:90001213\n\niyr:2012\nbyr:1987 eyr:2020 hgt:190cm cid:298 hcl:#866857\n\nhgt:161cm hcl:#efcc98 ecl:grn eyr:2028 iyr:2014\nbyr:1966 pid:769989459\n\nhgt:173cm pid:527615519 eyr:2024 hcl:#602927 byr:1949 ecl:oth cid:328\n\npid:679489285\nhgt:153cm byr:1963\nhcl:#602927 eyr:2026 ecl:blu\n\necl:blu hgt:186cm hcl:#c0946f pid:741255292 eyr:2022 byr:1996 iyr:2017\n\nhgt:172cm\nhcl:#888785 eyr:2022 pid:377797887 byr:1980\n\nhcl:z pid:399837694 iyr:2018 ecl:#33e59d eyr:2038\nhgt:60in\n\neyr:2027\nbyr:1923\nhgt:170cm pid:754104917\niyr:2020 cid:135 hcl:#341e13\necl:brn\n\necl:grn hcl:#c0946f\nbyr:2028 iyr:2016 pid:950191991\nhgt:193cm cid:93\neyr:1935\n\necl:brn hcl:#733820 eyr:2024\niyr:2017 pid:450063924\nbyr:2000 hgt:172cm\n\niyr:2008\ncid:229 byr:2023 eyr:2022 hcl:#341e13\necl:grn\nhgt:70in pid:104660281\n\neyr:2023 hgt:181cm cid:289 pid:828542447\niyr:2013 ecl:grn byr:1922 hcl:#866857\n\niyr:2030 pid:152cm cid:297 ecl:#75a512 hcl:z hgt:156in byr:2006\neyr:2035\n\niyr:2012 hcl:#18171d eyr:2025 hgt:188cm\necl:blu byr:1976\n\niyr:2018 hgt:157cm hcl:#b6652a\necl:oth byr:2002 eyr:2023\n\ncid:161\nhcl:#b6652a iyr:2016\nbyr:1930 ecl:oth pid:000425745 hgt:167cm eyr:2022\n\nhgt:160cm hcl:#89f1a0 eyr:2023 pid:867868252 byr:1976 iyr:2019 ecl:hzl\n\nbyr:1966 ecl:grn pid:597443937\niyr:2014 eyr:2029\n\npid:306301971 ecl:#a145cc\nhcl:z iyr:2018 cid:325 eyr:2023 byr:1942 hgt:157cm\n\necl:brn\npid:771134604 hgt:160cm\nbyr:1961 eyr:2020\niyr:2012 hcl:#6b5442\n\niyr:1922\necl:gmt\neyr:1963\npid:#d1a6f3 hcl:z byr:2015 hgt:153in\n\neyr:2022 ecl:gry\nhgt:156cm\npid:640711969\nhcl:#cfa07d\n\necl:grn\neyr:1980 pid:385212564 hcl:5b27f7 hgt:160cm iyr:2016 cid:171 byr:1990\n\niyr:2020\ncid:212 pid:959667791 byr:2002 ecl:amb\nhgt:75in eyr:2026 hcl:#888785\n\nbyr:1969 eyr:2021\niyr:2012\npid:318752605 hgt:179cm\ncid:81 hcl:#888785\n\nbyr:1926 hcl:#c0946f iyr:2010 hgt:155cm ecl:gry pid:475722917\neyr:2030\n\neyr:2025\necl:grn byr:1980 iyr:2010 hgt:160cm hcl:#d03ef0 pid:474973131\n\neyr:2020 iyr:2012 hgt:150cm\nhcl:#c0946f\nbyr:1924 ecl:amb\n\niyr:2016 hgt:173cm eyr:2029\nhcl:#888785 ecl:hzl byr:2001 cid:334 pid:291454183\n\niyr:2013\npid:909258239 byr:1970\necl:utc eyr:2026\ncid:312 hgt:158cm\nhcl:#18171d\n\necl:grn\nbyr:1941 pid:395943714\neyr:2027\nhcl:#7d3b0c\niyr:2011 hgt:158cm\n\necl:amb hcl:#fffffd\nbyr:1992\npid:266072435\neyr:2028 iyr:2020 hgt:161cm\n\nhcl:de3776 eyr:2021\ncid:234 ecl:#160982\niyr:2017 byr:1992\n\nbyr:1979 iyr:2020 ecl:brn\nhcl:#6b5442\npid:492860333 hgt:168cm eyr:2030\n\neyr:2025 hcl:#fffffd pid:776551474\necl:hzl hgt:169cm\niyr:2017\n\necl:hzl\neyr:2029\niyr:2013 byr:1952 hgt:152cm\npid:968064648 hcl:#6b5442\n\nbyr:1955\npid:947711080\ncid:149 ecl:amb\nhgt:150cm\nhcl:#341e13 eyr:2022 iyr:2016\n\nhgt:71cm ecl:#c6c47f\nbyr:2028 iyr:1994 eyr:2030 pid:0684877002 cid:237 hcl:#341e13\n\neyr:2030 hcl:#a97842 hgt:188cm byr:2000 pid:262013450\niyr:2018\n\nhgt:74in byr:1955 ecl:blu iyr:2012 hcl:#341e13 pid:165688658\n\nhgt:176cm cid:346 iyr:2012\npid:322396589\necl:gry eyr:2029\nbyr:1976\nhcl:#888785\n\neyr:2021\niyr:2015 hcl:3a6401 byr:1997 ecl:blu pid:188cm hgt:166in\n\necl:blu iyr:2010 byr:1984 hgt:183\npid:306571244 hcl:#623a2f eyr:2033 cid:113\n\necl:#804adb byr:2004 hgt:181cm\nhcl:#623a2f\neyr:2040 pid:#57e9d1\niyr:2028 cid:97\n\niyr:2015 pid:294753454 byr:1980 eyr:2020\nhgt:76in\necl:oth\nhcl:#a97842\n\nhcl:#a7a05c pid:0137262572 eyr:2023 cid:350 iyr:2015\necl:#52d3fe hgt:190cm\nbyr:2007\n\npid:826827136 eyr:2030 ecl:brn byr:1946 hcl:#a97842 iyr:2018\nhgt:173in\n\nbyr:1967\niyr:2015 pid:142177822 hgt:157cm ecl:oth eyr:2024 cid:221\n\niyr:2012 byr:1942 cid:187 pid:886132093\nhgt:158cm ecl:hzl hcl:#1bc909\n\npid:490847399\nbyr:1963\nhgt:69in\niyr:2011 ecl:gry\neyr:2027 hcl:#e4f497 cid:87\n\niyr:2014 ecl:hzl hgt:159cm hcl:#c0946f eyr:2028 byr:1926 pid:007819051\n\nhcl:#cfa07d pid:639664506 ecl:amb\nbyr:1997 cid:137 iyr:2014 eyr:2030 hgt:67in\n\nhgt:191in\neyr:2025\ncid:128\nbyr:2021 iyr:2015\nhcl:5ed1ae ecl:lzr\npid:406311551\n\neyr:2035\necl:gmt hcl:71e1ef iyr:2023\npid:4347854 byr:2017\n\nhgt:169cm\neyr:2028\necl:oth iyr:2016 byr:1954\npid:662755630 hcl:#733820\n\neyr:2029 pid:664032828 hgt:185cm hcl:#fffffd byr:1991 ecl:grn iyr:2017\n\npid:240747543 hgt:190cm\nhcl:#18171d iyr:2013 eyr:2021 ecl:grn byr:1920\n\niyr:2024 pid:87644548 hgt:126\nbyr:1971 ecl:brn\neyr:2040\n\niyr:2020\necl:lzr byr:2014 eyr:2027 pid:976290173\nhcl:#efcc98\nhgt:192in\n\npid:112431133 byr:1950 hgt:174cm\niyr:2020\ncid:118 hcl:#341e13 eyr:2023 ecl:amb\n\npid:034858755\nhcl:#d93689 iyr:2012 eyr:2025\nhgt:67cm\necl:brn byr:2027\ncid:306\n\neyr:2024 hcl:#fffffd ecl:hzl hgt:188cm cid:199 byr:2015 pid:983962091 iyr:1937\n\nhcl:#c0946f pid:899925634\neyr:2025 byr:2020\niyr:2016\necl:grt hgt:173cm\n\nhgt:59cm hcl:c5b2d7 byr:2008 iyr:2027\necl:lzr pid:155cm\neyr:2035\n\niyr:2014\neyr:2022 pid:850258746 hcl:#a97842 byr:2022 ecl:brn hgt:178cm\n\ncid:214 iyr:2017\necl:oth\nhcl:#866857 byr:1995 pid:793515973 hgt:193cm eyr:2023\n\nhcl:#18171d\niyr:2017 hgt:193cm cid:183 eyr:2025 pid:728034540 ecl:hzl byr:1969\n\neyr:2025 ecl:gry byr:2002 iyr:2019 hgt:174cm pid:603301922\nhcl:#fffffd\n\nbyr:2002\ncid:98 pid:828911903 eyr:2030 ecl:blu hgt:65in hcl:#74b1dc\n\nbyr:1969 hcl:#a97842 ecl:gry eyr:2027 pid:835656333 hgt:152cm cid:324 iyr:2014\n\npid:848442741 eyr:2030 hcl:#ceb3a1 byr:1984 iyr:2019 ecl:grn hgt:164cm\n\nhcl:#341e13 iyr:2019 hgt:166cm pid:419840849 byr:1974 eyr:2026 cid:211\n\nbyr:1945 pid:646444288 iyr:2020\neyr:2023 hgt:186cm\n\npid:375892516\nhgt:187cm\niyr:2010 eyr:2028 byr:1972 cid:272 ecl:blu hcl:#888785\n\nhgt:181in\necl:grn eyr:2034\nhcl:#7d3b0c byr:2018\npid:206240985 iyr:2015\n\nhgt:177 eyr:1973 pid:83092851 cid:92 ecl:utc byr:2023 hcl:z iyr:1948\n\neyr:2029 pid:1655089174 ecl:grn hgt:158cm iyr:2011 hcl:#b6652a byr:1926\ncid:158\n\nhcl:#341e13\niyr:2006\nbyr:2008 hgt:185 eyr:2024 ecl:utc\n\nhgt:171cm\npid:533365287 byr:1957 hcl:#ceb3a1 iyr:2014 ecl:amb eyr:2020\ncid:184\n\nhcl:#b6652a\npid:553897602 iyr:1929 ecl:grn cid:191 hgt:178cm byr:1991 eyr:2024\n\nbyr:1994 hgt:152cm pid:198152466\neyr:2022 ecl:hzl hcl:#4df239 iyr:2020\n\necl:grn\neyr:2022\nbyr:1968 iyr:2017 pid:044109096\n\nhcl:#d257c7 eyr:2036\niyr:2018\necl:#5b11eb\nbyr:1950"
	passports := strings.Split(input, "\n\n")
	validCount := 0
	for _, passport := range passports {
		pass := readPassport(passport)
		if checkRequiredFields(pass) {
			validCount++
			printPassport(pass)
		}
	}
	fmt.Println("Valid passports:", validCount)

}

func readPassport(passport string) Passport {
	passport = strings.Replace(passport, "\n", " ", -1)
	passportFields := strings.Split(passport, " ")
	passportObj := Passport{}
	for _, field := range passportFields {
		readField(&passportObj, field)
	}
	return passportObj
}

func readField(passport *Passport, field string) {
	splittedField := strings.Split(field, ":")
	fieldName := splittedField[0]
	fieldValue := strings.Trim(splittedField[1], " ")
	if strings.Contains(fieldName, "byr") {
		passport.byr, _ = strconv.Atoi(fieldValue)
	}
	if strings.Contains(fieldName, "iyr") {
		passport.iyr, _ = strconv.Atoi(fieldValue)
	}
	if strings.Contains(fieldName, "cid") {
		passport.cid = fieldValue
	}
	if strings.Contains(fieldName, "ecl") {
		passport.ecl = fieldValue
	}
	if strings.Contains(fieldName, "eyr") {
		passport.eyr, _ = strconv.Atoi(fieldValue)
	}
	if strings.Contains(fieldName, "hcl") {
		passport.hcl = fieldValue
	}
	if strings.Contains(fieldName, "hgt") {
		passport.hgt = fieldValue
	}
	if strings.Contains(fieldName, "pid") {
		passport.pid = fieldValue
	}
}

func checkRequiredFields(passport Passport) bool {
	return checkByr(passport) && checkIyr(passport) && checkEyr(passport) && checkHgt(passport) &&
		checkHcl(passport) && checkEcl(passport) && checkPid(passport)
}

func checkHgt(passport Passport) bool {
	heightValue, _ := strconv.Atoi(strings.Replace(strings.Replace(passport.hgt, "in", "", 1), "cm", "", 1))
	if strings.Contains(passport.hgt, "in") {
		return heightValue >= 59 && heightValue <= 76
	} else if strings.Contains(passport.hgt, "cm") {
		return heightValue >= 150 && heightValue <= 193
	} else {
		return false
	}
}

func checkEyr(passport Passport) bool {
	return passport.eyr >= 2020 && passport.eyr <= 2030
}

func checkIyr(passport Passport) bool {
	return passport.iyr >= 2010 && passport.iyr <= 2020
}

func checkByr(passport Passport) bool {
	return passport.byr >= 1920 && passport.byr <= 2002
}

func checkHcl(passport Passport) bool {
	if len(passport.hcl) != 7 {
		return false
	}

	hcl := passport.hcl[1:len(passport.hcl)]
	isValid := true
	for _, hclChar := range hcl {
		if !unicode.IsDigit(hclChar) && !unicode.In(hclChar, &unicode.RangeTable{
			R32: []unicode.Range32{{
				Lo:     'a',
				Hi:     'f',
				Stride: 1,
			}},
		}) {
			isValid = false
		}
	}
	return strings.HasPrefix(passport.hcl, "#") && isValid
}

func checkEcl(passport Passport) bool {
	return passport.ecl == "amb" || passport.ecl == "blu" || passport.ecl == "brn" ||
		passport.ecl == "gry" || passport.ecl == "grn" || passport.ecl == "hzl" ||
		passport.ecl == "oth"
}

func checkPid(passport Passport) bool {
	_, err := strconv.Atoi(passport.pid)
	return len(passport.pid) == 9 && err == nil
}

func printPassport(passport Passport) {
	fmt.Println("{byr:", passport.byr, "iyr:", passport.iyr,
		"eyr:", passport.eyr, "hgt:", passport.hgt, "hcl:", passport.hcl, "ecl:", passport.ecl, "pid:", passport.pid, "}")
}
