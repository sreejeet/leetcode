SELECT DISTINCT t1.Email FROM Person t1
INNER JOIN Person t2 
WHERE 
    t1.Id != t2.Id AND 
    t1.Email = t2.Email;