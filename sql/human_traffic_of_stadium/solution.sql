# https://leetcode.com/problems/human-traffic-of-stadium/

select distinct s1.*
from Stadium as s1, Stadium as s2, Stadium as s3
where s1.people >= 100 and s2.people >= 100 and s3.people >= 100 and
(
    (s1.id - s2.id=1 and s2.id - s3.id =1 and s1.id - s3.id=2 ) # s1, s2, s3
    or
    (s2.id-s1.id=1 and s1.id - s3.id=1 and s2.id-s3.id = 2) # s2, s1, s3
    or 
    (s2.id - s1.id=1 and  s3.id-s1.id=2 and s3.id - s2.id =1 ) # s3, s2, s1
) ORDER by s1.id