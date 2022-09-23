# 2 Join
SELECT Emp_Name, Emp_No, Emp_Profile, Emp_Country, Emp_Join_Date 
FROM Tb1_Employee INNER JOIN Tb2_Employment 
ON Tb1_Employee.Emp_ID=Tb2_Employment.Emp_ID;

# 3 Join
SELECT Emp_Name, Emp_No, Emp_Profile, Emp_Country, EmpJoinDate, 
FROM Employee e INNER JOIN Employment m ON
e.Emp_ID = m.EMP_ID INNER JOIN EmpDetail d on 
d.Emp_Email = m.Emp_Email;

# Self Join
SELECT e.Emp_ID, e.Emp_Name, m.FullName as ManagerName
FROM Employees e JOIN Employees m ON e.ManagerId = m.Emp_ID
