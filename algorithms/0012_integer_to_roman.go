func intToRoman(num int) string {
    
    
    dict1 := map[int]string {
        1: "I",
        2: "II",
        3: "III",
        4: "IV",
        5: "V",
        6: "VI",
        7: "VII",
        8: "VIII",
        9: "IX",
    }
    
    dict10 := map[int]string {
        1: "X",
        2: "XX",
        3: "XXX",
        4: "XL",
        5: "L",
        6: "LX",
        7: "LXX",
        8: "LXXX",
        9: "XC",
    }
    
    dict100 := map[int]string {
        1: "C",
        2: "CC",
        3: "CCC",
        4: "CD",
        5: "D",
        6: "DC",
        7: "DCC",
        8: "DCCC",
        9: "CM",
    }
    
    dict1000 := map[int]string {
        1: "M",
        2: "MM",
        3: "MMM",
    }
    
    var result string
    
    if si, ok := dict1000[num/1000]; ok {
        result = si
        num = num % 1000
    }
    
    if si, ok := dict100[num/100]; ok {
        result += si
        num = num % 100
    }
    
    if si, ok := dict10[num/10]; ok {
        result += si
    }

    if si, ok := dict1[num%10]; ok {
        result += si
    }
    
    return result
}
