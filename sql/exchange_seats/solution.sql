select
case
When id%2=0 then id-1
when (id%2!=0 AND id=(select max(id) from Seat)) then id
when id%2!=0 then id+1
end as id,
student
from Seat
order by id