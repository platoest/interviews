# https://leetcode.com/problems/market-analysis-i

SELECT
  u.user_id as buyer_id,
  u.join_date,
  COUNT(o.buyer_id) AS orders_in_2019
FROM
  Users as u
LEFT OUTER JOIN
  Orders as o
ON
    (o.buyer_id = u.user_id)
AND
    (year(o.order_date) = 2019)
GROUP BY
  u.user_id