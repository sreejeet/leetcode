SELECT t2.class FROM
    (SELECT
        t1.class, COUNT(DISTINCT t1.student) 's_count'
    FROM
        courses t1
    GROUP BY t1.class
    ) t2
WHERE s_count > 4