# https://leetcode.com/problems/trips-and-users/

SELECT
    t.request_at as Day,
    ROUND(
        COUNT(
            CASE WHEN t.status != 'completed' then 1 else null end
        ) / COUNT(*), 2)
    as 'Cancellation Rate' 
FROM Trips as t
WHERE t.request_at >= '2013-10-01' and t.request_at <= '2013-10-03'
AND t.client_id NOT IN (select users_id as id from Users where banned = 'Yes') 
AND t.driver_id NOT IN (select users_id as id from Users where banned = 'Yes') 
GROUP BY t.request_at