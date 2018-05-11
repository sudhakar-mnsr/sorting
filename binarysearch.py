array = [12,13,49,84,98,109,208,266]

def binarysearch(number):
    l=0
    r=len(array)-1
    while (l <= r):
       m = l+(r-l)/2
       if array[m] == number:
          return m
       if array[m] > number:
          r = m-1
       else:
          l=m+1
    return -1

number=277
found=binarysearch(number)
if found != -1:
   print "The number " + str(number) + ' ' + "is found at index" + ' ' +str(found)
else:
   print "The element is not present"
