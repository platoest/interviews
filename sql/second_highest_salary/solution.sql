SELECT SecondHighestSalary FROM
(
    SELECT DISTINCT salary AS SecondHighestSalary FROM Employee
    UNION ALL
    SELECT NULL
) t
ORDER BY SecondHighestSalary DESC LIMIT 1 OFFSET 1