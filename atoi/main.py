class Solution:
    lowerlimit = -(2**31)
    upperlimit = 2**31 - 1
    def myAtoi(self, s: str) -> int:
        s = s.strip()
        strnum = ""
        negative = False
        negative_set = False
        for i in s:
            if i == " " and len(strnum) == 0 and not negative_set:
                continue
            elif (i == "-" or i == "+") and len(strnum) == 0:
                if negative_set:
                    return 0
                else:
                    negative = True if i == "-" else False
                    negative_set = True
                    continue
            else:
                pass 
            try:
                int(i)
            except ValueError:
                break

            strnum += i
        
        if len(strnum) == 0:
            return 0
        elif negative:
            num = -int(strnum)
            if num < self.lowerlimit:
                return self.lowerlimit
            else:
                return num
        else:
            num = int(strnum)
            if num > self.upperlimit:
                return self.upperlimit
            else:
                return num


