def anagram(s1, s2):
    if(len(s1) != len(s2)):
        return False
    s1LetterCount = {}
    s2LetterCount = {}
    for ch in s1:
        if ch in s1LetterCount:
            s1LetterCount[ch] += 1
        else:
            s1LetterCount[ch] = 1
    for ch in s2:
        if ch in s2LetterCount:
            s2LetterCount[ch] +=1
        else:
            s2LetterCount[ch] = 1
    for key in s1LetterCount:
        if key not in s2LetterCount or s1LetterCount[key] != s2LetterCount[key]:
            return False
        return True

s1 = 'nameleses';
s2 = 'salesman'
print(anagram(s1,s2))